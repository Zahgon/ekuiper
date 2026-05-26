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
	"path/filepath"

	"github.com/mattn/go-tflite"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var ipManager *interpreterManager

func init() {
	path, err := conf.GetDataLoc()
	if err != nil {
		panic(err)
	}
	ipManager = &interpreterManager{
		registry: make(map[string]*tflite.Interpreter),
		path:     filepath.Join(path, "uploads"),
	}
}

type interpreterManager struct {
	syncx.Mutex
	registry map[string]*tflite.Interpreter
	path     string
}

func (m *interpreterManager) GetOrCreate(name string) (*tflite.Interpreter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
