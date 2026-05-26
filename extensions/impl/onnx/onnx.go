// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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

package onnx

import (
	_ "bytes"
	_ "image"
	_ "image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type OnnxFunc struct{}

// Validate the arguments.
// args[0]: string, model name which maps to a path
// args[1 to n]: tensors
func (f *OnnxFunc) Validate(args []interface{}) error { _ = "STUB: not implemented"; return nil }

func (f *OnnxFunc) Exec(ctx api.FunctionContext, args []any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Set input tensors

// only supports one dimensional arg. Even dim 0 must be an array of 1 element

// convert onnx's type float to float32 of golang

// not support

// The float16.Float16 type is just a uint16 underneath; write its
// bytes to the data slice.

// not support ,but dont need transfer becase string can look as []byte

// not support ，transfer to []byte

// support list see ：GetTensorElementDataType() and TensorElementDataType in onnxruntime_go

// todo :optimize: avoid creating output tensor every time

// for output , only transfer go build-in type

func (f *OnnxFunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }
