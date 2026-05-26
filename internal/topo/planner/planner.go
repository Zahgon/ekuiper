// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

func Plan(rule *def.Rule) (*topo.Topo, error) { _ = "STUB: not implemented"; return nil, nil }

// PlanSQLWithSourcesAndSinks For test only
func PlanSQLWithSourcesAndSinks(rule *def.Rule, mockSourcesProp map[string]map[string]any) (*topo.Topo, *ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// validation

// validate stmt

//if len(sources) > 0 && len(sources) != len(streamsFromStmt) {
//	return nil, fmt.Errorf("Invalid parameter sources or streams, the length cannot match the statement, expect %d sources.", len(streamsFromStmt))
//}

// Create the logical plan and optimize. Logical plans are a linked list

func updateFieldIndex(ctx api.StreamContext, stmt *ast.SelectStatement, af []*ast.Call, aff []*ast.Call) {
	_ = "STUB: not implemented"
	return
}

// ExprField: a visible, non-alias field whose top-level expression is
// not a *ast.FieldRef (e.g. CASE WHEN, arithmetic, bare function call).
// It gets exactly one output slot.  We assign SourceIndex to inner
// column FieldRefs but do NOT consume slots for them individually.

// handled by the generic walk below

// one slot for the whole ExprField output

// Add sink index for other non-select fields

// Set temp index for analytic funcs

func doUpdateIndex(ctx api.StreamContext, root ast.Node, index int, aliasIndex map[string]int) int {
	_ = "STUB: not implemented"
	return 0
}

func getSinkSchema(stmt *ast.SelectStatement) (map[string]*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateStmt(stmt *ast.SelectStatement) error { _ = "STUB: not implemented"; return nil }

func createTopo(rule *def.Rule, lp LogicalPlan, mockSourcesProp map[string]map[string]any, streamsFromStmt []string, schema map[string]*ast.JsonStreamField) (t *topo.Topo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create topology

// Cancel the partial topo on any planning error so that shared connection/stream
// subtopo refs and output channels registered during buildOps are cleaned up.
// Without this, a failed buildActions would leave the rule permanently registered
// in the shared subtopo's refRules with an unread output channel, causing
// "buffer full, drop message" errors for other rules.

// Add actions

func GetExplainInfoFromLogicalPlan(rule *def.Rule) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// validation

// Create logical plan and optimize. Logical plans are a linked list

func ExplainFromLogicalPlan(lp LogicalPlan, ruleID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Build the explainInfo of the current layer

// return the last schema if there are multiple sources
func buildOps(lp LogicalPlan, tp *topo.Topo, options *def.RuleOption, sources map[string]map[string]any, streamsFromStmt []string, index int) (node.Emitter, int, error) {
	_ = "STUB: not implemented"
	return *new(node.Emitter), 0, nil
}

// state window only support v2 window

func convertFromDuration(timeUnit ast.Token, length, interval int, delay int64) (time.Duration, time.Duration, time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration), *new(time.Duration)
}

func CreateLogicalPlan(stmt *ast.SelectStatement, opt *def.RuleOption, store kv.KeyValue) (LogicalPlan, error) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), nil
}

func checkSharedSourceOption(streams []*streamInfo, opt *def.RuleOption) error {
	_ = "STUB: not implemented"
	return nil
}

func createLogicalPlanFull(stmt *ast.SelectStatement, opt *def.RuleOption, store kv.KeyValue, isTemp bool) (LogicalPlan, []*ast.Call, []*ast.Call, error) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), nil, nil, nil
}

// If there are tables, the plan graph will be different for join/window

// If retainSize is not set, file table will try to read all the content in it

// TODO use interface to determine if the table is batch like file

//if opt.Experiment != nil && opt.Experiment.UseSliceTuple {
//	return nil, nil, nil, errors.New("slice tuple mode do not support dimensions or window yet")
//}

// if no interval value is set, and it's a count window, then set interval to length value.

// TODO calculate limit
// TODO incremental aggregate

// Not all joins are lookup joins, so we need to create a join plan for the remaining joins

// extract dedup trigger op

// In slice-tuple mode, assign a dedicated SinkContent slot to each
// ExprField (a non-alias, non-FieldRef visible expression such as a
// bare CASE WHEN or arithmetic expression).  The slot sequence must
// match the count used by updateFieldIndex: one slot per visible
// non-wildcard field in declaration order.

// wildcards consume no output slot

// extractSRFMapping extracts the set-returning-function in the field
func extractSRFMapping(stmt *ast.SelectStatement) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Transform(op node.UnOperation, name string, options *def.RuleOption) *node.UnaryOperator {
	_ = "STUB: not implemented"
	return nil
}

func extractWindowFuncFields(stmt *ast.SelectStatement) []*ast.Field {
	_ = "STUB: not implemented"
	return nil
}

func rewriteIfIncAggStmt(stmt *ast.SelectStatement, opt *def.RuleOption) []*ast.Field {
	_ = "STUB: not implemented"
	return nil
}

// TODO: support join later

func extractNodeIncAgg(node ast.Node, index *int) ([]*ast.Field, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func rewriteIntoBypass(newFieldRef *ast.FieldRef, f *ast.Call) { _ = "STUB: not implemented"; return }

func supportedWindowType(window *ast.Window) bool { _ = "STUB: not implemented"; return false }

var supportedWType = map[ast.WindowType]struct{}{
	ast.COUNT_WINDOW:    {},
	ast.SLIDING_WINDOW:  {},
	ast.HOPPING_WINDOW:  {},
	ast.TUMBLING_WINDOW: {},
}

func rewriteIfPushdownAlias(stmt *ast.SelectStatement, opt *def.RuleOption) map[ast.StreamName]map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func hasWildcard(stmt *ast.SelectStatement) bool { _ = "STUB: not implemented"; return false }

func searchColumnUsedCount(stmt *ast.SelectStatement, colName string) int {
	_ = "STUB: not implemented"
	return 0
}

func buildField(colName string, streamName ast.StreamName) ast.Field {
	_ = "STUB: not implemented"
	return *new(ast.Field)
}
