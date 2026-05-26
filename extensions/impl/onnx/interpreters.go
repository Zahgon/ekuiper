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

package onnx

import (
	_ "image"
	_ "image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"sync"

	ort "github.com/yalue/onnxruntime_go"

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
		registry: make(map[string]*InterPreter),
		path:     filepath.Join(path, "uploads"),
	}
	if conf.IsTesting {
		ipManager.path = "test"
	}
}

type interpreterManager struct {
	once       sync.Once
	envInitErr error
	syncx.Mutex
	registry map[string]*InterPreter
	path     string
}

func (m *interpreterManager) GetOrCreate(name string) (*InterPreter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDefaultSharedLibPath() string {
	_ = "STUB: not implemented"
	// For now, we only include libraries for ARM64 darwin and x86_64 or ARM64 Linux. In the future, libraries may be added or removed.
	return ""
}

type InterPreter struct {
	session    *ort.DynamicAdvancedSession
	inputInfo  []ort.InputOutputInfo
	outputInfo []ort.InputOutputInfo
}

func NewInterPreter(session *ort.DynamicAdvancedSession,
	inputInfo []ort.InputOutputInfo,
	outputInfo []ort.InputOutputInfo,
) *InterPreter {
	_ = "STUB: not implemented"
	return nil
}

func (ip *InterPreter) GetInputTensorCount() int { _ = "STUB: not implemented"; return 0 }

func (ip *InterPreter) GetEmptyOutputTensors() ([]ort.ArbitraryTensor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newEmptyArbitraryTensorBydataType(dataType ort.TensorElementDataType, shape ort.Shape) (ort.ArbitraryTensor, error) {
	_ = "STUB: not implemented"
	return *new(ort.ArbitraryTensor), nil
}

// todo more dataType
