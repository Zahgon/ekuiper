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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

// DecodeOp manages the format decoding (employ schema) and sending frequency (for batch decode, like a json array)
type DecodeOp struct {
	*defaultSinkNode
	converter message.Converter

	c *dconf
	// This is for first level decode, add the payload field to schema to make sure it is decoded
	forPayload     bool
	additionSchema string
	// hint for map allocation
	hint int
}

type dconf struct {
	// When receiving list, send them one by one, this is the sending interval between each
	// Typically set by file source
	SendInterval      cast.DurationConf `json:"sendInterval"`
	Format            string            `json:"format"`
	SchemaId          string            `json:"schemaId"`
	PayloadField      string            `json:"payloadField"`
	PayloadBatchField string            `json:"payloadBatchField"`
	PayloadFormat     string            `json:"payloadFormat"`
	PayloadSchemaId   string            `json:"payloadSchemaId"`
	PayloadDelimiter  string            `json:"payloadDelimiter"`
}

func NewDecodeOp(ctx api.StreamContext, forPayload bool, name string, rOpt *def.RuleOption, schema map[string]*ast.JsonStreamField, props map[string]any) (*DecodeOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// It is payload decoder

// Exec decode op receives raw data and converts it to message
func (o *DecodeOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (o *DecodeOp) Worker(ctx api.StreamContext, item any) []any {
	_ = "STUB: not implemented"
	return nil
}

func (o *DecodeOp) ResetSchema(ctx api.StreamContext, schema map[string]*ast.JsonStreamField) {
	_ = "STUB: not implemented"
	return
}

// append payload field to schema

// PayloadDecodeWorker each input has one message with the payload field to decode
//
//	{
//		"payloadField":"data","otherField":1
//	}
//
//	{
//		// parsed fields
//		"parsedField": 1,
//		"parsedField2": 2,
//		// keep the original field if in schema
//		"payloadField":"data",
//		"otherField":1
//	}
//
// If parse result is a list, it will also output a list
func (o *DecodeOp) PayloadDecodeWorker(ctx api.StreamContext, item any) []any {
	_ = "STUB: not implemented"
	return nil
}

// extract payload

// TODO do not update the tuple directly
// Currently, this op is generated implicitly and it is guarantee to not share the data, so we mutate it directly
func transTuple(d *xsql.Tuple, result any) []any { _ = "STUB: not implemented"; return nil }

// PayloadBatchDecodeWorker deals with payload like
//
//	{
//		"ts": 123456,
//		"batchField": [
//			{"payloadField":"data","otherField":1},
//			{"payloadField":"data2","otherField":2}
//		]
//	}
//
// It merges all payload result into one
//
//	{
//		"ts": 123456,
//		// parsed fields are merged
//		"parsedField": 1,
//		"parsedField": 2,
//		// other fields also merged and keep the latest
//		"otherField": 2
//	}
//
// If parse result is a list, it will also merge them in
func (o *DecodeOp) PayloadBatchDecodeWorker(ctx api.StreamContext, item any) []any {
	_ = "STUB: not implemented"
	return nil
}

// extract batch field

// TODO do not update the tuple directly
// Currently, this op is generated implicitly and it is guarantee to not share the data, so we mutate it directly
func mergeTuple(ctx api.StreamContext, d *xsql.Tuple, result any) {
	_ = "STUB: not implemented"
	return
}

func toTupleFromRawTuple(ctx api.StreamContext, v map[string]any, d *xsql.RawTuple) *xsql.Tuple {
	_ = "STUB: not implemented"
	return nil
}

func cloneTuple(d *xsql.Tuple, hint int) *xsql.Tuple { _ = "STUB: not implemented"; return nil }

func tupleAppend(d *xsql.Tuple, mv map[string]any) { _ = "STUB: not implemented"; return }

var _ SchemaNode = &DecodeOp{}
