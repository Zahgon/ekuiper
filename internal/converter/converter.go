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

package converter

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/converter/binary"
	"github.com/lf-edge/ekuiper/v2/internal/converter/delimited"
	"github.com/lf-edge/ekuiper/v2/internal/converter/json"
	"github.com/lf-edge/ekuiper/v2/internal/converter/urlencoded"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

func init() {
	modules.RegisterConverter(message.FormatJson, func(_ api.StreamContext, _ string, schema map[string]*ast.JsonStreamField, props map[string]any) (message.Converter, error) {
		return json.NewFastJsonConverter(schema, props), nil
	})
	modules.RegisterConverter(message.FormatBinary, func(_ api.StreamContext, _ string, _ map[string]*ast.JsonStreamField, props map[string]any) (message.Converter, error) {
		return binary.GetConverter()
	})
	modules.RegisterConverter(message.FormatDelimited, func(_ api.StreamContext, _ string, _ map[string]*ast.JsonStreamField, props map[string]any) (message.Converter, error) {
		return delimited.NewConverter(props)
	})
	modules.RegisterConverter(message.FormatUrlEncoded, func(_ api.StreamContext, _ string, _ map[string]*ast.JsonStreamField, props map[string]any) (message.Converter, error) {
		return urlencoded.NewConverter(props)
	})
	modules.RegisterWriterConverter(message.FormatDelimited, func(ctx api.StreamContext, _ string, _ map[string]*ast.JsonStreamField, props map[string]any) (message.ConvertWriter, error) {
		return delimited.NewCsvWriter(ctx, props)
	})
	modules.RegisterWriterConverter(message.FormatJson, func(ctx api.StreamContext, _ string, schema map[string]*ast.JsonStreamField, props map[string]any) (message.ConvertWriter, error) {
		return json.NewFastJsonConverter(schema, props), nil
	})
}

func GetOrCreateConverter(ctx api.StreamContext, format string, schemaId string, schemaFields map[string]*ast.JsonStreamField, props map[string]any) (c message.Converter, err error) {
	_ = "STUB: not implemented"
	return *new(message.Converter), nil
}

func GetConvertWriter(ctx api.StreamContext, format string, schemaId string, schema map[string]*ast.JsonStreamField, props map[string]any) (message.ConvertWriter, error) {
	_ = "STUB: not implemented"
	return *new(message.ConvertWriter), nil
}

func GetMerger(ctx api.StreamContext, format string, schemaId string, schemaFields map[string]*ast.JsonStreamField, props map[string]any) (modules.Merger, error) {
	_ = "STUB: not implemented"
	return *new(modules.Merger), nil
}

func transSchemaId(t, schemaId string, props map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If not require schema, just return the schemaId. And the register function need to deal with it by itself. Only the specific implementation defines the schemaId format.
