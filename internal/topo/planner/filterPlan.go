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
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type FilterPlan struct {
	baseLogicalPlan
	condition  ast.Expr
	stateFuncs []*ast.Call
}

func (p FilterPlan) Init() *FilterPlan { _ = "STUB: not implemented"; return nil }

func (p *FilterPlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

func (p *FilterPlan) PushDownPredicate(condition ast.Expr) (ast.Expr, LogicalPlan) {
	_ = "STUB: not implemented"
	// if no child, swallow all conditions
	return *new(ast.Expr), *new(LogicalPlan)
}

// eliminate this filter

func (p *FilterPlan) PruneColumns(fields []ast.Expr) error { _ = "STUB: not implemented"; return nil }

func (p *FilterPlan) ExtractStateFunc() { _ = "STUB: not implemented"; return }

func (p *FilterPlan) transform(f *ast.Call) { _ = "STUB: not implemented"; return }
