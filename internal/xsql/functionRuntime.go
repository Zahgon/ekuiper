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

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// Manage the function plugin instances
// Each operator has a single instance of this to hold the context
type funcRuntime struct {
	syncx.Mutex
	regs      []*funcReg
	parentCtx api.StreamContext
}

type funcReg struct {
	ins api.Function
	ctx api.FunctionContext
}

func NewFuncRuntime(ctx api.StreamContext) *funcRuntime { _ = "STUB: not implemented"; return nil }

// Get Each funcId returns a single instance of the function
// The funcId is assigned in operator instance level, thus each operator will have a single instance of the function
func (fp *funcRuntime) Get(name string, funcId int) (api.Function, api.FunctionContext, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), *new(api.FunctionContext), nil
}

// Check service extension and plugin extension if set
