package votecollector

import "github.com/onflow/flow-go/model/flow"

type BaseVoteCollector struct {
	blockID flow.Identifier
}

func NewBaseVoteCollector(blockID flow.Identifier) BaseVoteCollector {
	return BaseVoteCollector{
		blockID: blockID,
	}
}
