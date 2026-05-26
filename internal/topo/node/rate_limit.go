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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type limitConf struct {
	Interval        cast.DurationConf `json:"interval"`
	MergeField      string            `json:"mergeField"`
	Format          string            `json:"format"`
	Merger          string            `json:"merger"`
	PayloadSchemaId string            `json:"payloadSchemaId"`
	SchemaId        string            `json:"schemaId"`
}

// RateLimitOp handle messages at a regular rate, ignoring messages that arrive too quickly, only keep the most recent message. (default strategy)
// If strategy is set, send through all messages as well as trigger signal and let strategy node handle the merge.
// Otherwise, send the most recent message at trigger time
// Input: Raw
// Output: Raw as it is
// Concurrency: false
type RateLimitOp struct {
	*defaultSinkNode
	// configs
	c             *limitConf
	mergeStrategy int
	// state
	// keep last strategy
	latest any
	// merged items
	frameSet map[any]map[string]any
	// only when mergeField is set
	decoder message.PartialDecoder
	// only when using merger
	merger modules.Merger
}

func NewRateLimitOp(ctx api.StreamContext, name string, rOpt *def.RuleOption, schema map[string]*ast.JsonStreamField, props map[string]any) (*RateLimitOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec ratelimit op deal with 3 merge strategy
// - latest
// - merge by mergeField (when format and mergeField is set and no payload format)
// - merge by merger (when format, payloadFormat and merger is set)
func (o *RateLimitOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// get the latest

// by mergeField

// sort it

// end the last span after receiving a new item

// do not use onProcessEnd because span still need to exist

func (o *RateLimitOp) ResetSchema(ctx api.StreamContext, schema map[string]*ast.JsonStreamField) {
	_ = "STUB: not implemented"
	return
}
