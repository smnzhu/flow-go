package synchronization

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/lifecycle"
	"github.com/onflow/flow-go/state/protocol"
)

// finalizedSnapshot is a helper structure which contains latest finalized header and participants list
type finalizedSnapshot struct {
	head         *flow.Header
	participants flow.IdentityList
}

// FinalizedSnapshotCache represents a cached snapshot of the latest finalized header and participants list.
// It is used in Engine to access latest valid data.
type FinalizedSnapshotCache struct {
	mu sync.RWMutex

	log                       zerolog.Logger
	state                     protocol.State
	identityFilter            flow.IdentityFilter
	lastFinalizedSnapshot     *finalizedSnapshot
	finalizationEventNotifier engine.Notifier // notifier for finalization events

	lm *lifecycle.LifecycleManager
}

// NewFinalizedSnapshotCache creates a new finalized snapshot cache.
func NewFinalizedSnapshotCache(log zerolog.Logger, state protocol.State, participantsFilter flow.IdentityFilter) (*FinalizedSnapshotCache, error) {

	cache := &FinalizedSnapshotCache{
		state:          state,
		identityFilter: participantsFilter,
		lm:             lifecycle.NewLifecycleManager(),
		log:            log.With().Str("component", "finalized_snapshot_cache").Logger(),
	}

	snapshot, err := cache.getSnapshot()
	if err != nil {
		return nil, fmt.Errorf("could not apply last finalized state")
	}

	cache.lastFinalizedSnapshot = snapshot

	return cache, nil
}

// get returns last locally stored snapshot which contains final header
// and list of filtered identities
func (f *FinalizedSnapshotCache) get() *finalizedSnapshot {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.lastFinalizedSnapshot
}

func (f *FinalizedSnapshotCache) getSnapshot() (*finalizedSnapshot, error) {
	finalSnapshot := f.state.Final()
	head, err := finalSnapshot.Head()
	if err != nil {
		return nil, fmt.Errorf("could not get last finalized header: %w", err)
	}

	// get all participant nodes from the state
	participants, err := finalSnapshot.Identities(f.identityFilter)
	if err != nil {
		return nil, fmt.Errorf("could not get consensus participants at latest finalized block: %w", err)
	}

	return &finalizedSnapshot{
		head:         head,
		participants: participants,
	}, nil
}

// updateSnapshot updates latest locally cached finalized snapshot
func (f *FinalizedSnapshotCache) updateSnapshot() error {
	snapshot, err := f.getSnapshot()
	if err != nil {
		return err
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	if f.lastFinalizedSnapshot.head.Height < snapshot.head.Height {
		f.lastFinalizedSnapshot = snapshot
	}

	return nil
}

func (f *FinalizedSnapshotCache) Ready() <-chan struct{} {
	f.lm.OnStart(func() {
		go f.finalizationProcessingLoop()
	})
	return f.lm.Started()
}

func (f *FinalizedSnapshotCache) Done() <-chan struct{} {
	f.lm.OnStop()
	return f.lm.Stopped()
}

// OnFinalizedBlock implements the `OnFinalizedBlock` callback from the `hotstuff.FinalizationConsumer`
//  (1) Updates local state of last finalized snapshot.
// CAUTION: the input to this callback is treated as trusted; precautions should be taken that messages
// from external nodes cannot be considered as inputs to this function
func (f *FinalizedSnapshotCache) OnFinalizedBlock(flow.Identifier) {
	// notify that there is new finalized block
	f.finalizationEventNotifier.Notify()
}

// finalizationProcessingLoop is a separate goroutine that performs processing of finalization events
func (f *FinalizedSnapshotCache) finalizationProcessingLoop() {
	notifier := f.finalizationEventNotifier.Channel()
	for {
		select {
		case <-f.lm.ShutdownSignal():
			return
		case <-notifier:
			err := f.updateSnapshot()
			if err != nil {
				f.log.Fatal().Err(err).Msg("could not process latest finalized block")
			}
		}
	}
}
