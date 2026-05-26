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
	"io"

	_ "github.com/lf-edge/ekuiper/v2/internal/converter"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type Parser struct {
	s *Scanner

	i   int // buffer index
	n   int // buffer char count
	buf [3]struct {
		tok ast.Token
		lit string
	}
	inFunc      string // currently parsing function name
	f           int    // anonymous field index number
	fn          int    // function index number
	clause      string
	sourceNames []string // source names in the from/join clause
}

func (p *Parser) ParseCondition() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) ParseLimit() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) scan() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (p *Parser) curr() (ast.Token, string) { _ = "STUB: not implemented"; return *new(ast.Token), "" }

func (p *Parser) scanIgnoreWhitespace() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (p *Parser) unscan() { _ = "STUB: not implemented"; return }

func NewParser(r io.Reader) *Parser { _ = "STUB: not implemented"; return nil }

func NewParserWithSources(r io.Reader, sources []string) *Parser {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) ParseQueries() ([]ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) Parse() (*ast.SelectStatement, error) { _ = "STUB: not implemented"; return nil, nil }

// The source names may be injected from outside to parse part of the sql

func (p *Parser) parseSource() (ast.Sources, error) {
	_ = "STUB: not implemented"
	return *new(ast.Sources), nil
}

// TODO Current func has problems when the source includes white space.
func (p *Parser) parseSourceLiteral() (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// HASH, DIV & ADD token is specially support for MQTT topic name patterns.

func (p *Parser) parseFieldNameSections(isSubField bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parseJoins() (ast.Joins, error) {
	_ = "STUB: not implemented"
	return *new(ast.Joins), nil
}

func (p *Parser) ParseJoin(joinType ast.JoinType) (*ast.Join, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parseDimensions() (ast.Dimensions, error) {
	_ = "STUB: not implemented"
	return *new(ast.Dimensions), nil
}

func (p *Parser) parseHaving() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) parseSorts() (ast.SortFields, error) {
	_ = "STUB: not implemented"
	return *new(ast.SortFields), nil
}

func (p *Parser) parseFields() (ast.Fields, error) {
	_ = "STUB: not implemented"
	return *new(ast.Fields), nil
}

func (p *Parser) parseField() (*ast.Field, error) { _ = "STUB: not implemented"; return nil, nil }

func nameExpr(exp ast.Expr) string { _ = "STUB: not implemented"; return "" }

func (p *Parser) parseAlias() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *Parser) ParseExpr() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

// Change the asterisk to Mul token.

// LBRACKET is a special token, need to unscan

// IN is a special token, need to unscan

// IN is a special token, need to unscan

func (p *Parser) parseBetween(lhs ast.Expr, op ast.Token) (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) parseUnaryExpr(isSubField bool) (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

// Expect an RPAREN at the end.

// Back the Lparen token
// Back the ident token

func (p *Parser) parseValueSetExpr() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

// IN ("A", "B") or IN expression

// back to IN

func (p *Parser) parseBracketExpr() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

// field[]

// Such as field[2]

// Such as field[2:] or field[2:4]

// Such as field[:3] or [:]

// Such as field[2]

// Such as field[2:] or field[2:4]

func (p *Parser) parseColonExpr(start ast.Expr) (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) scanIgnoreWhiteSpaceWithNegativeNum() (ast.Token, string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (p *Parser) parseAs(f *ast.Field) (*ast.Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var WindowFuncs = map[string]struct{}{
	"tumblingwindow": {},
	"hoppingwindow":  {},
	"sessionwindow":  {},
	"slidingwindow":  {},
	"countwindow":    {},
	"statewindow":    {},
	"dedup_trigger":  {},
}

func convFuncName(n string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (p *Parser) parseCall(n string) (ast.Expr, error) {
	_ = "STUB: not implemented"
	// Check if n function exists and convert it to lowercase for built-in func
	return *new(ast.Expr), nil
}

// Add context for some aggregate func

// parse filter clause

// parse over when clause

func (p *Parser) parseCaseExpr() (*ast.CaseExpr, error) { _ = "STUB: not implemented"; return nil, nil }

// no condition value for case, additional validation needed

func validateWindows(fname string, args []ast.Expr) (ast.WindowType, error) {
	_ = "STUB: not implemented"
	return *new(ast.WindowType), nil
}

func validateWindow(funcName string, expectLen int, args []ast.Expr) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) ConvertToWindows(wtype ast.WindowType, args []ast.Expr) (*ast.Window, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) ParseCreateStmt() (ast.Statement, error) {
	_ = "STUB: not implemented"
	return *new(ast.Statement), nil
}

// Finish parsing create stream statement. Jump to validate

// TODO more accurate validation for table
func validateStream(stmt *ast.StreamStmt) error { _ = "STUB: not implemented"; return nil }

// do nothing for schemaless

func (p *Parser) parseShowStmt() (ast.Statement, error) {
	_ = "STUB: not implemented"
	return *new(ast.Statement), nil
}

func (p *Parser) parseDescribeStmt() (ast.Statement, error) {
	_ = "STUB: not implemented"
	return *new(ast.Statement), nil
}

func (p *Parser) parseExplainStmt() (ast.Statement, error) {
	_ = "STUB: not implemented"
	return *new(ast.Statement), nil
}

func (p *Parser) parseDropStmt() (ast.Statement, error) {
	_ = "STUB: not implemented"
	return *new(ast.Statement), nil
}

func (p *Parser) parseStreamFields() (ast.StreamFields, error) {
	_ = "STUB: not implemented"
	return *new(ast.StreamFields), nil
}

// For the schemaless streams
// create stream demo () WITH (FORMAT="JSON", DATASOURCE="demo" TYPE="edgex")

// Check the stack for LPAREN; If the stack for LPAREN is not zero, then it's not correct.

// The nested type definition of ARRAY and Struct, such as "field ARRAY(STRUCT(f BIGINT))"

func (p *Parser) parseStreamField() (*ast.StreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Just consume the comma.

func (p *Parser) parseStreamArrayType() (ast.FieldType, error) {
	_ = "STUB: not implemented"
	return *new(ast.FieldType), nil
}

func (p *Parser) parseStreamStructType() (ast.FieldType, error) {
	_ = "STUB: not implemented"
	return *new(ast.FieldType), nil
}

func (p *Parser) parseStreamOptions() (*ast.Options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// should not happen

func (p *Parser) ParseOver4Window(win *ast.Window) error { _ = "STUB: not implemented"; return nil }

// Only support filter on window now
func (p *Parser) parseFilter() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) parseAsterisk() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) inmeta() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) parseOver(c *ast.Call) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) parsePartitionBy() (*ast.PartitionExpr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parseWhen() (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func (p *Parser) parseInvisible() bool { _ = "STUB: not implemented"; return false }
