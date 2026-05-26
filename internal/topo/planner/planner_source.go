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
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

func transformSourceNode(ctx api.StreamContext, t *DataSourcePlan, mockSourcesProp map[string]map[string]any, ruleId string, options *def.RuleOption, index int) (node.DataSourceNode, []node.OperatorNode, int, error) {
	_ = "STUB: not implemented"
	return *new(node.DataSourceNode), nil, 0, nil
}

func splitSource(ctx api.StreamContext, t *DataSourcePlan, ss api.Source, options *def.RuleOption, mockProps map[string]any, index int, ruleId string, pp node.UnOperation, emitterName ast.StreamName) (_ node.DataSourceNode, _ []node.OperatorNode, _ int, err error) {
	_ = "STUB: not implemented"
	// Get all props
	return *new(node.DataSourceNode), nil, 0, nil
}

// Create the connector node as source node

// Some connection only allow one subscription. The source should implement UniqueSub to provide a subId to avoid multiple connection.

// Some connection only have on

// If having unique connection id AND unique sub id for each connection, need to share the sub node; Case 1 is neuron; Case 2 is edgeX

// connection selector is set as a one node sub_topo

// For shared stream, the rule id is the shared subtopo. Subtopo only has one run so runId is always 0

// If splitSource fails before returning, clean up the ref early,
// as it has not been attached to the topo yet.

// another node to set emitter

// Need to check after source has provisioned, so do not put it before provision

//if t.isWildCard {
//	schema = nil
//}
// Create the decode node

//if t.isWildCard {
//	schema = nil
//}
// Create the decode node

// Create the preprocessor node if needed

// Create subtopo in the end to avoid errors in the middle

type SourcePropsForSplit struct {
	Decompression string            `json:"decompression"`
	SelId         string            `json:"connectionSelector"`
	PayloadFormat string            `json:"payloadFormat"`
	Interval      cast.DurationConf `json:"interval"`
	// merger and mergerField should only set one
	MergeField string `json:"mergeField"`
	Merger     string `json:"merger"`
	Format     string `json:"format"`
}

type traits struct {
	needConnection    bool
	needCompression   bool
	needDecode        bool
	needPayloadDecode bool
	// rate limit will plan right after source read
	needRatelimit bool
	// rate limit merge will plan after decompress
	needRatelimitMerge bool
}

// function to return if a sub node is needed
func checkFeatures(ss api.Source, sp *SourcePropsForSplit, props map[string]any) (traits, error) {
	_ = "STUB: not implemented"
	// validate merger
	return *new(traits), nil
}

// Let merger payload format defaults to format

// TODO here is a hack for file source send interval. If it is sent in file, do not need to process sendInterval in decode

// pull source already pull internally which is also a rate limiter

// If rate limit merger is set, the first level decode will be done by merger

func checkByteSource(ss api.Source) model.NodeInfo {
	_ = "STUB: not implemented"
	return *new(model.NodeInfo)
}

func planLookupSource(ctx api.StreamContext, t *LookupPlan, ruleOption *def.RuleOption) (node.Emitter, error) {
	_ = "STUB: not implemented"
	return *new(node.Emitter), nil
}
