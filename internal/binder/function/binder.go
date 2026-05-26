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

package function

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/binder"
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

var ( // init once and read only
	funcFactories      []binder.FuncFactory
	funcFactoriesNames []string
)

func init() {
	f := binder.FactoryEntry{
		Name:    "built-in",
		Factory: GetManager(),
	}
	applyFactory(f)
}

// Initialize Only call once when server starts
func Initialize(factories []binder.FactoryEntry) error { _ = "STUB: not implemented"; return nil }

func applyFactory(f binder.FactoryEntry) { _ = "STUB: not implemented"; return }

func Function(name string) (api.Function, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), nil
}

func GetFunctionPlugin(name string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func HasFunctionSet(name string) bool { _ = "STUB: not implemented"; return false }

func ConvName(name string) (string, bool) { _ = "STUB: not implemented"; return "", false }

type multiAggFunc interface {
	GetFuncType(name string) ast.FuncType
}

func IsAggFunc(funcName string) bool { _ = "STUB: not implemented"; return false }

// NoAggFunc returns true if the function CANNOT be used in an aggregate query
func NoAggFunc(funcName string) bool { _ = "STUB: not implemented"; return false }

func GetFuncType(funcName string) ast.FuncType {
	_ = "STUB: not implemented"
	return *new(ast.FuncType)
}
