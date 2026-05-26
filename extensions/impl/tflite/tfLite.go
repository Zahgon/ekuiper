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

//go:build tflite

package tflite

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type Tffunc struct{}

// Validate the arguments.
// args[0]: string, model name which maps to a path
// args[1 to n]: tensors
func (f *Tffunc) Validate(args []interface{}) error { _ = "STUB: not implemented"; return nil }

func (f *Tffunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (f *Tffunc) Exec(ctx api.FunctionContext, args []any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Set input tensors

// only supports one dimensional arg. Even dim 0 must be an array of 1 element

//outputSize := output.Dim(output.NumDims() - 1)
//b := make([]byte, outputSize)
//status = output.CopyToBuffer(&b[0])
//if status != tflite.OK {
//	return fmt.Errorf("output failed"), false
//}
//results[i] = b
