// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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
	"fmt"
)

var (
	errorArrayFirstArgumentNotArrayError   = fmt.Errorf("first argument should be array of interface{}")
	errorArrayIndex                        = fmt.Errorf("index out of range")
	errorArraySecondArgumentNotArrayError  = fmt.Errorf("second argument should be array of interface{}")
	errorArrayFirstArgumentNotIntError     = fmt.Errorf("first argument should be int")
	errorArrayFirstArgumentNotStringError  = fmt.Errorf("first argument should be string")
	errorArraySecondArgumentNotIntError    = fmt.Errorf("second argument should be int")
	errorArraySecondArgumentNotStringError = fmt.Errorf("second argument should be string")
	errorArrayThirdArgumentNotIntError     = fmt.Errorf("third argument should be int")
	errorArrayThirdArgumentNotStringError  = fmt.Errorf("third argument should be string")
	errorArrayNotArrayElementError         = fmt.Errorf("array elements should be array")
	errorArrayNotStringElementError        = fmt.Errorf("array elements should be string")
)

func registerArrayFunc() { _ = "STUB: not implemented"; return }

// all un-hashable types are not deduplicated, including array, map, etc.

// Sort the slice if it contains elements that are comparable
