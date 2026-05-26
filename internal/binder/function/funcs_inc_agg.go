// Copyright 2024 EMQ Technologies Co., Ltd.
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

var supportedIncAggFunc = map[string]struct{}{
	"count":      {},
	"avg":        {},
	"max":        {},
	"min":        {},
	"sum":        {},
	"merge_agg":  {},
	"collect":    {},
	"last_value": {},
}

func IsSupportedIncAgg(name string) bool { _ = "STUB: not implemented"; return false }

func registerIncAggFunc() { _ = "STUB: not implemented"; return }

func incrementalLastValue(ctx api.FunctionContext, arg interface{}, ignoreNil bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementalCollect(ctx api.FunctionContext, arg interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementalMerge(ctx api.FunctionContext, arg map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementalMin(ctx api.FunctionContext, arg interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementalMax(ctx api.FunctionContext, arg interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrementalCount(ctx api.FunctionContext, arg interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func incrementalSum(ctx api.FunctionContext, arg float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
