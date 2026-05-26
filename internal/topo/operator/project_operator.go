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

package operator

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type ProjectOp struct {
	ColNames         [][]string // list of [col, table]
	ExceptNames      []string   // list of except name
	AllWildcard      bool
	WildcardEmitters map[string]bool
	AliasFields      ast.Fields
	ExprFields       ast.Fields
	Fields           ast.Fields
	// the length of fields exclude invisible
	FieldLen    int
	IsAggregate bool // Whether the project is used in an aggregate context. This is set by planner by analyzing the SQL query
	EnableLimit bool
	LimitCount  int
	// ExprIndices holds the SinkContent output slot index for each entry in
	// ExprFields when operating in slice-tuple mode.  Each ExprField (a
	// non-alias, non-FieldRef expression) gets a dedicated sink slot whose
	// position is pre-computed by the planner.
	ExprIndices []int

	SendMeta bool
	SendNil  bool

	kvs   []interface{}
	alias []interface{}
}

// Apply
//
//	input: *xsql.Tuple| xsql.Collection
//
// output: []map[string]interface{}
func (pp *ProjectOp) Apply(ctx api.StreamContext, data interface{}, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (pp *ProjectOp) getVE(tuple xsql.RawRow, agg xsql.AggregateData, wr *xsql.WindowRange, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer) *xsql.ValuerEval {
	_ = "STUB: not implemented"
	return nil
}

func (pp *ProjectOp) getRowVE(tuple xsql.Row, wr *xsql.WindowRange, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer) *xsql.ValuerEval {
	_ = "STUB: not implemented"
	return nil
}

func (pp *ProjectOp) project(row xsql.RawRow, ve *xsql.ValuerEval) error {
	_ = "STUB: not implemented"
	return nil
}

// set it to a typed nil to distinguish from nil
// so that the encoder can treat it differently from nil

// Evaluate expressions that are not plain field references (e.g. CASE WHEN,
// arithmetic, function calls without alias).  Their dedicated output slots
// are pre-assigned by the planner and stored in ExprIndices.

// ExprField: already handled in the ExprFields loop above.

// set it to a typed nil to distinguish from nil
// so that the encoder can treat it differently from nil

// Calculate all fields then pick the needed ones
// To make sure all calculations are run with the same context (e.g. alias values)
// Do not set value during calculations
