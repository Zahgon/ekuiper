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

package ast

type Node interface {
	node()
}

type NameNode interface {
	Node
	GetName() string
}

type ValidateAbleExpr interface {
	ValidateExpr() error
}

type Expr interface {
	Node
	expr()
	// String function for the explain grammar, convert Expr to String
	String() string
}

type Literal interface {
	Expr
	literal()
}

type ParenExpr struct {
	Expr Expr
}

type ArrowExpr struct {
	Expr Expr
}

type BracketExpr struct {
	Expr Expr
}

type ColonExpr struct {
	Start Expr
	End   Expr
}

func (c *ColonExpr) ValidateExpr() error { _ = "STUB: not implemented"; return nil }

type IndexExpr struct {
	Index Expr
}

type BooleanLiteral struct {
	Val bool
}

type TimeLiteral struct {
	Val Token
}

type IntegerLiteral struct {
	Val int64
}

type StringLiteral struct {
	Val string
}

type NumberLiteral struct {
	Val float64
}

type Wildcard struct {
	Token   Token
	Replace []Field
	Except  []string
}

func (pe *ParenExpr) expr()          { _ = "STUB: not implemented"; return }
func (pe *ParenExpr) node()          { _ = "STUB: not implemented"; return }
func (pe *ParenExpr) String() string { _ = "STUB: not implemented"; return "" }

func (ae *ArrowExpr) expr()          { _ = "STUB: not implemented"; return }
func (ae *ArrowExpr) node()          { _ = "STUB: not implemented"; return }
func (ae *ArrowExpr) String() string { _ = "STUB: not implemented"; return "" }

func (be *BracketExpr) expr()          { _ = "STUB: not implemented"; return }
func (be *BracketExpr) node()          { _ = "STUB: not implemented"; return }
func (be *BracketExpr) String() string { _ = "STUB: not implemented"; return "" }

func (be *ColonExpr) expr()          { _ = "STUB: not implemented"; return }
func (be *ColonExpr) node()          { _ = "STUB: not implemented"; return }
func (be *ColonExpr) String() string { _ = "STUB: not implemented"; return "" }

func (be *IndexExpr) expr()          { _ = "STUB: not implemented"; return }
func (be *IndexExpr) node()          { _ = "STUB: not implemented"; return }
func (be *IndexExpr) String() string { _ = "STUB: not implemented"; return "" }

func (w *Wildcard) expr()          { _ = "STUB: not implemented"; return }
func (w *Wildcard) node()          { _ = "STUB: not implemented"; return }
func (w *Wildcard) String() string { _ = "STUB: not implemented"; return "" }

func (bl *BooleanLiteral) expr()          { _ = "STUB: not implemented"; return }
func (bl *BooleanLiteral) literal()       { _ = "STUB: not implemented"; return }
func (bl *BooleanLiteral) node()          { _ = "STUB: not implemented"; return }
func (bl *BooleanLiteral) String() string { _ = "STUB: not implemented"; return "" }

func (tl *TimeLiteral) expr()          { _ = "STUB: not implemented"; return }
func (tl *TimeLiteral) literal()       { _ = "STUB: not implemented"; return }
func (tl *TimeLiteral) node()          { _ = "STUB: not implemented"; return }
func (tl *TimeLiteral) String() string { _ = "STUB: not implemented"; return "" }

func (il *IntegerLiteral) expr()          { _ = "STUB: not implemented"; return }
func (il *IntegerLiteral) literal()       { _ = "STUB: not implemented"; return }
func (il *IntegerLiteral) node()          { _ = "STUB: not implemented"; return }
func (il *IntegerLiteral) String() string { _ = "STUB: not implemented"; return "" }

func (nl *NumberLiteral) expr()          { _ = "STUB: not implemented"; return }
func (nl *NumberLiteral) literal()       { _ = "STUB: not implemented"; return }
func (nl *NumberLiteral) node()          { _ = "STUB: not implemented"; return }
func (nl *NumberLiteral) String() string { _ = "STUB: not implemented"; return "" }

func (sl *StringLiteral) expr()          { _ = "STUB: not implemented"; return }
func (sl *StringLiteral) literal()       { _ = "STUB: not implemented"; return }
func (sl *StringLiteral) node()          { _ = "STUB: not implemented"; return }
func (sl *StringLiteral) String() string { _ = "STUB: not implemented"; return "" }

type FuncType int

const (
	FuncTypeUnknown FuncType = iota - 1
	FuncTypeScalar
	FuncTypeAgg
	FuncTypeCols
	FuncTypeSrf
	FuncTypeWindow
	FuncTypeTrigger
)

