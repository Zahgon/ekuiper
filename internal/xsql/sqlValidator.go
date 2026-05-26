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

// Validate select statement without context.
// This is the pre-validation. In planner, there will be a more comprehensive validation after binding
func Validate(stmt *ast.SelectStatement) error { _ = "STUB: not implemented"; return nil }

func validateWindowFunction(stmt *ast.SelectStatement) error { _ = "STUB: not implemented"; return nil }

func validateSRFNestedForbidden(clause string, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMultiSRFForbidden(clause string, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSRFForbidden(node ast.Node) error { _ = "STUB: not implemented"; return nil }

func isSRFNested(node ast.Node) bool { _ = "STUB: not implemented"; return false }

func isWindowFunctionExists(node ast.Node) bool { _ = "STUB: not implemented"; return false }

// skip checking Fields
// TODO: support window functions in order by clause lately

func isSRFExists(node ast.Node) bool { _ = "STUB: not implemented"; return false }

// skip checking Fields

func validateFields(stmt *ast.SelectStatement, streamNames []string) {
	_ = "STUB: not implemented"
	return
}

// validateExpr checks if the streamName of a fieldRef is existed and covert it to json filed if not exist.
// The expr is the expression to be validated, and streamName is the stream name of the current select statement.
// The expr only contains the expression which is possible to be used in fields and join conditions
func validateExpr(expr ast.Expr, streamName []string) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

// Checks whether a slice contains an element
func contains(s []string, n string) bool { _ = "STUB: not implemented"; return false }

func getStreamNames(stmt *ast.SelectStatement) (result []string) {
	_ = "STUB: not implemented"
	return nil
}
