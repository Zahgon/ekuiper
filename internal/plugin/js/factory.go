// Copyright 2024 EMQ Technologies Co., Ltd.
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

package js

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/plugin"
)

func (m *Manager) Function(name string) (api.Function, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), nil
}

func (m *Manager) HasFunctionSet(_ string) bool { _ = "STUB: not implemented"; return false }

func (m *Manager) FunctionPluginInfo(funcName string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func (m *Manager) ConvName(n string) (string, bool) { _ = "STUB: not implemented"; return "", false }
