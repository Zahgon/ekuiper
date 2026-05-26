// Copyright 2021-2023 EMQ Technologies Co., Ltd.
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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type WindowPlan struct {
	baseLogicalPlan
	PartitionExpr    *ast.PartitionExpr
	singleCondition  ast.Expr
	beginCondition   ast.Expr
	emitCondition    ast.Expr
	triggerCondition ast.Expr
	condition        ast.Expr
	wtype            ast.WindowType
	delay            int64
	length           int
	interval         int // If interval is not set, it is equals to Length
	timeUnit         ast.Token
	limit            int // If limit is not positive, there will be no limit
	isEventTime      bool

	stateFuncs []*ast.Call
}

func (p WindowPlan) Init() *WindowPlan { _ = "STUB: not implemented"; return nil }

func (p *WindowPlan) WindowType() ast.WindowType {
	_ = "STUB: not implemented"
	return *new(ast.WindowType)
}

func (p *WindowPlan) GetTriggerCondition() ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (p *WindowPlan) GetBeginCondition() ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }

func (p *WindowPlan) GetSingleCondition() ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (p *WindowPlan) GetEmitCondition() ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }

func (p *WindowPlan) GetPartitionExpr() *ast.PartitionExpr { _ = "STUB: not implemented"; return nil }

func (p *WindowPlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

func (p *WindowPlan) PushDownPredicate(condition ast.Expr) (ast.Expr, LogicalPlan) {
	_ = "STUB: not implemented"
	// not time window depends on the event, so should not filter any
	return *new(ast.Expr), *new(LogicalPlan)
}

// TODO event time filter, need event window op support
//p.condition = combine(condition, p.condition)
//// push nil condition won't return any
//p.baseLogicalPlan.PushDownPredicate(nil)
// return nil, p

// Presume window condition are only one table related.
// TODO window condition validation

func (p *WindowPlan) PruneColumns(fields []ast.Expr) error { _ = "STUB: not implemented"; return nil }

func (p *WindowPlan) ExtractStateFunc() { _ = "STUB: not implemented"; return }

func (p *WindowPlan) transform(f *ast.Call) { _ = "STUB: not implemented"; return }

func (p *WindowPlan) GenWindowConfig() *node.WindowConfig { _ = "STUB: not implemented"; return nil }
