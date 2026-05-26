// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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
	"context"
	"encoding/gob"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.opentelemetry.io/otel/trace"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type WindowConfig struct {
	TriggerCondition ast.Expr
	StateFuncs       []*ast.Call
	Type             ast.WindowType
	// For time window
	Length   time.Duration
	Interval time.Duration // If the interval is not set, it is equals to Length
	Delay    time.Duration
	// For count window
	CountLength   int
	CountInterval int
	RawInterval   int
	TimeUnit      ast.Token

	// For SlidingWindow
	enableSlidingWindowSendTwice bool

	// For state window
	SingleCondition ast.Expr
	BeginCondition  ast.Expr
	EmitCondition   ast.Expr

	PartitionExpr *ast.PartitionExpr
}

type WindowOperator struct {
	*defaultSinkNode
	window          *WindowConfig
	interval        time.Duration
	duration        time.Duration
	isEventTime     bool
	isOverlapWindow bool
	trigger         *EventTimeTrigger // For event time only

	ticker *clock.Ticker // For processing time only
	// states
	triggerTime      time.Time
	msgCount         int
	delayTS          []time.Time
	triggerTS        []time.Time
	triggerCondition ast.Expr
	stateFuncs       []*ast.Call

	nextLink     trace.Link
	nextSpanCtx  context.Context
	nextSpan     trace.Span
	tupleSpanMap map[xsql.EventRow]trace.Span
}

const (
	WindowInputsKey = "$$windowInputs"
	TriggerTimeKey  = "$$triggerTime"
	MsgCountKey     = "$$msgCount"
)

func init() {
	gob.Register([]xsql.EventRow{})
	gob.Register([]map[string]interface{}{})
	gob.Register(map[string]time.Time{})
}

func validateWindowConfig(w WindowConfig) error { _ = "STUB: not implemented"; return nil }

func NewWindowOp(name string, w WindowConfig, options *def.RuleOption) (*WindowOperator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if no interval value is set, and it's a count window, then set interval to length value.

// Create watermark generator

func (o *WindowOperator) Close() { _ = "STUB: not implemented"; return }

// Exec is the entry point for the executor
// input: xsql.EventRow from preprocessor
// output: xsql.WindowTuplesSet
func (o *WindowOperator) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func getAlignedWindowEndTime(n time.Time, interval int, timeUnit ast.Token) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// The interval * days starting today

// should never happen

func getFirstTimer(ctx api.StreamContext, rawInerval int, timeUnit ast.Token) (time.Time, *clock.Timer) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (o *WindowOperator) execProcessingWindow(ctx api.StreamContext, inputs []xsql.EventRow, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// The first ticker to align the first window to the nature time

// resume the previous window

// send the last part

// process incoming item

// send the first part

// clear inputs if condition not matched
// TS add 1 to prevent remove current input

// For batching operator, do not end the span immediately so set it to nil

// If the deviation is less than 50ms, then process it. Otherwise, time may change and we'll start a new timer

// expire all inputs, so that when timer scans there is no item

// is cancelling

func (o *WindowOperator) setupTicker() { _ = "STUB: not implemented"; return }

func (o *WindowOperator) tick(ctx api.StreamContext, inputs []xsql.EventRow, n time.Time, log api.Logger) []xsql.EventRow {
	_ = "STUB: not implemented"
	return nil
}

type TupleList struct {
	tuples []xsql.EventRow
	index  int // Current index
	size   int // The size for count window
}

func NewTupleList(tuples []xsql.EventRow, windowSize int) (TupleList, error) {
	_ = "STUB: not implemented"
	return *new(TupleList), nil
}

func (tl *TupleList) hasMoreCountWindow() bool { _ = "STUB: not implemented"; return false }

func (tl *TupleList) count() int { _ = "STUB: not implemented"; return 0 }

func (tl *TupleList) nextCountWindow() *xsql.WindowTuples { _ = "STUB: not implemented"; return nil }

func (tl *TupleList) getRestTuples() []xsql.EventRow { _ = "STUB: not implemented"; return nil }

func (o *WindowOperator) isTimeRelatedWindow() bool { _ = "STUB: not implemented"; return false }

func isOverlapWindow(winType ast.WindowType) bool { _ = "STUB: not implemented"; return false }

func (o *WindowOperator) handleInputsForSlidingWindow(ctx api.StreamContext, inputs []xsql.EventRow, windowStart, windowEnd time.Time) ([]xsql.EventRow, []xsql.EventRow, []xsql.EventRow) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (o *WindowOperator) handleInputs(ctx api.StreamContext, inputs []xsql.EventRow, right time.Time) ([]xsql.EventRow, []xsql.EventRow, []xsql.EventRow) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Sync table

// this is to avoid always scan all tuples. better for performance if a window is big.

// Assume the inputs are sorted by timestamp

// Other window always discard the tuples that has been triggered.
// So the tuple in the inputs should all bigger than the current left (in the window)
// For hopping and sliding window, firstly check if the beginning tuples are expired and discard them

// Expired tuple, remove it by not adding back to inputs

// Now all tuples are in the window. Next step is to check if the tuple is in the current window
// If the tuple is beyond the right boundary, then it should be in the next window

func (o *WindowOperator) gcInputs(inputs []xsql.EventRow, triggerTime time.Time, ctx api.StreamContext) []xsql.EventRow {
	_ = "STUB: not implemented"
	return nil
}

func (o *WindowOperator) scan(inputs []xsql.EventRow, triggerTime time.Time, ctx api.StreamContext, length time.Duration, isFirstPart bool) []xsql.EventRow {
	_ = "STUB: not implemented"
	return nil
}

func (o *WindowOperator) calDelta(triggerTime time.Time, log api.Logger) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// max int, all events for the initial window

func (o *WindowOperator) isMatchCondition(ctx api.StreamContext, d xsql.EventRow) bool {
	_ = "STUB: not implemented"
	return false
}

// not match trigger condition

// match trigger condition

func (o *WindowOperator) handleTraceIngestTuple(_ api.StreamContext, t xsql.EventRow) {
	_ = "STUB: not implemented"
	return
}

func (o *WindowOperator) handleTraceDiscardTuple(ctx api.StreamContext, tuples []xsql.EventRow) {
	_ = "STUB: not implemented"
	return
}

func (o *WindowOperator) handleTraceEmitTuple(ctx api.StreamContext, wt *xsql.WindowTuples) {
	_ = "STUB: not implemented"
	return
}

// discard span if windowTuple is empty

func (o *WindowOperator) handleNextWindowTupleSpan(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}
