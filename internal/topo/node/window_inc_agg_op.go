// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	topoContext "github.com/lf-edge/ekuiper/v2/internal/topo/context"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

var EnableAlignWindow bool

func init() {
	EnableAlignWindow = true
	gob.Register(map[string]interface{}{})
	gob.Register(&xsql.Tuple{})
	gob.Register(&IncAggRange{})
	gob.Register(map[string]*IncAggWindow{})
	gob.Register(time.Time{})
	gob.Register(&IncAggWindow{})
	gob.Register([]*IncAggWindow{})
	gob.Register(CountWindowIncAggOpState{})
	gob.Register(TumblingWindowIncAggOpState{})
	gob.Register(SlidingWindowIncAggOpState{})
	gob.Register(SlidingWindowIncAggEventOpState{})
}

type WindowIncAggOperator struct {
	*defaultSinkNode
	windowConfig *WindowConfig
	Dimensions   ast.Dimensions
	aggFields    []*ast.Field
	WindowExec   windowIncAggExec

	putStateReqCh chan chan error
	restoreReqCh  chan chan error

	firstTimerMu      sync.Mutex
	firstTimerCreated bool
}

func NewWindowIncAggOp(name string, w *WindowConfig, dimensions ast.Dimensions, aggFields []*ast.Field, options *def.RuleOption) (*WindowIncAggOperator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *WindowIncAggOperator) Close() { _ = "STUB: not implemented"; return }

// Exec is the entry point for the executor
// input: *xsql.Tuple from preprocessor
// output: xsql.WindowTuplesSet
func (o *WindowIncAggOperator) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (o *WindowIncAggOperator) PutState4Test(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *WindowIncAggOperator) RestoreFromState4Test(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *WindowIncAggOperator) FirstTimerCreated4Test() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *WindowIncAggOperator) markFirstTimerCreated() { _ = "STUB: not implemented"; return }

func (o *WindowIncAggOperator) execStateCall4Test(ctx context.Context, reqCh chan chan error) error {
	_ = "STUB: not implemented"
	return nil
}

type windowIncAggExec interface {
	exec(ctx api.StreamContext, errCh chan<- error)
	PutState(ctx api.StreamContext)
	RestoreFromState(ctx api.StreamContext) error
}

type IncAggWindow struct {
	StartTime             time.Time
	EventTime             time.Time
	DimensionsIncAggRange map[string]*IncAggRange
}

func (w *IncAggWindow) Clone(ctx api.StreamContext) *IncAggWindow {
	_ = "STUB: not implemented"
	return nil
}

func (w *IncAggWindow) GenerateAllFunctionState() { _ = "STUB: not implemented"; return }

