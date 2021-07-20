package votecollector

import (
	"fmt"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
)

type CachingVoteCollector struct {
	BaseVoteCollector
	pendingVotes *PendingVotes
}

func NewCachingVoteCollector(base BaseVoteCollector) *CachingVoteCollector {
	return &CachingVoteCollector{
		BaseVoteCollector: base,
		pendingVotes:      NewPendingVotes(),
	}
}

func (c *CachingVoteCollector) AddVote(vote *model.Vote) (bool, error) {
	if vote.BlockID != c.blockID {
		return false, fmt.Errorf("this CachingVoteCollector processes votes for blockID (%x), "+
			"but got a vote for (%x)", c.blockID, vote.BlockID)
	}
	return c.pendingVotes.AddVote(vote), nil
}

func (c *CachingVoteCollector) BlockID() flow.Identifier {
	return c.blockID
}

func (c *CachingVoteCollector) ProcessingStatus() hotstuff.ProcessingStatus {
	return hotstuff.CachingVotes
}

func (c *CachingVoteCollector) GetVotes() []*model.Vote {
	return c.pendingVotes.All()
}
