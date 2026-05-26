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

package protobuf

import (

	// TODO: replace with `google.golang.org/protobuf/proto` pkg.
	//nolint:staticcheck

	"github.com/jhump/protoreflect/desc"    //nolint:staticcheck
	"github.com/jhump/protoreflect/dynamic" //nolint:staticcheck

	"github.com/lf-edge/ekuiper/v2/pkg/cast"
)

const (
	WrapperBool   = "google.protobuf.BoolValue"
	WrapperBytes  = "google.protobuf.BytesValue"
	WrapperDouble = "google.protobuf.DoubleValue"
	WrapperFloat  = "google.protobuf.FloatValue"
	WrapperInt32  = "google.protobuf.Int32Value"
	WrapperInt64  = "google.protobuf.Int64Value"
	WrapperString = "google.protobuf.StringValue"
	WrapperUInt32 = "google.protobuf.UInt32Value"
	WrapperUInt64 = "google.protobuf.UInt64Value"
	WrapperVoid   = "google.protobuf.EMPTY"
)

var WRAPPER_TYPES = map[string]struct{}{
	WrapperBool:   {},
	WrapperBytes:  {},
	WrapperDouble: {},
	WrapperFloat:  {},
	WrapperInt32:  {},
	WrapperInt64:  {},
	WrapperString: {},
	WrapperUInt32: {},
	WrapperUInt64: {},
}

var (
	fieldConverterIns = &FieldConverter{}
	mf                = dynamic.NewMessageFactoryWithDefaults()
)

type FieldConverter struct{}

func GetFieldConverter() *FieldConverter { _ = "STUB: not implemented"; return nil }

func (fc *FieldConverter) encodeMap(im *desc.MessageDescriptor, i interface{}) (*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FieldConverter) EncodeField(field *desc.FieldDescriptor, v interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FieldConverter) encodeSingleField(field *desc.FieldDescriptor, v interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FieldConverter) DecodeField(src interface{}, field *desc.FieldDescriptor, sn cast.Strictness) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FieldConverter) decodeSubMessage(input interface{}, ft *desc.MessageDescriptor, sn cast.Strictness) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FieldConverter) DecodeMap(src map[string]interface{}, ft *desc.MessageDescriptor, sn cast.Strictness) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fc *FieldConverter) decodeMessageField(src interface{}, field *desc.FieldDescriptor, result map[string]interface{}, sn cast.Strictness) error {
	_ = "STUB: not implemented"
	return nil
}

func (fc *FieldConverter) DecodeMessage(message *dynamic.Message, outputType *desc.MessageDescriptor) interface{} {
	_ = "STUB: not implemented"
	return nil
}
