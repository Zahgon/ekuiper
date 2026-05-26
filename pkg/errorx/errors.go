// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

type Error struct {
	msg  string
	code ErrorCode
}

func New(message string) *Error { _ = "STUB: not implemented"; return nil }

func NewWithCode(code ErrorCode, message string) *Error { _ = "STUB: not implemented"; return nil }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Code() ErrorCode { _ = "STUB: not implemented"; return *new(ErrorCode) }

type ErrorWithCode interface {
	Error() string
	Code() ErrorCode
}

func IsRecoverAbleError(err error) bool { _ = "STUB: not implemented"; return false }

// consider timeout and temporary error as recoverable

type MockTemporaryError struct{}

func (e *MockTemporaryError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *MockTemporaryError) Temporary() bool { _ = "STUB: not implemented"; return false }
