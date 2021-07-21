package votecollector

import (
	"fmt"
	"go.uber.org/atomic"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
)

type VerifyingVoteCollector struct {
	BaseVoteCollector

	aggregator    CombinedAggregator
	reconstructor RandomBeaconReconstructor
	done          atomic.Bool
}

func NewVerifyingVoteCollector(base BaseVoteCollector) *VerifyingVoteCollector {
	return &VerifyingVoteCollector{
		BaseVoteCollector: base,
	}
}

func (c *VerifyingVoteCollector) AddVote(vote *model.Vote) (bool, error) {
	if c.done.Load() {
		return false, nil
	}

	verified, sigType, err := c.aggregator.Verify(vote.SignerID, vote.SigData)
	if err != nil {
		return false, fmt.Errorf("could not verify vote signature: %w", err)
	}

	// TODO: handle if verified == false
	if !verified {
		return false, fmt.Errorf("could not verify vote signature: %w", err)
	}

	if c.done.Load() {
		return false, nil
	}

	added, err := c.aggregator.TrustedAdd(vote.SignerID, vote.SigData, sigType)
	if err != nil {
		return false, fmt.Errorf("could not aggregate vote signature: %w", err)
	}

	// TODO: properly fix this case
	if !added {
		return false, fmt.Errorf("could not aggregate vote signature: %w", err)
	}

	//c.reconstructor.TrustedAdd(vote)

	// we haven't collected sufficient weight, we have nothing to do further
	if !c.aggregator.HasSufficientWeight() {
		return true, nil
	}

	// we haven't collected sufficient shares, we have nothing to do further
	if !c.reconstructor.HasSufficientShares() {
		return true, nil
	}

	err = c.buildQC()
	if err != nil {
		return false, fmt.Errorf("could not build QC: %w", err)
	}

	return true, nil
}

func (c *VerifyingVoteCollector) buildQC() error {
	// other goroutine might be constructing QC at this time, check with CAS
	// and exit early
	if !c.done.CAS(false, true) {
		return nil
	}

	// at this point we can be sure that no one else is creating QC

	stakingSig, thresholdSig, err := c.aggregator.AggregateSignature()
	if err != nil {
		return fmt.Errorf("could not construct aggregated signatures: %w", err)
	}

	randomBeaconSig, err := c.reconstructor.Reconstruct()
	if err != nil {
		return fmt.Errorf("could not reconstruct random beacon signature: %w", err)
	}

	// TODO: using aggregated signatures, construct QC

	return nil
}

func (VerifyingVoteCollector) ProcessingStatus() hotstuff.ProcessingStatus {
	return hotstuff.VerifyingVotes
}
