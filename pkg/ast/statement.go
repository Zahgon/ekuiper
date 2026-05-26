// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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

type Statement interface {
	stmt()
	Node
}

type SelectStatement struct {
	Fields     Fields
	Sources    Sources
	Joins      Joins
	Condition  Expr
	Limit      Expr
	Dimensions Dimensions
	Having     Expr
	SortFields SortFields

	Statement
}

type Fields []Field

func (f Fields) node() { _ = "STUB: not implemented"; return }

func (f Fields) Len() int { _ = "STUB: not implemented"; return 0 }

func (f Fields) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (f Fields) Less(i int, j int) bool { _ = "STUB: not implemented"; return false }

type Field struct {
	Name      string
	AName     string
	Expr      Expr
	Invisible bool
	Node
}

func (f *Field) GetName() string { _ = "STUB: not implemented"; return "" }

func (f *Field) IsSelectionField() bool { _ = "STUB: not implemented"; return false }

func (f *Field) IsColumn() bool { _ = "STUB: not implemented"; return false }

type Sources []Source

func (s Sources) node() { _ = "STUB: not implemented"; return }

type Source interface {
	Node
	source()
}

type Table struct {
	Name  string
	Alias string
	Source
}

type JoinType int

const (
	LEFT_JOIN JoinType = iota
	INNER_JOIN
	RIGHT_JOIN
	FULL_JOIN
	CROSS_JOIN
)

func (j JoinType) String() string { _ = "STUB: not implemented"; return "" }

type Join struct {
	Name     string
	Alias    string
	JoinType JoinType
	Expr     Expr

	Node
}

type Joins []Join

func (j Joins) node() { _ = "STUB: not implemented"; return }

type Dimension struct {
	Expr Expr

	Node
}

type Dimensions []Dimension

func (d Dimensions) node() { _ = "STUB: not implemented"; return }

func (d *Dimensions) GetWindow() *Window { _ = "STUB: not implemented"; return nil }

func (d *Dimensions) GetGroups() Dimensions { _ = "STUB: not implemented"; return *new(Dimensions) }

type WindowType int

const (
	NOT_WINDOW WindowType = iota
	TUMBLING_WINDOW
	HOPPING_WINDOW
	SLIDING_WINDOW
	SESSION_WINDOW
	COUNT_WINDOW
	STATE_WINDOW
)

func (w WindowType) String() string { _ = "STUB: not implemented"; return "" }

type Window struct {
	PartitionExpr    *PartitionExpr
	TriggerCondition Expr
	SingleCondition  Expr
	BeginCondition   Expr
	EmitCondition    Expr
	WindowType       WindowType
	Delay            *IntegerLiteral
	Length           *IntegerLiteral
	Interval         *IntegerLiteral
	TimeUnit         *TimeLiteral
	Filter           Expr
	Expr
}

type SortField struct {
	Name       string
	StreamName StreamName
	Uname      string // unique name of a field
	Ascending  bool
	FieldExpr  Expr

	Expr
}

func (sf *SortField) String() string { _ = "STUB: not implemented"; return "" }

func (wd *Window) String() string { _ = "STUB: not implemented"; return "" }

type SortFields []SortField

func (d SortFields) node() { _ = "STUB: not implemented"; return }

const (
	RowkindInsert = "insert"
	RowkindUpdate = "update"
	RowkindUpsert = "upsert"
	RowkindDelete = "delete"
)
