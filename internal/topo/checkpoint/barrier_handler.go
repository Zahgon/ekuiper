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
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type BarrierHandler interface {
	Process(data *BufferOrEvent, ctx api.StreamContext) bool // If data is barrier return true, else return false
	SetOutput(chan<- *BufferOrEvent)                         // It is using for block a channel
}

// For qos 1, simple track barriers
type BarrierTracker struct {
	responder          Responder
	inputCount         int
	pendingCheckpoints map[int64]int
}

func NewBarrierTracker(responder Responder, inputCount int) *BarrierTracker {
	_ = "STUB: not implemented"
	return nil
}

func (h *BarrierTracker) Process(data *BufferOrEvent, ctx api.StreamContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *BarrierTracker) SetOutput(_ chan<- *BufferOrEvent) {
	_ = "STUB: not implemented"
	// do nothing, does not need it
	return
}

func (h *BarrierTracker) processBarrier(b *Barrier, ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

// For qos 2, block an input until all barriers are received
type BarrierAligner struct {
	responder           Responder
	inputCount          int
	currentCheckpointId int64
	output              chan<- *BufferOrEvent
	blockedChannels     map[string]bool
	buffer              []*BufferOrEvent
}

func NewBarrierAligner(responder Responder, inputCount int) *BarrierAligner {
	_ = "STUB: not implemented"
	return nil
}

func (h *BarrierAligner) Process(data *BufferOrEvent, ctx api.StreamContext) bool {
	_ = "STUB: not implemented"
	return false
}

// If blocking, save to buffer

func (h *BarrierAligner) processBarrier(b *Barrier, ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

// TODO Abort checkpoint

// clean up all the buffer

func (h *BarrierAligner) onBarrier(name string, ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (h *BarrierAligner) SetOutput(output chan<- *BufferOrEvent) { _ = "STUB: not implemented"; return }

func (h *BarrierAligner) releaseBlocksAndResetBarriers() { _ = "STUB: not implemented"; return }

func (h *BarrierAligner) beginNewAlignment(barrier *Barrier, ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}
