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

//go:build plugin || !core

package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/internal/plugin/native"
)

var nativeManager *native.Manager

func init() {
	components["plugin"] = pluginComp{}
}

type pluginComp struct{}

func (p pluginComp) register() { _ = "STUB: not implemented"; return }

func (p pluginComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

func (p pluginComp) exporter() ConfManager { _ = "STUB: not implemented"; return *new(ConfManager) }

func pluginsHandler(w http.ResponseWriter, r *http.Request, t plugin.PluginType) {
	_ = "STUB: not implemented"
	return
}

// Problems decoding

func pluginHandler(w http.ResponseWriter, r *http.Request, t plugin.PluginType) {
	_ = "STUB: not implemented"
	return
}

// Problems decoding

// list or create source plugin
func sourcesHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// delete a source plugin
func sourceHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list or create sink plugin
func sinksHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// delete a sink plugin
func sinkHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list or create function plugin
func functionsHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list all user-defined functions in all function plugins
func functionsListHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func functionsGetHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// delete a function plugin
func functionHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type functionList struct {
	Functions []string `json:"functions,omitempty"`
}

// Register function list for function plugin. If a plugin exports multiple functions, the function list must be registered
// either by create or register. If the function plugin has been loaded because of auto load through so file, the function
// list MUST be registered by this API or only the function with the same name as the plugin can be used.
func functionRegisterHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Problems decoding

func prebuildSourcePlugins(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func prebuildSinkPlugins(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func prebuildFuncsPlugins(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func prebuildPluginsHandler(w http.ResponseWriter, _ *http.Request, t plugin.PluginType) {
	_ = "STUB: not implemented"
	return
}

func fetchPluginList(t plugin.PluginType, hosts, os, arch string) (result map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The url is similar to http://host:port/kuiper-plugins/0.9.1/debian/sinks/

type pluginExporter struct{}

func (e pluginExporter) Import(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e pluginExporter) PartialImport(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e pluginExporter) Export() map[string]string { _ = "STUB: not implemented"; return nil }

func (e pluginExporter) Status() map[string]string { _ = "STUB: not implemented"; return nil }

func (e pluginExporter) Reset() { _ = "STUB: not implemented"; return }
