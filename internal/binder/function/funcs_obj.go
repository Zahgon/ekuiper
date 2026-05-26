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
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

func registerObjectFunc() { _ = "STUB: not implemented"; return }

// The argument can be {obj}, {string arr} OR {obj}, {string}...

func pick(ctx api.FunctionContext, res map[string]any, argMap map[string]any, k string) {
	_ = "STUB: not implemented"
	return
}

func sliceStringContains(s []string, target string) bool { _ = "STUB: not implemented"; return false }