func (w *IncAggWindow) restoreState(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

type IncAggRange struct {
	fv   *xsql.FunctionValuer
	fctx *topoContext.DefaultContext

	FunctionState map[string]interface{}
	LastRow       *xsql.Tuple
	Fields        map[string]interface{}
}

func (r *IncAggRange) Clone(ctx api.StreamContext) *IncAggRange {
	_ = "STUB: not implemented"
	return nil
}

func (r *IncAggRange) generateFunctionState() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (r *IncAggRange) restoreState(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

type CountWindowIncAggOp struct {
	*WindowIncAggOperator
	windowSize int
	CountWindowIncAggOpState
}

type CountWindowIncAggOpState struct {
	CurrWindow     *IncAggWindow
	CurrWindowSize int
}

func (co *CountWindowIncAggOp) PutState(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (co *CountWindowIncAggOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (co *CountWindowIncAggOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (co *CountWindowIncAggOp) setIncAggWindow(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (co *CountWindowIncAggOp) incAggCal(ctx api.StreamContext, dimension string, row *xsql.Tuple, incAggWindow *IncAggWindow) {
	_ = "STUB: not implemented"
	return
}

func (co *CountWindowIncAggOp) emit(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

type TumblingWindowIncAggOp struct {
	*WindowIncAggOperator
	ticker     *clock.Ticker
	FirstTimer *clock.Timer
	Interval   time.Duration
	TumblingWindowIncAggOpState
}

type TumblingWindowIncAggOpState struct {
	CurrWindow *IncAggWindow
}

func NewTumblingWindowIncAggOp(o *WindowIncAggOperator) *TumblingWindowIncAggOp {
	_ = "STUB: not implemented"
	return nil
}

func (to *TumblingWindowIncAggOp) PutState(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (to *TumblingWindowIncAggOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (to *TumblingWindowIncAggOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (to *TumblingWindowIncAggOp) emit(ctx api.StreamContext, errCh chan<- error, now time.Time) {
	_ = "STUB: not implemented"
	return
}

type SlidingWindowIncAggOp struct {
	*WindowIncAggOperator
	triggerCondition ast.Expr
	Length           time.Duration
	Delay            time.Duration
	taskCh           chan *IncAggOpTask
	SlidingWindowIncAggOpState
}

type SlidingWindowIncAggOpState struct {
	CurrWindowList []*IncAggWindow
}

type IncAggOpTask struct {
	window *IncAggWindow
}

func NewSlidingWindowIncAggOp(o *WindowIncAggOperator) *SlidingWindowIncAggOp {
	_ = "STUB: not implemented"
	return nil
}

func (so *SlidingWindowIncAggOp) PutState(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (so *SlidingWindowIncAggOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (so *SlidingWindowIncAggOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (so *SlidingWindowIncAggOp) appendIncAggWindow(ctx api.StreamContext, errCh chan<- error, fv *xsql.FunctionValuer, row *xsql.Tuple, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (so *SlidingWindowIncAggOp) emit(ctx api.StreamContext, errCh chan<- error, window *IncAggWindow, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (so *SlidingWindowIncAggOp) isMatchCondition(ctx api.StreamContext, fv *xsql.FunctionValuer, d *xsql.Tuple) bool {
	_ = "STUB: not implemented"
	return false
}

// not match trigger condition

type HoppingWindowIncAggOp struct {
	*WindowIncAggOperator
	FirstTimer *clock.Timer
	ticker     *clock.Ticker
	Length     time.Duration
	Interval   time.Duration
	taskCh     chan *IncAggOpTask
	HoppingWindowIncAggOpState
}

type HoppingWindowIncAggOpState struct {
	CurrWindowList []*IncAggWindow
}

func NewHoppingWindowIncAggOp(o *WindowIncAggOperator) *HoppingWindowIncAggOp {
	_ = "STUB: not implemented"
	return nil
}

func (ho *HoppingWindowIncAggOp) PutState(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (ho *HoppingWindowIncAggOp) RestoreFromState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (ho *HoppingWindowIncAggOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggOp) newIncWindow(ctx api.StreamContext, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggOp) emit(ctx api.StreamContext, errCh chan<- error, window *IncAggWindow, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (ho *HoppingWindowIncAggOp) calIncAggWindow(ctx api.StreamContext, fv *xsql.FunctionValuer, row *xsql.Tuple, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func incAggCal(ctx api.StreamContext, dimension string, row *xsql.Tuple, incAggWindow *IncAggWindow, aggFields []*ast.Field) {
	_ = "STUB: not implemented"
	return
}

func newIncAggRange(ctx api.StreamContext) *IncAggRange { _ = "STUB: not implemented"; return nil }

func newIncAggWindow(ctx api.StreamContext, now time.Time) *IncAggWindow {
	_ = "STUB: not implemented"
	return nil
}

func calDimension(fv *xsql.FunctionValuer, dimensions ast.Dimensions, row *xsql.Tuple) string {
	_ = "STUB: not implemented"
	return ""
}

func gcIncAggWindow(currWindowList []*IncAggWindow, windowLength time.Duration, now time.Time) []*IncAggWindow {
	_ = "STUB: not implemented"
	return nil
}

func buildStateKey(ctx api.StreamContext) string { _ = "STUB: not implemented"; return "" }
