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

type (
	funcExe      func(ctx api.FunctionContext, args []interface{}) (interface{}, bool)
	funcVal      func(ctx api.FunctionContext, args []ast.Expr) error
	funcCheckNil func(args []interface{}) (interface{}, bool)
)

type builtinFunc struct {
	fType ast.FuncType
	exec  funcExe
	val   funcVal
	check funcCheckNil
}

var (
	builtins            map[string]builtinFunc
	builtinStatfulFuncs map[string]func() api.Function
	kvPairKName         = "key"
	kvPairVName         = "value"
)

func init() {
	builtins = make(map[string]builtinFunc)
	builtinStatfulFuncs = make(map[string]func() api.Function)
	registerAggFunc()
	registerIncAggFunc()
	registerMathFunc()
	registerStrFunc()
	registerMiscFunc()
	registerAnalyticFunc()
	registerColsFunc()
	registerSetReturningFunc()
	registerArrayFunc()
	registerObjectFunc()
	registerGlobalStateFunc()
	registerDateTimeFunc()
	registerGlobalAggFunc()
	registerWindowFunc()
}

//var funcWithAsteriskSupportMap = map[string]string{
//	"collect": "",
//	"count":   "",
//}

var analyticFuncs = map[string]struct{}{
	"lag":         {},
	"changed_col": {},
	"had_changed": {},
	"latest":      {},
	"acc_sum":     {},
	"acc_min":     {},
	"acc_max":     {},
	"acc_avg":     {},
	"acc_count":   {},
}

var windowFuncs = map[string]struct{}{
	"row_number": {},
}

const AnalyticPrefix = "$$a"

func IsWindowFunc(name string) bool { _ = "STUB: not implemented"; return false }

func IsAnalyticFunc(name string) bool { _ = "STUB: not implemented"; return false }

type Manager struct{}

// Function the name is converted to lowercase if needed during parsing
func (m *Manager) Function(name string) (api.Function, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), nil
}

func (m *Manager) FunctionPluginInfo(funcName string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func (m *Manager) HasFunctionSet(name string) bool { _ = "STUB: not implemented"; return false }

func (m *Manager) ConvName(n string) (string, bool) { _ = "STUB: not implemented"; return "", false }

var (
	m                    = &Manager{}
	_ binder.FuncFactory = m
)

func GetManager() *Manager { _ = "STUB: not implemented"; return nil }

func returnNilIfHasAnyNil(args []interface{}) (returned interface{}, skipExec bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func returnFalseIfHasAnyNil(args []interface{}) (returned interface{}, skipExec bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func return0IfHasAnyNil(args []interface{}) (returned interface{}, skipExec bool) {
	_ = "STUB: not implemented"
	return nil, false
}
