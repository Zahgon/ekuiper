// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package js

import (
	"github.com/dop251/goja"
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

// JSFunc is stateful
// Each instance has its own vm
type JSFunc struct {
	vm     *goja.Runtime
	jsfunc goja.Callable
	isAgg  bool
	// state, use this to avoid creating new array each time
	args []goja.Value
}

func NewJSFunc(symbolName string) (*JSFunc, error) { _ = "STUB: not implemented"; return nil, nil }

// Get the text from the symbol table

// Should not happen, already verify when install
//if err != nil {
//	return nil, fmt.Errorf("failed to interpret script: %v", err)
//}

// Should not happen, already verify when install
//if !ok {
//	return nil, fmt.Errorf("cannot find function \"%s\" in script", symbolName)
//}

func (f *JSFunc) Validate(_ []interface{}) error { _ = "STUB: not implemented"; return nil }

func (f *JSFunc) Exec(ctx api.FunctionContext, args []any) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f *JSFunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (f *JSFunc) Close() error { _ = "STUB: not implemented"; return nil }
