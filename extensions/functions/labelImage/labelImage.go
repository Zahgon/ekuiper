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

//go:build tflite

package main

import (
	_ "image/jpeg"
	_ "image/png"
	"sync"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/mattn/go-tflite" //nolint:typecheck
)

type labelImage struct {
	modelPath   string
	labelPath   string
	once        sync.Once
	interpreter *tflite.Interpreter
	labels      []string
}

func (f *labelImage) Validate(args []interface{}) error { _ = "STUB: not implemented"; return nil }

func (f *labelImage) Exec(args []interface{}, ctx api.FunctionContext) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TODO If created, the interpreter will be kept through the whole life of kuiper. Refactor this later.
// defer interpreter.Delete()

// output is the biggest score labelImage

func (f *labelImage) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func loadLabels(filename string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

var LabelImage = labelImage{
	modelPath: "labelImage/mobilenet_quant_v1_224.tflite",
	labelPath: "labelImage/labels.txt",
}
