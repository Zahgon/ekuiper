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

type LogicalPlan interface {
	ExplainInfo
	Children() []LogicalPlan
	SetChildren(children []LogicalPlan)
	// PushDownPredicate pushes down the filter in the filter/where/on/having clauses as deeply as possible.
	// It will accept a condition that is an expression slice, and return the expressions that can't be pushed.
	// It is also return the new tree of plan as it can possibly change the tree
	PushDownPredicate(ast.Expr) (ast.Expr, LogicalPlan)
	// PruneColumns Prune the unused columns in the data source level, by pushing all needed columns down
	PruneColumns(fields []ast.Expr) error
}

type baseLogicalPlan struct {
	children []LogicalPlan
	// Can be used to return the derived instance from the base type
	self LogicalPlan
	// Interface for explaining
	ExplainInfo *PlanExplainInfo
}

type ExplainInfo interface {
	ID() int64
	Type() string
	ChildrenID() []int64
	Explain() string
	BuildExplainInfo()
	SetID(id int64)
}

type RuleRuntimeInfo interface {
	BuildSchemaInfo(ruleID string)
}

type PlanExplainInfo struct {
	T    PlanType `json:"-"`
	ID   int64    `json:"-"`
	Op   string   `json:"op"`
	Info string   `json:"info"`
}

func (p *PlanExplainInfo) SetOp() { _ = "STUB: not implemented"; return }

func (p *baseLogicalPlan) Explain() string { _ = "STUB: not implemented"; return "" }

func (p *baseLogicalPlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

func (p *baseLogicalPlan) SetID(id int64) { _ = "STUB: not implemented"; return }

func (p *baseLogicalPlan) Type() string { _ = "STUB: not implemented"; return "" }

func (p *baseLogicalPlan) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *baseLogicalPlan) ChildrenID() []int64 { _ = "STUB: not implemented"; return nil }

func (p *baseLogicalPlan) setPlanType(planType PlanType) { _ = "STUB: not implemented"; return }

func (p *baseLogicalPlan) Children() []LogicalPlan { _ = "STUB: not implemented"; return nil }

func (p *baseLogicalPlan) SetChildren(children []LogicalPlan) { _ = "STUB: not implemented"; return }

// PushDownPredicate By default, push down the predicate to the first child instead of the children
// as most plan cannot have multiple children
func (p *baseLogicalPlan) PushDownPredicate(condition ast.Expr) (ast.Expr, LogicalPlan) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(LogicalPlan)
}

func (p *baseLogicalPlan) PruneColumns(fields []ast.Expr) error {
	_ = "STUB: not implemented"
	return nil
}

type PlanType string

const (
	AGGREGATE     PlanType = "AggregatePlan"
	ANALYTICFUNCS PlanType = "AnalyticFuncsPlan"
	DATASOURCE    PlanType = "DataSourcePlan"
	FILTER        PlanType = "FilterPlan"
	HAVING        PlanType = "HavingPlan"
	JOINALIGN     PlanType = "JoinAlignPlan"
	JOIN          PlanType = "JoinPlan"
	LOOKUP        PlanType = "LookupPlan"
	ORDER         PlanType = "OrderPlan"
	PROJECT       PlanType = "ProjectPlan"
	PROJECTSET    PlanType = "ProjectSetPlan"
	WINDOW        PlanType = "WindowPlan"
	WINDOWFUNC    PlanType = "WindowFuncPlan"
	WATERMARK     PlanType = "WatermarkPlan"
	IncAggWindow  PlanType = "IncAggWindowPlan"
	AggFunc       PlanType = "AggFunc"
)
