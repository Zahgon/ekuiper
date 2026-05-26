// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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

type WindowFuncOperator struct {
	WindowFuncField *ast.Field
}

type windowFuncHandle interface {
	handleTuple(input xsql.Row) xsql.Row
	handleCollection(input xsql.Collection) xsql.Collection
}

type rowNumberFuncHandle struct {
	name string
}

func (rh *rowNumberFuncHandle) handleTuple(input xsql.Row) xsql.Row {
	_ = "STUB: not implemented"
	return *new(xsql.Row)
}

func (rh *rowNumberFuncHandle) handleCollection(input xsql.Collection) xsql.Collection {
	_ = "STUB: not implemented"
	return *new(xsql.Collection)
}

func (wf *WindowFuncOperator) Apply(ctx api.StreamContext, data interface{}, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// handle the following case:
// 1: row_number() over (partition by a)
// 2: row_number() over (partition by a order by b)

// handle the following case:
// 1: row_number() over (order by a)

// handle the following case:
// 1: row_number() without over clause

func getWindowFuncHandle(funcName, colName string) (windowFuncHandle, error) {
	_ = "STUB: not implemented"
	return *new(windowFuncHandle), nil
}

func sortCollection(ctx api.StreamContext, data xsql.Collection, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer, sortFields ast.SortFields) xsql.Collection {
	_ = "STUB: not implemented"
	return *new(xsql.Collection)
}

func partitionCollection(ctx api.StreamContext, input xsql.Collection, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer, prs *ast.PartitionExpr, sortFields ast.SortFields, wh windowFuncHandle) (xsql.Collection, error) {
	_ = "STUB: not implemented"
	return *new(xsql.Collection), nil
}

// visit result by order
