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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type DataSourcePlan struct {
	baseLogicalPlan
	name ast.StreamName
	// calculated properties
	// initialized with stream definition, pruned with rule
	metaFields []string
	// pass-on and converted state. For schemaless, the value is always nil
	streamFields map[string]*ast.JsonStreamField
	// pass-on properties
	isSchemaless    bool
	streamStmt      *ast.StreamStmt
	allMeta         bool
	isBinary        bool
	iet             bool
	timestampFormat string
	timestampField  string
	// col -> alias
	colAliasMapping map[string]string
	// intermediate status
	isWildCard  bool
	fields      map[string]*ast.JsonStreamField
	metaMap     map[string]string
	pruneFields []string
	// inRuleTest means whether in the rule test mode
	inRuleTest    bool
	useSliceTuple bool
}

func (p DataSourcePlan) Init() *DataSourcePlan { _ = "STUB: not implemented"; return nil }

func (p *DataSourcePlan) BuildSchemaInfo(ruleID string) { _ = "STUB: not implemented"; return }

func (p *DataSourcePlan) buildSchemaInfo(ruleID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *DataSourcePlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

// PushDownPredicate Presume no children for data source
func (p *DataSourcePlan) PushDownPredicate(condition ast.Expr) (ast.Expr, LogicalPlan) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(LogicalPlan)
}

// Add a filter plan for children

func (p *DataSourcePlan) extract(expr ast.Expr) (ast.Expr, ast.Expr) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), *new(ast.Expr)
}

func (p *DataSourcePlan) PruneColumns(fields []ast.Expr) error {
	_ = "STUB: not implemented"
	// init values
	return nil
}

// only allowed case like a.b.c

func buildArrowReference(cur ast.Expr, root map[string]interface{}) (map[string]interface{}, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// handleArrowFields mark the field and subField for the arrowFields which should be remained
// Then pruned the field which is not used.
func (p *DataSourcePlan) handleArrowFields(arrowFields []*ast.BinaryExpr) {
	_ = "STUB: not implemented"
	return
}

func pruneJSONStreamField(cur *ast.JsonStreamField) { _ = "STUB: not implemented"; return }

func markPruneJSONStreamField(cur interface{}, field *ast.JsonStreamField) {
	_ = "STUB: not implemented"
	return
}

func (p *DataSourcePlan) getField(name string, strict bool) (*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// always return nil for schemaless

// Do not prune fields now for preprocessor
// TODO provide field information to the source for it to prune
func (p *DataSourcePlan) getAllFields() { _ = "STUB: not implemented"; return }

// for consistency of results for testing

func (p *DataSourcePlan) getProps() error { _ = "STUB: not implemented"; return nil }
