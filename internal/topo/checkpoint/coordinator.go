// Copyright 2021-2024 EMQ Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checkpoint

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type pendingCheckpoint struct {
	checkpointId   int64
	isDiscarded    bool
	notYetAckTasks map[string]bool
}

func newPendingCheckpoint(checkpointId int64, tasksToWaitFor []Responder) *pendingCheckpoint {
	_ = "STUB: not implemented"
	return nil
}

func (c *pendingCheckpoint) ack(opId string) bool { _ = "STUB: not implemented"; return false }

// TODO serialize state

func (c *pendingCheckpoint) isFullyAck() bool { _ = "STUB: not implemented"; return false }

func (c *pendingCheckpoint) finalize() *completedCheckpoint { _ = "STUB: not implemented"; return nil }

func (c *pendingCheckpoint) dispose(_ bool) { _ = "STUB: not implemented"; return }

type completedCheckpoint struct {
	checkpointId int64
}

type checkpointStore struct {
	syncx.RWMutex
	maxNum      int
	checkpoints []*completedCheckpoint
}

func (s *checkpointStore) add(c *completedCheckpoint) { _ = "STUB: not implemented"; return }

func (s *checkpointStore) getLatest() *completedCheckpoint { _ = "STUB: not implemented"; return nil }

func (s *checkpointStore) getCount() int { _ = "STUB: not implemented"; return 0 }

type Coordinator struct {
	toBeClean               int
	tasksToTrigger          []Responder
	tasksToWaitFor          []Responder
	sinkTasks               []SinkTask
	pendingCheckpoints      *sync.Map
	completedCheckpoints    *checkpointStore
	ruleId                  string
	baseInterval            time.Duration
	cleanThreshold          int
	advanceToEndOfEventTime bool
	ticker                  *clock.Ticker // For processing time only
	signal                  chan *Signal
	store                   api.Store
	ctx                     api.StreamContext
	activated               atomic.Bool

	inForceSaveState     atomic.Bool
	forceSaveStateNotify chan any
}

func NewCoordinator(ruleId string, sources []StreamTask, operators []NonSourceTask, sinks []SinkTask, qos def.Qos, store api.Store, interval time.Duration, ctx api.StreamContext) *Coordinator {
	_ = "STUB: not implemented"
	return nil
}

// 5 minutes by default

func createBarrierHandler(re Responder, inputCount int, qos def.Qos) BarrierHandler {
	_ = "STUB: not implemented"
	return *new(BarrierHandler)
}

func (c *Coordinator) Activate() error { _ = "STUB: not implemented"; return nil }

func (c *Coordinator) saveState(n time.Time, logger api.Logger) {
	_ = "STUB: not implemented"
	// trigger checkpoint
	// TODO pose max attempt and min pause check for consequent pendingCheckpoints
	return
}

// TODO Check if all tasks are running

// Create a pending checkpoint

// Let the sources send out a barrier

func (c *Coordinator) Deactivate() error { _ = "STUB: not implemented"; return nil }

func (c *Coordinator) ForceSaveState() (chan any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Coordinator) FinishForceSaveState() { _ = "STUB: not implemented"; return }

func (c *Coordinator) cancel(checkpointId int64) { _ = "STUB: not implemented"; return }

func (c *Coordinator) complete(checkpointId int64) { _ = "STUB: not implemented"; return }

// TODO handle checkpoint error

// Drop the previous pendingCheckpoints

// TODO revisit how to abort a checkpoint, discard callback

// For testing
func (c *Coordinator) GetCompleteCount() int { _ = "STUB: not implemented"; return 0 }

func (c *Coordinator) GetLatest() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Coordinator) IsActivated() bool { _ = "STUB: not implemented"; return false }

func (c *Coordinator) ActiveForceSaveState() { _ = "STUB: not implemented"; return }
