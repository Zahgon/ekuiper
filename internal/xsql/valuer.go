// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

var (
	// implicitValueFuncs is a set of functions that event implicitly passes the value.
	implicitValueFuncs = map[string]bool{
		"window_start":   true,
		"window_end":     true,
		"event_time":     true,
		"window_trigger": true,
	}
	// ImplicitStateFuncs is a set of functions that read/update global state implicitly.
	ImplicitStateFuncs = map[string]bool{
		"last_hit_time":      true,
		"last_hit_count":     true,
		"last_agg_hit_time":  true,
		"last_agg_hit_count": true,
	}
)

/*
 *  Valuer definitions
 */

// Valuer is the interface that wraps the Value() method.
type Valuer interface {
	// Value returns the value and existence flag for a given key.
	Value(key, table string) (interface{}, bool)
	Meta(key, table string) (interface{}, bool)
}

// AliasValuer is used to calculate and cache the alias value
type AliasValuer interface {
	// AliasValue Get the value of alias
	AliasValue(name string) (interface{}, bool)
	// AppendAlias set the alias result
	AppendAlias(key string, value interface{}) bool
}

// CallValuer implements the Call method for evaluating function calls.
type CallValuer interface {
	Valuer

	// Call is invoked to evaluate a function call (if possible).
	Call(name string, funcId int, args []interface{}) (interface{}, bool)
}

// FuncValuer can calculate function type value like window_start and window_end
type FuncValuer interface {
	FuncValue(key string) (interface{}, bool)
}

type AggregateCallValuer interface {
	CallValuer
	GetAllTuples() AggregateData
	GetSingleCallValuer() CallValuer
}

type WildcardValuer struct {
	Data Wildcarder
}

func (wv *WildcardValuer) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (wv *WildcardValuer) Meta(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"

	// MultiValuer returns a Valuer that iterates over multiple Valuer instances
	// to find a match.
	return nil, false
}

func MultiValuer(valuers ...Valuer) Valuer { _ = "STUB: not implemented"; return *new(Valuer) }

type multiValuer []Valuer

func (a multiValuer) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a multiValuer) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a multiValuer) AppendAlias(key string, value interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (a multiValuer) AliasValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a multiValuer) FuncValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a multiValuer) Call(name string, funcId int, args []interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a multiValuer) ValueByIndex(index int, sourceIndex int) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (a multiValuer) SetByIndex(index int, value any) { _ = "STUB: not implemented"; return }

func (a multiValuer) TempByIndex(index int) any { _ = "STUB: not implemented"; return *new(any) }

func (a multiValuer) SetTempByIndex(index int, value any) { _ = "STUB: not implemented"; return }

type multiAggregateValuer struct {
	data AggregateData
	multiValuer
	singleCallValuer CallValuer
}

func MultiAggregateValuer(data AggregateData, singleCallValuer CallValuer, valuers ...Valuer) Valuer {
	_ = "STUB: not implemented"
	return *new(Valuer)
}

func (a *multiAggregateValuer) Call(name string, funcId int, args []interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	// assume the aggFuncMap already cache the custom agg funcs in IsAggFunc()
	return nil, false
}

func (a *multiAggregateValuer) GetAllTuples() AggregateData {
	_ = "STUB: not implemented"
	return *new(AggregateData)
}

func (a *multiAggregateValuer) GetSingleCallValuer() CallValuer {
	_ = "STUB: not implemented"
	return *new(CallValuer)
}

func (a *multiAggregateValuer) AppendAlias(key string, value interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *multiAggregateValuer) AliasValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

/*
 * Eval Logics
 */

// Eval evaluates expr against a map.
func Eval(expr ast.Expr, m Valuer) interface{} { _ = "STUB: not implemented"; return nil }

// ValuerEval will evaluate an expression using the Valuer.
type ValuerEval struct {
	Valuer Valuer

	// IntegerFloatDivision will set the eval system to treat
	// a division between two integers as a floating point division.
	IntegerFloatDivision bool
}

// Eval evaluates an expression and returns a value.
// map the expression to the correct valuer
func (v *ValuerEval) Eval(expr ast.Expr) interface{} { _ = "STUB: not implemented"; return nil }

// The analytic functions are calculated prior to all ops, so just get the cached field value

// nil is also cached

// This is the implicit arg set by the filter planner
// If set, it will only return the value, no updating the value

// In the parser, the col func arguments must be ColField

// won't happen

// won't happen

// this data should be recorded or not ? default answer is yes

// analytic func must put the partition key into the args

// TODO possible performance elevation to eliminate this cal

// The field specified with stream source

func (v *ValuerEval) evalBinaryExpr(expr *ast.BinaryExpr) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// shortcut for bool

func (v *ValuerEval) evalCase(expr *ast.CaseExpr) interface{} {
	_ = "STUB: not implemented"
	return nil
	// compare value to all when clause
}

func (v *ValuerEval) evalValueSet(expr *ast.ValueSetExpr) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *ValuerEval) evalSetsExpr(lhs interface{}, op ast.Token, rhsSet interface{}) interface{} {
	_ = "STUB: not implemented"

	/*Semantic rules

	When using the IN operator, the following semantics apply in this order:

	Returns FALSE if value_set is empty.
	Returns NULL if search_value is NULL.
	Returns TRUE if value_set contains a value equal to search_value.
	Returns NULL if value_set contains a NULL.
	Returns FALSE.
	When using the NOT IN operator, the following semantics apply in this order:

	Returns TRUE if value_set is empty.
	Returns NULL if search_value is NULL.
	Returns FALSE if value_set contains a value equal to search_value.
	Returns NULL if value_set contains a NULL.
	Returns TRUE.
	*/return nil
}

func (v *ValuerEval) evalJsonExpr(result interface{}, op ast.Token, expr ast.Expr) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (v *ValuerEval) subset(result interface{}, expr ast.Expr) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// SimpleDataEval lhs and rhs are non-nil
func (v *ValuerEval) SimpleDataEval(lhs, rhs any, op ast.Token) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// for relationship, return false

// Evaluate if both sides are simple types.

// Try the rhs as a float64, int64, or uint64

// Try as a float64 to see if a float cast is required.

// Try as a float64 to see if a float cast is required.

/*
 * Helper functions
 */

type BracketEvalResult struct {
	Start, End int
}

func (ber *BracketEvalResult) isIndex() bool { _ = "STUB: not implemented"; return false }

func isSliceOrArray(v interface{}) bool { _ = "STUB: not implemented"; return false }

func isSetOperator(op ast.Token) bool { _ = "STUB: not implemented"; return false }

func invalidOpError(lhs interface{}, op ast.Token, rhs interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func convertNum(para interface{}) interface{} {
	_ = "STUB: not implemented"

	// Already check type of para so that there will be no error, just ignore error
	return nil
}

func isInt(para interface{}) bool { _ = "STUB: not implemented"; return false }

func isFloat(para interface{}) bool { _ = "STUB: not implemented"; return false }
