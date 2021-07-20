package votecollector

import (
	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
)

type CollectorStateMachine struct{}

func (CollectorStateMachine) AddVote(vote *model.Vote) (bool, error) {
	panic("implement me")
}

func (CollectorStateMachine) BlockID() flow.Identifier {
	panic("implement me")
}

func (CollectorStateMachine) ProcessingStatus() hotstuff.ProcessingStatus {
	panic("implement me")
}

func (CollectorStateMachine) ChangeProcessingStatus(expectedValue, newValue hotstuff.ProcessingStatus) error {
	panic("implement me")
}
