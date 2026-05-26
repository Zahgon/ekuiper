// Copyright 2023-2025 EMQ Technologies Co., Ltd.
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

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type funcExecutor struct{}

func (f *funcExecutor) ValidateWithName(args []ast.Expr, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO pass in ctx

func (f *funcExecutor) Validate(_ []interface{}) error { _ = "STUB: not implemented"; return nil }

func (f *funcExecutor) Exec(ctx api.FunctionContext, args []any) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f *funcExecutor) ExecWithName(args []interface{}, ctx api.FunctionContext, name string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (f *funcExecutor) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (f *funcExecutor) GetFuncType(name string) ast.FuncType {
	_ = "STUB: not implemented"
	return *new(ast.FuncType)
}

var staticFuncExecutor = &funcExecutor{}
