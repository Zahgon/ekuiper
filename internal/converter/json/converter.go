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

package json

import (
	"bytes"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/valyala/fastjson"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type FastJsonConverter struct {
	syncx.RWMutex
	schema map[string]*ast.JsonStreamField
	FastJsonConverterConf
	isSlice bool
	buffer  bytes.Buffer
	isNew   bool
}

func (f *FastJsonConverter) New(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (f *FastJsonConverter) Write(ctx api.StreamContext, d any) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FastJsonConverter) Flush(_ api.StreamContext) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FastJsonConverterConf struct {
	UseInt64        bool              `json:"useInt64ForWholeNumber"`
	ColAliasMapping map[string]string `json:"colAliasMapping"`
}

func NewFastJsonConverter(schema map[string]*ast.JsonStreamField, props map[string]any) *FastJsonConverter {
	_ = "STUB: not implemented"
	return nil
}

func (f *FastJsonConverter) setupProps(props map[string]any) { _ = "STUB: not implemented"; return }

func (f *FastJsonConverter) ResetSchema(schema map[string]*ast.JsonStreamField) {
	_ = "STUB: not implemented"
	return
}

func (f *FastJsonConverter) Encode(ctx api.StreamContext, d any) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FastJsonConverter) Decode(ctx api.StreamContext, b []byte) (m any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *FastJsonConverter) DecodeField(_ api.StreamContext, b []byte, field string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *FastJsonConverter) decodeWithSchema(b []byte, schema map[string]*ast.JsonStreamField) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *FastJsonConverter) decodeArray(array []*fastjson.Value, field *ast.JsonStreamField) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FastJsonConverter) decodeObject(obj *fastjson.Object, schema map[string]*ast.JsonStreamField, isOuter bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FastJsonConverter) checkSchema(key, typ string, schema map[string]*ast.JsonStreamField) (add, valid bool) {
	_ = "STUB: not implemented"
	// for schemaless, allow to decode the key value
	return false, false
}

// for defined schema, skip to decode undefined key

// for the schema we didn't parse,allow to decode eg: results[0].a.b

// for the defined schema type, directly to check

func (f *FastJsonConverter) extractNumberValue(name string, v *fastjson.Value, field *ast.JsonStreamField) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FastJsonConverter) extractStringValue(name string, v *fastjson.Value, field *ast.JsonStreamField) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FastJsonConverter) extractBooleanFromValue(name string, v *fastjson.Value, field *ast.JsonStreamField) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FastJsonConverter) extractNumber(v *fastjson.Value) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *FastJsonConverter) decodeToSlice(v *fastjson.Value, schema map[string]*ast.JsonStreamField) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *FastJsonConverter) decodeObject2Slice(obj *fastjson.Object, schema map[string]*ast.JsonStreamField, isOuter bool) (model.SliceVal, error) {
	_ = "STUB: not implemented"
	return *new(model.SliceVal), nil
}

func getBooleanFromValue(value *fastjson.Value) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getType(t *ast.JsonStreamField) string { _ = "STUB: not implemented"; return "" }

func isFloat64(v string) bool { _ = "STUB: not implemented"; return false }

var _ message.ConvertWriter = &FastJsonConverter{}
