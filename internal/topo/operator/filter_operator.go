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

type FilterOp struct {
	Condition  ast.Expr
	StateFuncs []*ast.Call
}

// Apply the filter operator to each message in the stream
// The input data could be a xsql.Row or a xsql.Collection
// For xsql.Row, apply the condition to the row and return the row if the condition is true
// For xsql.Collection, apply the condition to each row and return the rows that meet the condition
// If error happens, return the error
func (p *FilterOp) Apply(ctx api.StreamContext, data interface{}, fv *xsql.FunctionValuer, _ *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// nil is false

// Only return if any row meets the condition, otherwise filter all
