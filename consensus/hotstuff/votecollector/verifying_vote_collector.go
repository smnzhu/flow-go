package votecollector

import (
	"fmt"
	"github.com/onflow/flow-go/model/flow"

	"go.uber.org/atomic"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
)

type VerifyingVoteCollector struct {
	BaseVoteCollector

	dkg           hotstuff.DKG
	aggregator    CombinedAggregator
	reconstructor RandomBeaconReconstructor
	qcBuilder     QCBuilder
	done          atomic.Bool
}

func NewVerifyingVoteCollector(base BaseVoteCollector) *VerifyingVoteCollector {
	return &VerifyingVoteCollector{
		BaseVoteCollector: base,
	}
}

func (c *VerifyingVoteCollector) AddVote(vote *model.Vote) (*flow.QuorumCertificate, bool, error) {
	if c.done.Load() {
		return nil, false, nil
	}

	verified, sigType, err := c.aggregator.Verify(vote.SignerID, vote.SigData)
	if err != nil {
		return nil, false, fmt.Errorf("could not verify vote signature: %w", err)
	}

	// TODO: handle if verified == false
	if !verified {
		return nil, false, fmt.Errorf("could not verify vote signature: %w", err)
	}

	if c.done.Load() {
		return nil, false, nil
	}

	_, err = c.aggregator.TrustedAdd(vote.SignerID, vote.SigData, sigType)
	if err != nil {
		return nil, false, fmt.Errorf("could not aggregate vote signature: %w", err)
	}

	if sigType == SigTypeThreshold {
		index, err := c.dkg.Index(vote.SignerID)
		if err != nil {
			return nil, false, fmt.Errorf("could not retrieve dkg index for signer (%v): %w", vote.SignerID, err)
		}
		_, err = c.reconstructor.TrustedAdd(index, vote.SigData)
		if err != nil {
			return nil, false, fmt.Errorf("could not add random beacon sig share: %w", err)
		}

	}

	// we haven't collected sufficient weight, we have nothing to do further
	if !c.aggregator.HasSufficientWeight() {
		return nil, false, nil
	}

	// we haven't collected sufficient shares, we have nothing to do further
	if !c.reconstructor.HasSufficientShares() {
		return nil, false, nil
	}

	qc, err := c.buildQC()
	if err != nil {
		return nil, false, fmt.Errorf("could not build QC: %w", err)
	}

	return qc, qc != nil, nil
}

func (c *VerifyingVoteCollector) buildQC() (*flow.QuorumCertificate, error) {
	// other goroutine might be constructing QC at this time, check with CAS
	// and exit early
	if !c.done.CAS(false, true) {
		return nil, nil
	}

	// at this point we can be sure that no one else is creating QC

	// aggregator returns two signatures, one is aggregated staking signature
	// another one is aggregated threshold signature
	stakingSig, thresholdSig, err := c.aggregator.AggregateSignature()
	if err != nil {
		return nil, fmt.Errorf("could not construct aggregated signatures: %w", err)
	}

	// reconstructor returns random beacon signature reconstructed from threshold signature shares
	randomBeaconSig, err := c.reconstructor.Reconstruct()
	if err != nil {
		return nil, fmt.Errorf("could not reconstruct random beacon signature: %w", err)
	}

	qc, err := c.qcBuilder.CreateQC(stakingSig, thresholdSig, randomBeaconSig)
	if err != nil {
		return nil, fmt.Errorf("could not create quorum certificate: %w", err)
	}

	return qc, nil
}

func (VerifyingVoteCollector) ProcessingStatus() hotstuff.ProcessingStatus {
	return hotstuff.VerifyingVotes
}
