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
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	"github.com/lf-edge/ekuiper/v2/internal/topo/graph"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/internal/topo/operator"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

type genNodeFunc func(name string, props map[string]interface{}, options *def.RuleOption) (node.TopNode, error)

var extNodes = map[string]genNodeFunc{}

type sourceType int

const (
	ILLEGAL sourceType = iota
	STREAM
	SCANTABLE
	LOOKUPTABLE
)

// PlanByGraph returns a topo.Topo object by a graph
func PlanByGraph(rule *def.Rule) (*topo.Topo, error) { _ = "STUB: not implemented"; return nil, nil }

// handled above,

// Not all joins are lookup joins, so we need to create a join plan for the remaining joins

// validate source node

// reverse edges, value is a 2-dim array. Only switch node will have the second dim

// never happen

// sort the nodes by topological order

// validate the typo
// the map is to record the output for each node

// special case for join which does not allow multiple streams

// convert filter to having if the input is aggregated

// add the linkages

func genNodesInOrder(toNodes []string, edges map[string][]interface{}, flatReversedEdges map[string][]string, nodesInOrder []string, i int) int {
	_ = "STUB: not implemented"
	return 0
}

func parseSource(nodeName string, gn *def.GraphNode, rule *def.Rule, tp *topo.Topo, store kv.KeyValue, lookupTableChildren map[string]*ast.Options, isTemp bool) (node.DataSourceNode, sourceType, string, []node.OperatorNode, error) {
	_ = "STUB: not implemented"
	return *new(node.DataSourceNode), *new(sourceType), "", nil, nil
}

// If source name is specified, find the created stream/table from store

// Validate temp streams can only be used by temp rules

// Use the plan to calculate the schema and other meta info

func parseOrderBy(props map[string]interface{}, sourceNames []string) (*operator.OrderOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseGroupBy(props map[string]interface{}, sourceNames []string) (*operator.AggregateOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseJoinAst(props map[string]interface{}, sourceNames []string) (*ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseWatermark(props map[string]interface{}, streamEmitters map[string]struct{}) (*graph.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseWindow(props map[string]interface{}) (*node.WindowConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parsePick(props map[string]interface{}, sourceNames []string) (*operator.ProjectOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFunc(props map[string]interface{}, sourceNames []string) (*operator.FuncOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// never happen

func parseFilter(props map[string]interface{}, sourceNames []string) (*operator.FilterOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseHaving(props map[string]interface{}, sourceNames []string) (*operator.HavingOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSwitch(props map[string]interface{}, sourceNames []string) (*node.SwitchConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
