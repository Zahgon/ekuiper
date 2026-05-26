// Copyright 2021 EMQ Technologies Co., Ltd.
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

import "github.com/lf-edge/ekuiper/v2/pkg/ast"

type JoinPlan struct {
	baseLogicalPlan
	from  *ast.Table
	joins ast.Joins
}

func (p JoinPlan) Init() *JoinPlan { _ = "STUB: not implemented"; return nil }

func (p *JoinPlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

func (p *JoinPlan) PushDownPredicate(condition ast.Expr) (ast.Expr, LogicalPlan) {
	_ = "STUB: not implemented"
	// TODO multiple join support
	// Assume only one join
	return *new(ast.Expr), *new(LogicalPlan)
}

// always swallow all conditions

// TODO fine grain handling for left/right join

// never swallow anything

// Return the unpushable condition and pushable condition
func extractCondition(condition ast.Expr) (unpushable ast.Expr, pushable ast.Expr) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(ast.Expr)
}

// default case: all condition are unpushable

func (p *JoinPlan) PruneColumns(fields []ast.Expr) error { _ = "STUB: not implemented"; return nil }
