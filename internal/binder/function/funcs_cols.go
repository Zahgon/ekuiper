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
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type ResultCols map[string]interface{}

// ColFunc Functions which will return columns directly instead of a map
type ColFunc func(ctx api.FunctionContext, args []interface{}, keys []string) (ResultCols, error)

func wrapColFunc(colFunc ColFunc) funcExe { _ = "STUB: not implemented"; return *new(funcExe) }

func registerColsFunc() { _ = "STUB: not implemented"; return }

func changedFunc(ctx api.FunctionContext, args []interface{}, keys []string) (ResultCols, error) {
	_ = "STUB: not implemented"
	// validation
	return *new(ResultCols), nil
}