type Call struct {
	Name     string
	FuncId   int
	FuncType FuncType
	Args     []Expr
	// This is used for analytic functions.
	// In planner, all analytic functions are planned to calculate in analytic_op which produce a new field.
	// This cachedField cached the new field name and when evaluating, just returned the field access evaluated value.
	CachedField string
	CacheIndex  int
	Cached      bool
	Partition   *PartitionExpr
	WhenExpr    Expr

	// This is used for window functions.
	SortFields SortFields
}

func (c *Call) expr()          { _ = "STUB: not implemented"; return }
func (c *Call) literal()       { _ = "STUB: not implemented"; return }
func (c *Call) node()          { _ = "STUB: not implemented"; return }
func (c *Call) String() string { _ = "STUB: not implemented"; return "" }

type PartitionExpr struct {
	Exprs []Expr
}

func (pe *PartitionExpr) expr()          { _ = "STUB: not implemented"; return }
func (pe *PartitionExpr) node()          { _ = "STUB: not implemented"; return }
func (pe *PartitionExpr) String() string { _ = "STUB: not implemented"; return "" }

type BinaryExpr struct {
	OP  Token
	LHS Expr
	RHS Expr
}

func (be *BinaryExpr) ValidateExpr() error { _ = "STUB: not implemented"; return nil }

func (be *BinaryExpr) expr()          { _ = "STUB: not implemented"; return }
func (be *BinaryExpr) node()          { _ = "STUB: not implemented"; return }
func (be *BinaryExpr) String() string { _ = "STUB: not implemented"; return "" }

type WhenClause struct {
	// The condition Expression
	Expr   Expr
	Result Expr
}

func (w *WhenClause) expr()          { _ = "STUB: not implemented"; return }
func (w *WhenClause) node()          { _ = "STUB: not implemented"; return }
func (w *WhenClause) String() string { _ = "STUB: not implemented"; return "" }

type CaseExpr struct {
	// The compare value Expression. It can be a value Expression or nil.
	// When it is nil, the WhenClause Expr must be a logical(comparison) Expression
	Value       Expr
	WhenClauses []*WhenClause
	ElseClause  Expr
}

func (c *CaseExpr) expr()          { _ = "STUB: not implemented"; return }
func (c *CaseExpr) node()          { _ = "STUB: not implemented"; return }
func (c *CaseExpr) String() string { _ = "STUB: not implemented"; return "" }

type ValueSetExpr struct {
	LiteralExprs []Expr // ("A", "B", "C") or (1, 2, 3)
	ArrayExpr    Expr
}

func (c *ValueSetExpr) expr()          { _ = "STUB: not implemented"; return }
func (c *ValueSetExpr) node()          { _ = "STUB: not implemented"; return }
func (c *ValueSetExpr) String() string { _ = "STUB: not implemented"; return "" }

type BetweenExpr struct {
	Lower  Expr
	Higher Expr
}

func (b *BetweenExpr) expr()          { _ = "STUB: not implemented"; return }
func (b *BetweenExpr) node()          { _ = "STUB: not implemented"; return }
func (b *BetweenExpr) String() string { _ = "STUB: not implemented"; return "" }

type LimitExpr struct {
	LimitCount *IntegerLiteral
}

func (l *LimitExpr) expr()          { _ = "STUB: not implemented"; return }
func (l *LimitExpr) node()          { _ = "STUB: not implemented"; return }
func (l *LimitExpr) String() string { _ = "STUB: not implemented"; return "" }

type StreamName string

func (sn *StreamName) node() { _ = "STUB: not implemented"; return }

const (
	DefaultStream = StreamName("$$default")
	AliasStream   = StreamName("$$alias")
)

type MetaRef struct {
	StreamName StreamName
	Name       string
}

func (fr *MetaRef) expr()          { _ = "STUB: not implemented"; return }
func (fr *MetaRef) node()          { _ = "STUB: not implemented"; return }
func (fr *MetaRef) String() string { _ = "STUB: not implemented"; return "" }

type JsonFieldRef struct {
	Name string
}

func (fr *JsonFieldRef) expr()          { _ = "STUB: not implemented"; return }
func (fr *JsonFieldRef) node()          { _ = "STUB: not implemented"; return }
func (fr *JsonFieldRef) String() string { _ = "STUB: not implemented"; return "" }

type ColFuncField struct {
	Name string
	Expr Expr
}

func (fr *ColFuncField) expr()          { _ = "STUB: not implemented"; return }
func (fr *ColFuncField) node()          { _ = "STUB: not implemented"; return }
func (fr *ColFuncField) String() string { _ = "STUB: not implemented"; return "" }
