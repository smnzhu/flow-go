package votecollector

import (
	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
)

type VerifyingVoteCollector struct {
	BaseVoteCollector
}

func NewVerifyingVoteCollector(base BaseVoteCollector) *VerifyingVoteCollector {
	return &VerifyingVoteCollector{
		BaseVoteCollector: base,
	}
}

func (VerifyingVoteCollector) AddVote(vote *model.Vote) (bool, error) {
	panic("implement me")
}

func (VerifyingVoteCollector) ProcessingStatus() hotstuff.ProcessingStatus {
	panic("implement me")
}
