// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

// Plugin runtime to control the whole plugin with control channel: Distribute symbol data connection, stop symbol and stop plugin

package runtime

import (
	"sync"

	"github.com/lf-edge/ekuiper/sdk/go/api"
)

var (
	logger api.Logger
	reg    runtimes
)

func initVars(args []string, conf *PluginConfig) { _ = "STUB: not implemented"; return }

// parse Args

type (
	NewSourceFunc   func() api.Source
	NewFunctionFunc func() api.Function
	NewSinkFunc     func() api.Sink
)

// PluginConfig construct once and then read only
type PluginConfig struct {
	Name      string
	Sources   map[string]NewSourceFunc
	Functions map[string]NewFunctionFunc
	Sinks     map[string]NewSinkFunc
}

func (conf *PluginConfig) Get(pluginType string, symbolName string) (builderFunc interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// Start Connect to control plane
// Only run once at process startup
func Start(args []string, conf *PluginConfig) { _ = "STUB: not implemented"; return }

// not parallel run now

// never stop a function symbol here.

// Stop the whole plugin

//nolint:staticcheck

// key is rule_op_ins_symbol
type runtimes struct {
	content map[string]RuntimeInstance
	sync.RWMutex
}

func (r *runtimes) Set(name string, instance RuntimeInstance) { _ = "STUB: not implemented"; return }

func (r *runtimes) Get(name string) (RuntimeInstance, bool) {
	_ = "STUB: not implemented"
	return *new(RuntimeInstance), false
}

func (r *runtimes) Delete(name string) { _ = "STUB: not implemented"; return }
