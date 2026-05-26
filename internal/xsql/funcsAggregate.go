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

package xsql

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type AggregateFunctionValuer struct {
	data AggregateData
	fv   *FunctionValuer
}

func NewFunctionValuersForOp(ctx api.StreamContext) (*FunctionValuer, *AggregateFunctionValuer) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Should only be called by stream to make sure a single instance for an operation
func NewAggregateFunctionValuers(p *funcRuntime) (*FunctionValuer, *AggregateFunctionValuer) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *AggregateFunctionValuer) SetData(data AggregateData) { _ = "STUB: not implemented"; return }

func (v *AggregateFunctionValuer) GetSingleCallValuer() CallValuer {
	_ = "STUB: not implemented"
	return *new(CallValuer)
}

func (v *AggregateFunctionValuer) Value(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (v *AggregateFunctionValuer) Meta(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (v *AggregateFunctionValuer) FuncValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (*AggregateFunctionValuer) AppendAlias(string, interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *AggregateFunctionValuer) AliasValue(_ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (v *AggregateFunctionValuer) Call(name string, funcId int, args []interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// do nothing, continue

func (v *AggregateFunctionValuer) GetAllTuples() AggregateData {
	_ = "STUB: not implemented"
	return *new(AggregateData)
}
