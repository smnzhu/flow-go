package votecollector

import (
	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
)

type VerifyingVoteCollector struct{}

func (VerifyingVoteCollector) AddVote(vote *model.Vote) (bool, error) {
	panic("implement me")
}

func (VerifyingVoteCollector) ProcessingStatus() hotstuff.ProcessingStatus {
	panic("implement me")
}
