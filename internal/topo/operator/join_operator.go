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

// JoinOp TODO join expr should only be the equal op between 2 streams like tb1.id = tb2.id
type JoinOp struct {
	From  *ast.Table
	Joins ast.Joins
}

// Apply JoinOp to join two streams. If running in continuous query, the inner join will always return empty result because there is only one stream data.
func (jp *JoinOp) Apply(ctx api.StreamContext, data interface{}, fv *xsql.FunctionValuer, _ *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (jp *JoinOp) getStreamNames(join *ast.Join) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exclude default stream as it is a virtual stream name.

func (jp *JoinOp) evalSet(ctx api.StreamContext, input xsql.Collection, join ast.Join, fv *xsql.FunctionValuer) (*xsql.JoinTuples, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If no messages in the right

func evalOn(join ast.Join, ve *xsql.ValuerEval, left interface{}, right xsql.Row) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// if no on expression

func (jp *JoinOp) evalSetWithRightJoin(input xsql.Collection, join ast.Join, excludeJoint bool, fv *xsql.FunctionValuer) (*xsql.JoinTuples, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jp *JoinOp) evalJoinSets(set *xsql.JoinTuples, input xsql.Collection, join ast.Join, fv *xsql.FunctionValuer) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jp *JoinOp) evalRightJoinSets(set *xsql.JoinTuples, input xsql.Collection, join ast.Join, excludeJoint bool, fv *xsql.FunctionValuer) (*xsql.JoinTuples, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
