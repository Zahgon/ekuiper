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

//go:build test

package converter

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/schema"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

func init() {
	modules.RegisterConverter("mock", func(_ api.StreamContext, _ string, _ map[string]*ast.JsonStreamField, props map[string]any) (message.Converter, error) {
		return &MockConverter{}, nil
	})
	modules.RegisterSchemaType(modules.CUSTOM, &schema.CustomType{}, ".so")
	modules.RegisterConverterSchemas("mock", "protobuf")
}

// MockConverter mocks a slow converter for benchmark test
type MockConverter struct{}

func (m MockConverter) Encode(ctx api.StreamContext, d any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m MockConverter) Decode(ctx api.StreamContext, b []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
