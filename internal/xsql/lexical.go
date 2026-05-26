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

package xsql

import (
	"bufio"
	"bytes"
	"io"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type Scanner struct {
	r   *bufio.Reader
	buf *bytes.Buffer
}

func NewScanner(r io.Reader) *Scanner { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Scan() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

// s.unread()

func (s *Scanner) ScanIdent() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (s *Scanner) ScanString(isSingle bool) (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (s *Scanner) ScanDigit() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (s *Scanner) ScanNumber(startWithDot bool, isNeg bool) (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (s *Scanner) ScanBackquoteIdent() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (s *Scanner) skipUntilNewline() { _ = "STUB: not implemented"; return }

func (s *Scanner) skipUntilEndComment() error { _ = "STUB: not implemented"; return nil }

// We might be at the end.

func (s *Scanner) ScanWhiteSpace() (tok ast.Token, lit string) {
	_ = "STUB: not implemented"
	return *new(ast.Token), ""
}

func (s *Scanner) read() rune { _ = "STUB: not implemented"; return 0 }

func (s *Scanner) unread() { _ = "STUB: not implemented"; return }

var eof = rune(0)

func isWhiteSpace(r rune) bool { _ = "STUB: not implemented"; return false }

func isLetter(ch rune) bool { _ = "STUB: not implemented"; return false }

func isDigit(ch rune) bool { _ = "STUB: not implemented"; return false }

func isQuotation(ch rune) bool { _ = "STUB: not implemented"; return false }

func isBackquote(ch rune) bool { _ = "STUB: not implemented"; return false }
