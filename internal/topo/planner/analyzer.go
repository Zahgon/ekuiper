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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type streamInfo struct {
	stmt   *ast.StreamStmt
	schema ast.StreamFields
}

// Analyze the select statement by decorating the info from stream statement.
// Typically, set the correct stream name for fieldRefs
func decorateStmt(s *ast.SelectStatement, opt *def.RuleOption, isTemp bool) ([]*streamInfo, []*ast.Call, []*ast.Call, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Validate temp streams can only be used by temp rules

// [fieldName][streamsName][*aliasRef] if alias, with special key alias/default. Each key has exactly one value

// unfold the wildcard for slice mode

// flat wildcard

// Scan columns fields: bind all field refs, collect alias

// bind alias field expressions

// Bind field ref for alias AND set StreamName for all field ref

// do not bind selection fields, should have done above

// check if stream exists

// Collect all analytic function calls so that we can let them run firstly

// walk sources at last to let them run firstly
// because another clause may depend on the alias defined here

type aliasTopoDegree struct {
	alias  string
	degree int
	field  ast.Field
}

type aliasTopoDegrees []*aliasTopoDegree

func (a aliasTopoDegrees) Len() int { _ = "STUB: not implemented"; return 0 }

func (a aliasTopoDegrees) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (a aliasTopoDegrees) Swap(i, j int) { _ = "STUB: not implemented"; return }

// checkAliasReferenceCycle checks whether exists select a + 1 as b, b + 1 as a from demo;
func checkAliasReferenceCycle(s *ast.SelectStatement) bool { _ = "STUB: not implemented"; return false }

func dfsRef(aliasRef map[string]map[string]struct{}, walked map[string]struct{}, currentName, targetName string) bool {
	_ = "STUB: not implemented"
	return false
}

func aliasFieldTopoSort(s *ast.SelectStatement, streamStmts []*streamInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// the unknownFieldRef name belongs to a alias

func isFieldRefNameExists(name string, streamStmts []*streamInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func isAliasFieldTopoSortFinish(aliasDegrees map[string]*aliasTopoDegree) bool {
	_ = "STUB: not implemented"
	return false
}

type validateOptStmt interface {
	validate(statement *ast.SelectStatement) error
}

func validate(stmt *ast.SelectStatement) error { _ = "STUB: not implemented"; return nil }

var stmtCheckers = []validateOptStmt{
	&aggFuncChecker{},
	&groupChecker{},
}

type aggFuncChecker struct{}

func (c *aggFuncChecker) validate(s *ast.SelectStatement) (err error) {
	_ = "STUB: not implemented"

	// lazy set isAgg flag
	return nil
}

// aggregate call should not have any aggregate arg

// agg func check is done in dimensions.
// in window trigger condition, NoAggFunc is allowed unlike normal condition so return false to skip that check

type groupChecker struct{}

func (c *groupChecker) validate(s *ast.SelectStatement) error {
	_ = "STUB: not implemented"
	return nil
}

// file-private functions below
// allAggregate checks if all expressions of binary expression are aggregate
func allAggregate(expr ast.Expr) (r bool) { _ = "STUB: not implemented"; return false }

// do nothing

func convertStreamInfo(streamStmt *ast.StreamStmt) (*streamInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone the statement to avoid data race when multiple rules sharing the same stream definition
// The planner may modify the options, e.g. set the type to default value

type fieldsMap struct {
	content       map[string]streamFieldStore
	aliasNames    map[string]struct{}
	isSchemaless  bool
	defaultStream ast.StreamName
}

func newFieldsMap(isSchemaless bool, defaultStream ast.StreamName) *fieldsMap {
	_ = "STUB: not implemented"
	return nil
}

func (f *fieldsMap) reserve(fieldName string, streamName ast.StreamName) {
	_ = "STUB: not implemented"
	return
}

func (f *fieldsMap) save(fieldName string, streamName ast.StreamName, field *ast.AliasRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fieldsMap) bindAlias(aliasName string) { _ = "STUB: not implemented"; return }

func (f *fieldsMap) bind(fr *ast.FieldRef) error { _ = "STUB: not implemented"; return nil }

type streamFieldStore interface {
	add(k ast.StreamName)
	ref(k ast.StreamName, v *ast.AliasRef) error
	bindRef(f *ast.FieldRef) error
}

func newStreamFieldStore(isSchemaless bool, defaultStream ast.StreamName) streamFieldStore {
	_ = "STUB: not implemented"
	return *new(streamFieldStore)
}

type streamFieldMap struct {
	content map[ast.StreamName]*ast.AliasRef
}

// add the stream name must not be default.
// This is used when traversing stream schema
func (s *streamFieldMap) add(k ast.StreamName) {
	_ = "STUB: not implemented"

	// bind for schema field, all keys must be created before running bind
	// can bind alias & col. For alias, the stream name must be empty; For col, the field must be a col
	return
}

func (s *streamFieldMap) ref(k ast.StreamName, v *ast.AliasRef) error {
	_ = "STUB: not implemented"
	return nil
	// must not exist, save alias ref for alias
}

// the key must exist after the schema travers, do validation
// In schema mode, default stream won't be a key

// valid, do nothing

func (s *streamFieldMap) bindRef(fr *ast.FieldRef) error { _ = "STUB: not implemented"; return nil }

// if alias, return this

// if alias exists

type streamFieldMapSchemaless struct {
	content       map[ast.StreamName]*ast.AliasRef
	defaultStream ast.StreamName
}

// add this should not be called for schemaless
func (s *streamFieldMapSchemaless) add(k ast.StreamName) {
	_ = "STUB: not implemented"

	// bind for schemaless field, create column if not exist
	// can bind alias & col. For alias, the stream name must be empty; For col, the field must be a col
	return
}

func (s *streamFieldMapSchemaless) ref(k ast.StreamName, v *ast.AliasRef) error {
	_ = "STUB: not implemented"
	return nil
	// must not exist
}

// the key may or may not exist. But always have only one default stream field.
// Replace with stream name if another stream found. The key can be duplicate

// In schemaless mode, default stream can only exist when length is 1

// valid, do nothing

func (s *streamFieldMapSchemaless) bindRef(fr *ast.FieldRef) error {
	_ = "STUB: not implemented"
	return nil
}

// must be a column because alias are fields and have been traversed
// reserve a hole and do nothing

// if alias or single col, return this

// if alias exists

// reserver a hole
