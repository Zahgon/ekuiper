// Copyright 2024 EMQ Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

type HoppingWindowIncAggEventOp struct {
	op *HoppingWindowIncAggOp
	HoppingWindowIncAggEventOpState
}

type HoppingWindowIncAggEventOpState struct {
	CurrWindowList        []*IncAggWindow
	NextTriggerWindowTime time.Time
}

func NewHoppingWindowIncAggEventOp(o *WindowIncAggOperator) *HoppingWindowIncAggEventOp {
	_ = "STUB: not implemented"
	return nil
}

func (ho *HoppingWindowIncAggEventOp) PutState(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggEventOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (ho *HoppingWindowIncAggEventOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggEventOp) calIncAggWindow(ctx api.StreamContext, fv *xsql.FunctionValuer, row *xsql.Tuple, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggEventOp) emitWindow(ctx api.StreamContext, errCh chan<- error, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggEventOp) calIncAggWindowInEvent(ctx api.StreamContext, fv *xsql.FunctionValuer, row *xsql.Tuple) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggEventOp) triggerWindow(ctx api.StreamContext, now time.Time) {
	_ = "STUB: not implemented"
	return
}

type SlidingWindowIncAggEventOp struct {
	op *SlidingWindowIncAggOp
	SlidingWindowIncAggEventOpState
}

type SlidingWindowIncAggEventOpState struct {
	SlidingWindowIncAggOpState
	EmitList []*IncAggWindow
}

func NewSlidingWindowIncAggEventOp(o *WindowIncAggOperator) *SlidingWindowIncAggEventOp {
	_ = "STUB: not implemented"
	return nil
}

func (so *SlidingWindowIncAggEventOp) PutState(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (so *SlidingWindowIncAggEventOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (so *SlidingWindowIncAggEventOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (so *SlidingWindowIncAggEventOp) emitList(ctx api.StreamContext, errCh chan<- error, triggerTS time.Time) {
	_ = "STUB: not implemented"
	return
}

// emit nothing

// emit all windows

// emit part of windows

func (so *SlidingWindowIncAggEventOp) appendIncAggWindowInEvent(ctx api.StreamContext, errCh chan<- error, fv *xsql.FunctionValuer, row *xsql.Tuple) {
	_ = "STUB: not implemented"
	return
}

func (so *SlidingWindowIncAggEventOp) appendDelayIncAggWindowInEvent(ctx api.StreamContext, errCh chan<- error, fv *xsql.FunctionValuer, row *xsql.Tuple) {
	_ = "STUB: not implemented"
	return
}

type TumblingWindowIncAggEventOp struct {
	*HoppingWindowIncAggEventOp
}

func NewTumblingWindowIncAggEventOp(o *WindowIncAggOperator) *TumblingWindowIncAggEventOp {
	_ = "STUB: not implemented"
	return nil
}

func (to *TumblingWindowIncAggEventOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (o *WindowIncAggOperator) ingest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// watermark tuple should return

type CountWindowIncAggEventOp struct {
	op *CountWindowIncAggOp
	CountWindowIncAggEventOpState
}

type CountWindowIncAggEventOpState struct {
	CurrWindow     *IncAggWindow
	CurrWindowSize int
	EmitList       []*IncAggWindow
}

func NewCountWindowIncAggEventOp(o *WindowIncAggOperator) *CountWindowIncAggEventOp {
	_ = "STUB: not implemented"
	return nil
}

func (co *CountWindowIncAggEventOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (co *CountWindowIncAggEventOp) emitWindow(ctx api.StreamContext, errCh chan<- error, window *IncAggWindow, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (co *CountWindowIncAggEventOp) PutState(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (co *CountWindowIncAggEventOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}
