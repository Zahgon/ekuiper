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

package main

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

/**
 **	A function which will count how many words had been received from the beginning
 ** to demonstrate how to use states
 ** There are 2 arguments:
 **  0: column, the column to be calculated. The column value type must be string
 **  1: separator, a string literal for word separator
 **/

type accumulateWordCountFunc struct{}

func (f *accumulateWordCountFunc) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (f *accumulateWordCountFunc) Exec(ctx api.FunctionContext, args []any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (f *accumulateWordCountFunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func AccumulateWordCount() api.Function { _ = "STUB: not implemented"; return *new(api.Function) }
