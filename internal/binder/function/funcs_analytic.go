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

package function

// registerAnalyticFunc registers the analytic functions
// The last parameter of the function is always the partition key
func registerAnalyticFunc() { _ = "STUB: not implemented"; return }

// first time call, need create state for lag

// notice nil is ignored in latest

func getMax(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }

func getMin(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }
