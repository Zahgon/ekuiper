// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

package service

import (
	"sync"

	"github.com/jhump/protoreflect/desc"            //nolint:staticcheck
	"github.com/jhump/protoreflect/desc/protoparse" //nolint:staticcheck
	"github.com/jhump/protoreflect/dynamic"         //nolint:staticcheck

	// introduce annotations
	_ "google.golang.org/genproto/googleapis/api/annotations"

	"github.com/lf-edge/ekuiper/v2/internal/converter/protobuf"
)

type descriptor interface {
	GetFunctions() []string
}

type protoDescriptor interface {
	ConvertParamsToMessage(method string, params []interface{}) (*dynamic.Message, error)
	ConvertReturnMessage(method string, returnVal *dynamic.Message) (interface{}, error)
	MethodDescriptor(method string) *desc.MethodDescriptor
	MessageFactory() *dynamic.MessageFactory
}

type jsonDescriptor interface {
	ConvertParamsToJson(method string, params []interface{}) ([]byte, error)
	ConvertReturnJson(method string, returnVal []byte) (interface{}, error)
}

type textDescriptor interface {
	ConvertParamsToText(method string, params []interface{}) ([]byte, error)
	ConvertReturnText(method string, returnVal []byte) (interface{}, error)
}

type interfaceDescriptor interface {
	ConvertParams(method string, params []interface{}) ([]interface{}, error)
	ConvertReturn(method string, returnVal interface{}) (interface{}, error)
}

type multiplexDescriptor interface {
	jsonDescriptor
	textDescriptor
	interfaceDescriptor
	httpMapping
}

var ( // Do not call these directly, use the get methods
	protoParser *protoparse.Parser
	// A buffer of descriptor for schemas
	reg = &sync.Map{}
)

func ProtoParser() *protoparse.Parser { _ = "STUB: not implemented"; return nil }

func parse(schema schema, file string, schemaless bool) (descriptor, error) {
	_ = "STUB: not implemented"
	return *new(descriptor), nil
}

type wrappedSchemalessDescriptor struct{}

func (d *wrappedSchemalessDescriptor) GetFunctions() (result []string) {
	_ = "STUB: not implemented"
	return nil
}

func (d *wrappedSchemalessDescriptor) ConvertParamsToMessage(_ string, _ []interface{}) (*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) ConvertReturnMessage(_ string, _ *dynamic.Message) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) MethodDescriptor(_ string) *desc.MethodDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (d *wrappedSchemalessDescriptor) MessageFactory() *dynamic.MessageFactory {
	_ = "STUB: not implemented"
	return nil
}

func (d *wrappedSchemalessDescriptor) ConvertParamsToJson(_ string, params []interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) ConvertReturnJson(_ string, returnVal []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) ConvertParamsToText(_ string, params []interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) ConvertReturnText(_ string, returnVal []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) ConvertParams(_ string, params []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedSchemalessDescriptor) ConvertReturn(_ string, params interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type wrappedProtoDescriptor struct {
	*desc.FileDescriptor
	methodOptions map[string]*httpOptions
	mf            *dynamic.MessageFactory
	fc            *protobuf.FieldConverter
}

// GetFunctions TODO support for duplicate names
func (d *wrappedProtoDescriptor) GetFunctions() (result []string) {
	_ = "STUB: not implemented"
	return nil
}

func (d *wrappedProtoDescriptor) MessageFactory() *dynamic.MessageFactory {
	_ = "STUB: not implemented"

	// ConvertParams TODO support optional field, support enum type
	// Parameter mapping for protobuf
	// 1. If param length is 1, it can either a map contains all field or a field only.
	// 2. If param length is more then 1, they will map to message fields in the order
	return nil
}

func (d *wrappedProtoDescriptor) ConvertParams(method string, params []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) ConvertParamsToMessage(method string, params []interface{}) (*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) ConvertParamsToJson(method string, params []interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	// Deal with encoded json string. Just return the string
	return nil, nil
}

func (d *wrappedProtoDescriptor) ConvertParamsToText(method string, params []interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) convertParams(im *desc.MessageDescriptor, params []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If it is map, try unfold it
// TODO custom error for non map or map name not match

// For non map params, treat it as special case of multiple params

func (d *wrappedProtoDescriptor) ConvertReturn(method string, returnVal interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MUST be a map

func (d *wrappedProtoDescriptor) ConvertReturnMessage(method string, returnVal *dynamic.Message) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) ConvertReturnJson(method string, returnVal []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) ConvertReturnText(method string, returnVal []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) MethodDescriptor(name string) *desc.MethodDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (d *wrappedProtoDescriptor) unfoldMap(ft *desc.MessageDescriptor, i interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
