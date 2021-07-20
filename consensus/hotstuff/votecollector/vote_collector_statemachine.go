package votecollector

import (
	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
)

type VoteCollectorStateMachine struct{}

func (VoteCollectorStateMachine) AddVote(vote *model.Vote) (bool, error) {
	panic("implement me")
}

func (VoteCollectorStateMachine) BlockID() flow.Identifier {
	panic("implement me")
}

func (VoteCollectorStateMachine) ProcessingStatus() hotstuff.ProcessingStatus {
	panic("implement me")
}

func (VoteCollectorStateMachine) ChangeProcessingStatus(expectedValue, newValue hotstuff.ProcessingStatus) error {
	panic("implement me")
}
