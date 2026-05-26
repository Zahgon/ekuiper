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

package transform

import (
	"text/template"
)

func GenTp(dt string) (*template.Template, error) { _ = "STUB: not implemented"; return nil, nil }

// TransItem If you do not need to convert data to []byte, you can use this function directly. Otherwise, use TransFunc.
func TransItem(input interface{}, dataField string, fields []string, excludeFields []string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

// selectMap select fields from input map or array of map.
func selectMap(input any, fields []string, excludeFields []string) (any, error) {
	_ = "STUB: not implemented"
	// can only have fields or excludeFields
	return *new(any), nil
}
