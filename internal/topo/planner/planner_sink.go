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

package planner

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

// SinkPlanner is the planner for sink node. It transforms logical sink plan to multiple physical nodes.
// It will split the sink plan into multiple sink nodes according to its sink configurations.

func buildActions(tp *topo.Topo, rule *def.Rule, inputs []node.Emitter, streamCount int, schema map[string]*ast.JsonStreamField) error {
	_ = "STUB: not implemented"
	return nil
}

func copyProps(raw map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func PlanSinkOps(tp *topo.Topo, inputs []node.Emitter, cn node.CompNode) {
	_ = "STUB: not implemented"
	return
}

// The case order is important, because sink node is also operator node

// resend

func SinkToComp(tp *topo.Topo, sinkType string, sinkName string, props map[string]any, rule *def.Rule, streamCount int, schema map[string]*ast.JsonStreamField) (node.CompNode, error) {
	_ = "STUB: not implemented"
	return *new(node.CompNode), nil
}

// consume common conf in sink itself, swallow and not pass it to common conf

// Split sink node

// Cache in alter queue, the topo becomes sink (fail) -> cache -> resendSink
// If no alter queue, the topo is cache -> sink

// TODO currently, the destination prop must be named topic

func findTemplateProps(props map[string]any) []string { _ = "STUB: not implemented"; return nil }

// Split sink node according to the sink configuration. Return the new input emitters.
func splitSink(tp *topo.Topo, s api.Sink, sinkName string, options *def.RuleOption, sc *node.SinkConf, templates []string, schema map[string]*ast.JsonStreamField) ([]node.TopNode, error) {
	_ = "STUB: not implemented"
	// tailor schema, each sink may have different transform field
	return nil, nil
}

// Batch enabled

// Shallow copy sc to avoid modifying the original

// Transform enabled
// Currently, the row to map is done here and is required. TODO: eliminate map and this could become optional

// Encode will convert the result to []byte

// Caching

func washSchema(sc *node.SinkConf, schema map[string]*ast.JsonStreamField) map[string]*ast.JsonStreamField {
	_ = "STUB: not implemented"
	return nil
}

// set DataField, we don't know the final schema then

type SinkCompNode struct {
	name  string
	nodes []node.TopNode
}

func (s *SinkCompNode) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *SinkCompNode) Nodes() []node.TopNode { _ = "STUB: not implemented"; return nil }

var _ node.CompNode = &SinkCompNode{}
