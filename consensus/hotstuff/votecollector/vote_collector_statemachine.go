package votecollector

import (
	"fmt"
	"sync"

	"go.uber.org/atomic"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
)

type VoteCollectorStateMachine struct {
	BaseVoteCollector

	sync.Mutex
	collector atomic.Value
}

func (csm *VoteCollectorStateMachine) atomicLoadCollector() hotstuff.VoteCollectorState {
	return csm.collector.Load().(*atomicValueWrapper).collector
}

// atomic.Value doesn't allow storing interfaces as atomic values,
// it requires that stored type is always the same so we need a wrapper that will mitigate this restriction
// https://github.com/golang/go/issues/22550
type atomicValueWrapper struct {
	collector hotstuff.VoteCollectorState
}

func NewVoteCollectorStateMachine(base BaseVoteCollector) *VoteCollectorStateMachine {
	sm := &VoteCollectorStateMachine{
		BaseVoteCollector: base,
	}

	// by default start with caching collector
	sm.collector.Store(&atomicValueWrapper{
		collector: NewCachingVoteCollector(base),
	})
	return sm
}

func (csm *VoteCollectorStateMachine) AddVote(vote *model.Vote) (bool, error) {
	var (
		added bool
		err   error
	)
	for {
		collector := csm.atomicLoadCollector()
		currentState := collector.ProcessingStatus()
		added, err = collector.AddVote(vote)
		if err != nil {
			return false, fmt.Errorf("could not add vote %v: %w", vote.ID(), err)
		}
		if currentState != csm.ProcessingStatus() {
			continue
		}
		break
	}
	return added, nil
}

func (csm *VoteCollectorStateMachine) ProcessingStatus() hotstuff.ProcessingStatus {
	return csm.atomicLoadCollector().ProcessingStatus()
}

func (csm *VoteCollectorStateMachine) ChangeProcessingStatus(expectedValue, newValue hotstuff.ProcessingStatus) error {
	panic("implement me")
}
