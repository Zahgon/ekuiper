// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package function

import (
	"errors"
)

var errTooManyArguments = errors.New("too many arguments")

type IntervalUnit string

// registerDateTimeFunc registers the date and time functions.
func registerDateTimeFunc() { _ = "STUB: not implemented"; return }

func execGetCurrentDate() funcExe { _ = "STUB: not implemented"; return *new(funcExe) }

// validFspArgs returns a function that validates the 'fsp' arg.
func validFspArgs() funcVal { _ = "STUB: not implemented"; return *new(funcVal) }

func execGetCurrentDateTime(timeOnly bool) funcExe { _ = "STUB: not implemented"; return *new(funcExe) }

// getCurrentWithFsp returns the current date/time with the specified number of fractional seconds precision.
func getCurrentWithFsp(fsp int, timeOnly bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
