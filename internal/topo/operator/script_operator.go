// Copyright 2023 EMQ Technologies Co., Ltd.
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

//go:build script

package operator

import (
	"github.com/dop251/goja"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

type ScriptOp struct {
	vm     *goja.Runtime
	jsfunc goja.Callable
	isAgg  bool
}

func NewScriptOp(script string, isAgg bool) (*ScriptOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ScriptOp) Apply(ctx api.StreamContext, data interface{}, _ *xsql.FunctionValuer, _ *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}
