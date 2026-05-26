// Copyright 2021-2023 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

// IsAggregate check if an expression is aggregate with the binding alias info
func IsAggregate(expr ast.Expr) (r bool) { _ = "STUB: not implemented"; return false }

// lazy calculate

func getOrCalculateAgg(f *ast.FieldRef) bool { _ = "STUB: not implemented"; return false }

func WithAggFields(stmt *ast.SelectStatement) bool { _ = "STUB: not implemented"; return false }

func HasAggFuncs(node ast.Node) bool { _ = "STUB: not implemented"; return false }
