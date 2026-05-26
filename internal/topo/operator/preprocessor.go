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

package operator

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

// Preprocessor only planned when
// 1. eventTime, to convert the timestamp field
// 2. schema validate and convert, when strict_validation is on and field type is not binary
// Do not convert types
type Preprocessor struct {
	// Pruned stream fields. Could be streamField(with data type info) or string
	defaultFieldProcessor
	// allMeta        bool
	// metaFields     []string //only needed if not allMeta
	isEventTime    bool
	timestampField string
	checkSchema    bool
	isBinary       bool
}

func NewPreprocessor(isSchemaless bool, fields map[string]*ast.JsonStreamField, _ bool, _ []string, iet bool, timestampField string, timestampFormat string, isBinary bool, strictValidation bool) (*Preprocessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply the preprocessor to the tuple
/*	input: *xsql.Tuple
 *	output: *xsql.Tuple
 */
func (p *Preprocessor) Apply(ctx api.StreamContext, data interface{}, _ *xsql.FunctionValuer, _ *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// No need to reconstruct meta as the memory has been allocated earlier
//if !p.allMeta && p.metaFields != nil && len(p.metaFields) > 0 {
//	newMeta := make(xsql.Metadata)
//	for _, f := range p.metaFields {
//		if m, ok := tuple.Metadata.Value(f, ""); ok {
//			newMeta[f] = m
//		}
//	}
//	tuple.Metadata = newMeta
//}
