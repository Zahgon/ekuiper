// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

package errorx

type ErrorCode int

const (
	Undefined_Err ErrorCode = 1000
	GENERAL_ERR   ErrorCode = 1001
	NOT_FOUND     ErrorCode = 1002
	IOErr         ErrorCode = 1003
	CovnerterErr  ErrorCode = 1004
	EOF           ErrorCode = 1005

	// error code for sql

	ParserError   ErrorCode = 2001
	PlanError     ErrorCode = 2101
	ExecutorError ErrorCode = 2201

	StreamTableError ErrorCode = 3000
	RuleErr          ErrorCode = 4000
	ConfKeyError     ErrorCode = 5000
)

var NotFoundErr = NewWithCode(NOT_FOUND, "not found")

func NewIOErr(msg string) error { _ = "STUB: not implemented"; return nil }

func NewEOF(msg string) error { _ = "STUB: not implemented"; return nil }

func IsIOError(err error) bool { _ = "STUB: not implemented"; return false }

func IsEOF(err error) bool { _ = "STUB: not implemented"; return false }

func IsUnexpectedErr(err error) bool { _ = "STUB: not implemented"; return false }

func NewParserError(msg string) error { _ = "STUB: not implemented"; return nil }

func GetErrorCode(err error) (ErrorCode, bool) {
	_ = "STUB: not implemented"
	return *new(ErrorCode), false
}
