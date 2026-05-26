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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/lookup/cache"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

type LookupConf struct {
	Cache           bool              `json:"cache"`
	CacheTTL        cast.DurationConf `json:"cacheTtl"`
	CacheMissingKey bool              `json:"cacheMissingKey"`
}

type srcConf struct {
	PayloadField     string `json:"payloadField"`
	PayloadFormat    string `json:"payloadFormat"`
	PayloadSchemaId  string `json:"payloadSchemaId"`
	PayloadDelimiter string `json:"payloadDelimiter"`
}

// LookupNode will look up the data from the external source when receiving an event
type LookupNode struct {
	*defaultSinkNode

	conf *LookupConf
	c    *srcConf

	joinType ast.JoinType
	vals     []ast.Expr
	fields   []string
	keys     []string
	// If lookupByteSource, the decoders are needed
	isBytesLookup  bool
	formatDecoder  message.Converter
	payloadDecoder message.Converter
}

func NewLookupNode(ctx api.StreamContext, name string, isBytesLookup bool, fields []string, keys []string, joinType ast.JoinType, vals []ast.Expr, srcOptions *ast.Options, options *def.RuleOption, props map[string]any) (*LookupNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *LookupNode) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// Start the lookup source loop

// process incoming item from both streams(transformed) and tables

// lookup will lookup the cache firstly, if expires, read the external source
func (n *LookupNode) lookup(ctx api.StreamContext, d xsql.Row, fv *xsql.FunctionValuer, ns api.Source, tuples *xsql.JoinTuples, c *cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

// if any of the value is nil, the lookup will always return empty result

func (n *LookupNode) doLookup(ctx api.StreamContext, ns api.Source, cvs []any) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only called when isBytesLookup is true
// Must guarantee decoders are set
func (n *LookupNode) decode(ctx api.StreamContext, row []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func decodePayload(ctx api.StreamContext, decoder message.Converter, rt map[string]any, field string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
