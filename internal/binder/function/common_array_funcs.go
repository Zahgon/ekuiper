// Copyright 2023 EMQ Technologies Co., Ltd.
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

// The functions here are used to implement the array functions to be referred in
// 1. Aggregate function
// 2. Array function

func max(arr []interface{}) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func min(arr []interface{}) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func getCount(s []interface{}) int { _ = "STUB: not implemented"; return 0 }

func getFirstValidArg(s []interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func sliceIntTotal(s []interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func sliceFloatTotal(s []interface{}) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func sliceIntMax(s []interface{}, max int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sliceFloatMax(s []interface{}, max float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sliceStringMax(s []interface{}, max string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func sliceIntMin(s []interface{}, min int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sliceFloatMin(s []interface{}, min float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sliceStringMin(s []interface{}, min string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func dedup(r []interface{}, col []interface{}, all bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
