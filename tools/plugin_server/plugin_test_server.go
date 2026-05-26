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

package main

import (
	"net/http"
	"sync"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/internal/plugin/portable"
	"github.com/lf-edge/ekuiper/v2/internal/plugin/portable/runtime"
	"github.com/lf-edge/ekuiper/v2/internal/topo/context"
	"github.com/lf-edge/ekuiper/v2/internal/topo/state"
)

// Only support to test a single plugin Testing process.
// 0. Edit the testingPlugin variable to match your plugin meta.
// 1. Start this server, and wait for handshake.
// 2. Start or debug your plugin. Make sure the handshake completed.
// 3. Issue startSymbol/stopSymbol REST API to debug your plugin symbol.

// EDIT HERE: Define the plugins that you want to test.
var testingPlugin = &portable.PluginInfo{
	PluginMeta: runtime.PluginMeta{
		Name:       "pysam",
		Version:    "v1",
		Language:   "python",
		Executable: "pysam.py",
	},
	Sources:   []string{"pyjson"},
	Sinks:     []string{"print"},
	Functions: []string{"revert"},
}

var mockSinkData = []map[string]interface{}{
	{
		"name":  "hello",
		"count": 5,
	}, {
		"name":  "world",
		"count": 10,
	},
}

var mockFuncData = [][]interface{}{
	{"twelve"},
	{"eleven"},
}

var (
	m       *portable.Manager
	ctx     api.StreamContext
	cancels sync.Map
)

func main() {
	var err error
	m, err = portable.MockManager(map[string]*portable.PluginInfo{testingPlugin.Name: testingPlugin})
	if err != nil {
		panic(err)
	}
	ins, err := startPluginIns(testingPlugin)
	if err != nil {
		panic(err)
	}
	defer ins.Stop()
	runtime.GetPluginInsManager().AddPluginIns(testingPlugin.Name, ins)
	c := context.WithValue(context.Background(), context.LoggerKey, conf.Log)
	ctx = c.WithMeta("rule1", "op1", &state.MemoryStore{}).WithInstance(1)
	server := createRestServer("127.0.0.1", 33333)
	server.ListenAndServe()
}

func startPluginIns(info *portable.PluginInfo) (*runtime.PluginIns, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createRestServer(ip string, port int) *http.Server { _ = "STUB: not implemented"; return nil }

// Good practice to set timeouts to avoid Slowloris attacks.

func startSymbolHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func stopSymbolHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func decode(r *http.Request) (*runtime.Control, error) { _ = "STUB: not implemented"; return nil, nil }
