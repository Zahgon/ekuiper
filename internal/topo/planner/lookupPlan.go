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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

// LookupPlan is the plan for table lookup and then merged/joined
type LookupPlan struct {
	baseLogicalPlan
	joinExpr   ast.Join
	keys       []string
	fields     []string
	valvars    []ast.Expr
	options    *ast.Options
	conditions ast.Expr
}

// Init must run validateAndExtractCondition before this func
func (p LookupPlan) Init() *LookupPlan { _ = "STUB: not implemented"; return nil }

func (p *LookupPlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

// PushDownPredicate do not deal with conditions, push down or return up
func (p *LookupPlan) PushDownPredicate(condition ast.Expr) (ast.Expr, LogicalPlan) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(LogicalPlan)
}

// Swallow all filter conditions. If there are other filter plans, there may have multiple filters

// Add a filter plan for children

// Return the unpushable condition and pushable condition
func extractLookupCondition(condition ast.Expr, tableName string) (unpushable ast.Expr, pushable ast.Expr) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(ast.Expr)
}

// default case: all condition are unpushable

// validateAndExtractCondition Make sure the join condition is equi-join and extreact other conditions
func (p *LookupPlan) validateAndExtractCondition() bool { _ = "STUB: not implemented"; return false }

// No equal predict condition found

// Extract equi-join condition

// flatConditions flat the join condition. Only binary condition of EQ and AND are allowed
func flatConditions(condition ast.Expr) ([]*ast.BinaryExpr, []ast.Expr) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LookupPlan) PruneColumns(fields []ast.Expr) error { _ = "STUB: not implemented"; return nil }
