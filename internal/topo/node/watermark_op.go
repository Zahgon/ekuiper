// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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

package node

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.opentelemetry.io/otel/trace"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

// WatermarkOp is used when event time is enabled.
// It is used to align the event time of the input streams
// It sends out the data in time order with watermark.
type WatermarkOp struct {
	*defaultSinkNode
	// config
	lateTolerance time.Duration
	sendWatermark bool
	// state
	events          []*xsql.Tuple // All the cached events in order
	rowHandle       map[any]trace.Span
	streamWMs       map[string]time.Time
	lastWatermarkTs time.Time
}

var _ OperatorNode = &WatermarkOp{}

const (
	WatermarkKey  = "$$wartermark"
	EventInputKey = "$$eventinputs"
	StreamWMKey   = "$$streamwms"
)

func NewWatermarkOp(name string, sendWatermark bool, streams []string, options *def.RuleOption) *WatermarkOp {
	_ = "STUB: not implemented"
	return nil
}

func (w *WatermarkOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// restore state

// whether to drop the late event

// If not drop, check if it can be sent out

func (w *WatermarkOp) track(ctx api.StreamContext, emitter string, ts time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// Add an event and check if watermark proceeds
// If yes, send out all events before the watermark
func (w *WatermarkOp) addAndTrigger(ctx api.StreamContext, d *xsql.Tuple) {
	_ = "STUB: not implemented"
	// Insert into the sorted array, should be faster than append then sort
	return
}

// Make sure watermark time proceeds

// Send out all events before the watermark

// Find out the last event to send in this watermark change

// Send out all events before the watermark

// The first event processing time start at the beginning of event receiving

// set the current span which will be set in broadcast

// Update watermark

// watermark is the minimum timestamp of all input topics
func (w *WatermarkOp) computeWatermarkTs() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
