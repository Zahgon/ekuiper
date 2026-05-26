// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

package operator

import (
	"time"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

// Only run when strict validation mode is on, fields is defined and is not binary
// Do not convert types
type defaultFieldProcessor struct {
	streamFields    map[string]*ast.JsonStreamField
	timestampFormat string
}

func (p *defaultFieldProcessor) validateAndConvert(tuple *xsql.Tuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *defaultFieldProcessor) validateAndConvertMessage(schema map[string]*ast.JsonStreamField, message xsql.Message) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate and convert field value to the type defined in schema
func (p *defaultFieldProcessor) validateAndConvertField(sf *ast.JsonStreamField, t interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *defaultFieldProcessor) parseTime(s string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
