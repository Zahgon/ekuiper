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

package ast

import (
	"regexp"
)

type LikePattern struct {
	Expr    Expr
	Pattern *regexp.Regexp
}

func (l *LikePattern) expr()          { _ = "STUB: not implemented"; return }
func (l *LikePattern) node()          { _ = "STUB: not implemented"; return }
func (l *LikePattern) String() string { _ = "STUB: not implemented"; return "" }

func (l *LikePattern) Compile(likestr string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FieldRef could be
//  1. SQL Field
//     1.1 Explicit field "stream.col"
//     1.2 Implicit field "col"  -> only exist in schemaless stream. Otherwise, explicit stream name will be bound
//     1.3 Alias field "expr as c" -> refer to an Expression or column
type FieldRef struct {
	// optional, bind in analyzer, empty means alias, default means not set
	// MUST have after binding for SQL fields. For 1.2,1.3 and 1.4, use special constant as stream name
	StreamName StreamName
	// optional, set only once. For selections, empty name will be assigned a default name
	// MUST have after binding, assign a name for 1.4
	Name string
	// whether the index is set. Use this to distinguish uninitialized index and zero index
	HasIndex bool
	// optional, set when it is a Field define in stream. If not a field, set to -1
	SourceIndex int
	// must have if it is from sink.
	Index int
	// Only for alias
	*AliasRef
}

func (fr *FieldRef) expr()          { _ = "STUB: not implemented"; return }
func (fr *FieldRef) node()          { _ = "STUB: not implemented"; return }
func (fr *FieldRef) IsColumn() bool { _ = "STUB: not implemented"; return false }

func (fr *FieldRef) String() string { _ = "STUB: not implemented"; return "" }

func (fr *FieldRef) IsAlias() bool { _ = "STUB: not implemented"; return false }

func (fr *FieldRef) RefSelection(a *AliasRef) {
	_ = "STUB: not implemented"

	// RefSources Must call after binding or will get empty
	return
}

func (fr *FieldRef) RefSources() []StreamName { _ = "STUB: not implemented"; return nil }

// SetRefSource Only call this for alias field ref
func (fr *FieldRef) SetRefSource(names []StreamName) { _ = "STUB: not implemented"; return }

type AliasRef struct {
	// MUST have, It is used for evaluation
	Expression Expr
	// MUST have after binding, calculate once in initializer. Could be 0 when alias an Expression without col like "1+2"
	RefSources []StreamName
	// optional, lazy set when calculating isAggregate
	IsAggregate *bool
}

func (a *AliasRef) String() string { _ = "STUB: not implemented"; return "" }

// SetRefSource only used for unit test
func (a *AliasRef) SetRefSource(names []string) { _ = "STUB: not implemented"; return }

func NewAliasRef(e Expr) (*AliasRef, error) { _ = "STUB: not implemented"; return nil, nil }

// MockAliasRef is for testing only.
func MockAliasRef(e Expr, r []StreamName, a *bool) *AliasRef { _ = "STUB: not implemented"; return nil }
