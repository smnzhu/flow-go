package voteaggregatorv2

import (
	"fmt"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
)

type VoteAggregator struct {
	collectors VoteCollectors
}

func (va *VoteAggregator) AddVote(vote *model.Vote) error {

	// TODO: This code should be executed by worker goroutine

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
