// Copyright 2025 EMQ Technologies Co., Ltd.
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
	"encoding/gob"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

const (
	V2WindowInputsKey = "$$v2windowInputs"
)

var InfTime = time.Unix(1<<63-62135596801, 999999999)

func init() {
	gob.Register([]*xsql.Tuple{})
	gob.Register(&WindowScanner{})
	gob.Register(time.Time{})
	gob.Register(&StateWindowStatus{})
	gob.Register(map[string]*StateWindowStatus{})
}

type WindowV2Operator struct {
	*defaultSinkNode
	windowConfig WindowConfig
	wExec        WindowV2Exec
	scanner      *WindowScanner
}

func NewWindowV2Op(name string, w WindowConfig, options *def.RuleOption) (*WindowV2Operator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *WindowV2Operator) Close() { _ = "STUB: not implemented"; return }

func (o *WindowV2Operator) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (o *WindowV2Operator) emitWindow(ctx api.StreamContext, startTime, endTime time.Time) {
	_ = "STUB: not implemented"
	return
}

type WindowV2Exec interface {
	exec(ctx api.StreamContext, errCh chan<- error)
}

type StateWindowOp struct {
	*WindowV2Operator
	status          map[string]*StateWindowStatus
	PartitionExpr   *ast.PartitionExpr
	SingleCondition ast.Expr
	BeginCondition  ast.Expr
	EmitCondition   ast.Expr
	stateFuncs      []*ast.Call
}

type StateWindowStatus struct {
	StartTime time.Time
	EndTime   time.Time
	OnBegin   bool
	Scanner   *WindowScanner
}

func NewStateWindowOp(o *WindowV2Operator) *StateWindowOp { _ = "STUB: not implemented"; return nil }

func (s *StateWindowOp) emit(ctx api.StreamContext, status *StateWindowStatus) {
	_ = "STUB: not implemented"
	return
}

func calPartition(fv *xsql.FunctionValuer, partitionExpr *ast.PartitionExpr, row *xsql.Tuple) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *StateWindowOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (s *StateWindowOp) handleTupleWithBeginEmitCondition(ctx api.StreamContext, fv *xsql.FunctionValuer, row *xsql.Tuple, status *StateWindowStatus) {
	_ = "STUB: not implemented"
	return
}

func (s *StateWindowOp) handleTupleWithSingleCondition(ctx api.StreamContext, fv *xsql.FunctionValuer, row *xsql.Tuple, status *StateWindowStatus) {
	_ = "STUB: not implemented"
	return
}

type SlidingWindowOp struct {
	*WindowV2Operator
	Delay            time.Duration
	Length           time.Duration
	stateFuncs       []*ast.Call
	triggerCondition ast.Expr
	delayNotify      chan time.Time
}

func NewSlidingWindowOp(o *WindowV2Operator) *SlidingWindowOp {
	_ = "STUB: not implemented"
	return nil
}

func (s *SlidingWindowOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func isMatchCondition(ctx api.StreamContext, condition ast.Expr, fv *xsql.FunctionValuer, d *xsql.Tuple, stateFuncs []*ast.Call) bool {
	_ = "STUB: not implemented"
	return false
}

// not match trigger condition

type WindowScanner struct {
	Tuples []*xsql.Tuple
}

func (s *WindowScanner) addTuple(tuple *xsql.Tuple) { _ = "STUB: not implemented"; return }

// scan left-open, right-closed window
func (s *WindowScanner) scanWindow(windowStart, windowEnd time.Time) []*xsql.Tuple {
	_ = "STUB: not implemented"
	return nil
}

// gc the tuples which earlier than gcTime
func (s *WindowScanner) gc(gcTime time.Time) { _ = "STUB: not implemented"; return }
