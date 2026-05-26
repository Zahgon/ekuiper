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

package portable

import (
	"sync"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/plugin"
)

func (m *Manager) Source(name string) (api.Source, error) {
	_ = "STUB: not implemented"
	return *new(api.Source), nil
}

func (m *Manager) SourcePluginInfo(name string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func (m *Manager) LookupSource(_ string) (api.Source, error) {
	_ = "STUB: not implemented"
	// TODO add support
	return *new(api.Source), nil
}

func (m *Manager) Sink(name string) (api.Sink, error) {
	_ = "STUB: not implemented"
	return *new(api.Sink), nil
}

func (m *Manager) SinkPluginInfo(name string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

// The function instance are kept forever even after deletion
// The instance is actually a wrapper of the nng channel which is dependant from the plugin instance
// Even updated plugin instance can reuse the channel if the function name is not changed
// It is not used to check if the function is bound, use ConvName which checks the meta
var funcInsMap = &sync.Map{}

func (m *Manager) Function(name string) (api.Function, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), nil
}

func (m *Manager) HasFunctionSet(funcName string) bool { _ = "STUB: not implemented"; return false }

func (m *Manager) FunctionPluginInfo(funcName string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func (m *Manager) ConvName(funcName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
