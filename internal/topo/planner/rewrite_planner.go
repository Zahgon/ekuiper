// Copyright 2025 EMQ Technologies Co., Ltd.
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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type rewriteResult struct {
	windowFuncFields     []*ast.Field
	incAggFields         []*ast.Field
	dsColAliasMapping    map[ast.StreamName]map[string]string
	aggFuncsFieldInWhere []*ast.Field
}

// rewrite stmt will do following things:
// 1. extract and rewrite the window function
// 2. extract and rewrite the aggregation function
func rewriteStmt(stmt *ast.SelectStatement, opt *def.RuleOption) rewriteResult {
	_ = "STUB: not implemented"
	return *new(rewriteResult)
}

// extract agg function from filter condition and rewrite with bypass fields
func rewriteAggFunctionInWhere(stmt *ast.SelectStatement, _ *def.RuleOption) []*ast.Field {
	_ = "STUB: not implemented"
	return nil
}

func RewriteAggFunctionInWhere(stmt *ast.SelectStatement) []*ast.Field {
	_ = "STUB: not implemented"
	return nil
}
