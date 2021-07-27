package voteaggregatorv2

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/engine/common/fifoqueue"
)

// defaultVoteAggregatorWorkers number of workers to dispatch events for vote aggregators
const defaultVoteAggregatorWorkers = 8

type VoteAggregator struct {
	unit                 engine.Unit
	log                  zerolog.Logger
	collectors           VoteCollectors
	pendingVotes         *fifoqueue.FifoQueue
	pendingVotesNotifier engine.Notifier
}

// Ready returns a ready channel that is closed once the engine has fully
// started. For the propagation engine, we consider the engine up and running
// upon initialization.
func (va *VoteAggregator) Ready() <-chan struct{} {
	// launch as many workers as we need
	for i := 0; i < defaultVoteAggregatorWorkers; i++ {
		va.unit.Launch(va.pendingVotesProcessingLoop)
	}

	return va.unit.Ready()
}

func (va *VoteAggregator) Done() <-chan struct{} {
	return va.unit.Done()
}

func (va *VoteAggregator) pendingVotesProcessingLoop() {
	notifier := va.pendingVotesNotifier.Channel()
	for {
		select {
		case <-va.unit.Quit():
			return
		case <-notifier:
			err := va.processPendingVoteEvents()
			if err != nil {
				va.log.Fatal().Err(err).Msg("internal error processing block incorporated queued message")
			}
		}
	}
}

func (va *VoteAggregator) processPendingVoteEvents() error {
	for {
		select {
		case <-va.unit.Quit():
			return nil
		default:
		}

		msg, ok := va.pendingVotes.Pop()
		if ok {
			err := va.processPendingVote(msg.(*model.Vote))
			if err != nil {
				return fmt.Errorf("could not process incorporated block: %w", err)
			}
			continue
		}

		// when there is no more messages in the queue, back to the loop to wait
		// for the next incoming message to arrive.
		return nil
	}
}

func (va *VoteAggregator) processPendingVote(vote *model.Vote) error {
	lazyInitCollector, err := va.collectors.GetOrCreateCollector(vote.View, vote.BlockID)
	if err != nil {
		return fmt.Errorf("could not lazy init collector for view %d, blockID %v: %w",
			vote.View, vote.BlockID, err)
	}
	err = lazyInitCollector.Collector.AddVote(vote)
	if err != nil {
		return fmt.Errorf("could not process vote for view %d, blockID %v: %w",
			vote.View, vote.BlockID, err)
	}

	return nil
}

func (va *VoteAggregator) AddVote(vote *model.Vote) error {
	// It's ok to silently drop votes in case our processing pipeline is full.
	// It means that we are probably catching up.
	if ok := va.pendingVotes.Push(vote); ok {
		va.pendingVotesNotifier.Notify()
	}

	return nil
}

func (va *VoteAggregator) AddBlock(block *model.Block) error {
	// TODO: check block signature

	err := va.collectors.ProcessBlock(block)
	if err != nil {
		return fmt.Errorf("could not process block %v: %w", block.BlockID, err)
	}

	// after this call, collector might change state

	return nil
}

func (va *VoteAggregator) InvalidBlock(block *model.Block) {
	panic("implement me")
}

func (va *VoteAggregator) PruneByView(view uint64) error {
	err := va.collectors.PruneByView(view)
	if err != nil {
		return fmt.Errorf("could not prune by view %d: %w", view, err)
	}
	return nil
}
