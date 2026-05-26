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

type AnalyticFuncsOp struct {
	Funcs       []*ast.Call
	FieldFuncs  []*ast.Call
	transformed bool
}

func (p *AnalyticFuncsOp) evalTupleFunc(calls []*ast.Call, ve *xsql.ValuerEval, input xsql.Row) (xsql.Row, error) {
	_ = "STUB: not implemented"
	return *new(xsql.Row), nil
}

func (p *AnalyticFuncsOp) evalCollectionFunc(calls []*ast.Call, fv *xsql.FunctionValuer, input xsql.Collection) (xsql.Collection, error) {
	_ = "STUB: not implemented"
	return *new(xsql.Collection), nil
}

func (p *AnalyticFuncsOp) Apply(ctx api.StreamContext, data interface{}, fv *xsql.FunctionValuer, _ *xsql.AggregateFunctionValuer) (got interface{}) {
	_ = "STUB: not implemented"
	return nil
}
