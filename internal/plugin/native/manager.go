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

// Manage the loading of both native and portable plugins

package native

import (
	"context"
	"plugin"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/binder"
	plugin2 "github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// isSafeArchiveEntry checks if the entry name is safe for extraction
func isSafeArchiveEntry(name string) bool {
	_ = "STUB: not implemented"
	// Disallow absolute paths
	return false
}

// Disallow parent traversal

// Prevent special device files (on Windows)

// Manager Initialized in the binder
var (
	manager *Manager
	_       binder.SourceFactory = manager
	_       binder.SinkFactory   = manager
	_       binder.FuncFactory   = manager
)

const DELETED = "$deleted"

// Manager is appended only because plugin cannot delete or reload. To delete a plugin, restart the server to reindex
type Manager struct {
	syncx.RWMutex
	// 3 maps for source/sink/function. In each map, key is the plugin name, value is the version
	plugins []map[string]string
	// A map from function name to its plugin file name. It is constructed during initialization by reading kv info. All functions must have at least an entry, even the function resizes in a one function plugin.
	symbols map[string]string
	// loaded symbols in current runtime
	runtime map[string]*plugin.Plugin
	// dirs
	pluginDir     string
	pluginConfDir string
	pluginDataDir string
	// the access to func symbols db
	funcSymbolsDb kv.KeyValue
	// the access to plugin install script db
	plgInstallDb kv.KeyValue
	// the access to plugin install status db
	plgStatusDb kv.KeyValue
}

// InitManager must only be called once
func InitManager() (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

func findAll(t plugin2.PluginType, pluginDir string) (result map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// load the plugins when ekuiper set up

func (rr *Manager) get(t plugin2.PluginType, name string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (rr *Manager) store(t plugin2.PluginType, name string, version string) {
	_ = "STUB: not implemented"
	return
}

func (rr *Manager) storeSymbols(name string, symbols []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rr *Manager) removeSymbols(symbols []string) { _ = "STUB: not implemented"; return }

// API for management

func (rr *Manager) List(t plugin2.PluginType) []string { _ = "STUB: not implemented"; return nil }

func (rr *Manager) ListSymbols() []string { _ = "STUB: not implemented"; return nil }

func (rr *Manager) GetPluginVersionBySymbol(t plugin2.PluginType, symbolName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (rr *Manager) GetPluginBySymbol(t plugin2.PluginType, symbolName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (rr *Manager) storePluginInstallScript(name string, t plugin2.PluginType, j plugin2.Plugin) {
	_ = "STUB: not implemented"
	return
}

func (rr *Manager) removePluginInstallScript(name string, t plugin2.PluginType) {
	_ = "STUB: not implemented"
	return
}

func (rr *Manager) Register(t plugin2.PluginType, j plugin2.Plugin) error {
	_ = "STUB: not implemented"
	return nil
}

// Validation

// download

// clean up: delete zip file and unzip files in error

// unzip and copy to destination

// Revert for any errors

// RegisterFuncs prerequisite：function plugin of name exists
func (rr *Manager) RegisterFuncs(name string, functions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rr *Manager) Delete(t plugin2.PluginType, name string, stop bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Find etc folder

// Find etc folder

func (rr *Manager) GetPluginInfo(t plugin2.PluginType, name string) (map[string]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// ignore the error

func (rr *Manager) install(t plugin2.PluginType, name, src string, shellParas []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Prevent Zip Slip: only allow safe archive entries

// skip yaml file if exists

// run install script if there is

// load the runtime first

// binder factory implementations

func (rr *Manager) Source(name string) (api.Source, error) {
	_ = "STUB: not implemented"
	return *new(api.Source), nil
}

func (rr *Manager) SourcePluginInfo(name string) (plugin2.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin2.EXTENSION_TYPE), "", ""
}

func (rr *Manager) LookupSource(name string) (api.Source, error) {
	_ = "STUB: not implemented"
	return *new(api.Source), nil
}

func (rr *Manager) Sink(name string) (api.Sink, error) {
	_ = "STUB: not implemented"
	return *new(api.Sink), nil
}

func (rr *Manager) SinkPluginInfo(name string) (plugin2.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin2.EXTENSION_TYPE), "", ""
}

func (rr *Manager) Function(name string) (api.Function, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), nil
}

func (rr *Manager) HasFunctionSet(name string) bool { _ = "STUB: not implemented"; return false }

func (rr *Manager) FunctionPluginInfo(funcName string) (plugin2.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin2.EXTENSION_TYPE), "", ""
}

func (rr *Manager) ConvName(name string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// If not found, return nil,nil; Other errors return nil, err
func (rr *Manager) loadRuntime(t plugin2.PluginType, soName, soFilepath, symbolName string) (plugin.Symbol, error) {
	_ = "STUB: not implemented"
	return *new(plugin.Symbol), nil
}

// Return the lowercase version of so name. It may be upper case in path.
func (rr *Manager) getSoFilePath(t plugin2.PluginType, name string, isSoName bool) (string, error) {
	_ = "STUB: not implemented"
	// Validate plugin name to prevent path traversal or absolute paths
	return "", nil
}

// We must identify plugin or symbol when deleting function plugin

func parseName(n string) (string, string) { _ = "STUB: not implemented"; return "", "" }

func ucFirst(str string) string { _ = "STUB: not implemented"; return "" }

func lcFirst(str string) string { _ = "STUB: not implemented"; return "" }

func (rr *Manager) UninstallAllPlugins() { _ = "STUB: not implemented"; return }

func (rr *Manager) GetAllPlugins() map[string]string { _ = "STUB: not implemented"; return nil }

func (rr *Manager) GetAllPluginsStatus() map[string]string { _ = "STUB: not implemented"; return nil }

const BOOT_INSTALL = "$boot_install"

// PluginImport save the plugin install information and wait for restart
func (rr *Manager) PluginImport(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// set the flag to install the plugins when eKuiper reboot

// PluginPartialImport compare the plugin to be installed and the one in database
// if not exist in database, install;
// if exist, ignore
func (rr *Manager) PluginPartialImport(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (rr *Manager) hasInstallFlag() bool { _ = "STUB: not implemented"; return false }

func (rr *Manager) clearInstallFlag() { _ = "STUB: not implemented"; return }

func (rr *Manager) pluginRegisterForImport(key, script string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rr *Manager) pluginInstallWhenReboot() { _ = "STUB: not implemented"; return }
