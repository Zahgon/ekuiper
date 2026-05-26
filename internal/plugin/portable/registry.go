// Copyright 2021 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type registry struct {
	syncx.RWMutex
	plugins map[string]*PluginInfo
	// mapping from symbol to plugin. Deduced from plugin set.
	sources   map[string]string
	sinks     map[string]string
	functions map[string]string
}

// Set prerequisite: the pluginInfo must have been validated that the names are valid
func (r *registry) Set(name string, pi *PluginInfo) { _ = "STUB: not implemented"; return }

func (r *registry) Get(name string) (*PluginInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *registry) GetSymbol(pt plugin.PluginType, symbolName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *registry) List() []*PluginInfo { _ = "STUB: not implemented"; return nil }

// return empty slice instead of nil to help json marshal

func (r *registry) Delete(name string) { _ = "STUB: not implemented"; return }
