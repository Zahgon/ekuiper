// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

// FunctionValuer ONLY use NewFunctionValuer function to initialize
type FunctionValuer struct {
	runtime *funcRuntime
}

// Should only be called by stream to make sure a single instance for an operation
func NewFunctionValuer(p *funcRuntime) *FunctionValuer { _ = "STUB: not implemented"; return nil }

func (*FunctionValuer) Value(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (*FunctionValuer) Meta(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (*FunctionValuer) AppendAlias(string, interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (*FunctionValuer) AliasValue(string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (fv *FunctionValuer) Call(name string, funcId int, args []interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// do nothing, continue
