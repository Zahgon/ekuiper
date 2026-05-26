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

package portable

import (
	"context"

	"github.com/lf-edge/ekuiper/v2/internal/binder"
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/internal/plugin/portable/runtime"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

var (
	manager *Manager
	_       binder.SourceFactory = manager
	_       binder.SinkFactory   = manager
	_       binder.FuncFactory   = manager
)

type Manager struct {
	pluginDir     string
	pluginConfDir string
	pluginDataDir string
	reg           *registry // can be replaced with kv
	// the access to plugin install script db
	plgInstallDb kv.KeyValue
	// the access to plugin install status db
	plgStatusDb kv.KeyValue
}

// InitManager must only be called once
func InitManager() (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

// Read plugin info from file system

func MockManager(plugins map[string]*PluginInfo) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) syncRegistry() (err error) { _ = "STUB: not implemented"; return nil }

// Parse plugins asyncly because it will run the plugin which may block in handshake

func (m *Manager) parsePlugin(name string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) doRegister(name string, pi *PluginInfo, isInit bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO install async? This may take time at start up

func (m *Manager) parsePluginJson(name string) (info *PluginInfo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) storePluginInstallScript(name string, j plugin.Plugin) {
	_ = "STUB: not implemented"
	return
}

func (m *Manager) removePluginInstallScript(name string) { _ = "STUB: not implemented"; return }

func (m *Manager) Register(p plugin.Plugin) error { _ = "STUB: not implemented"; return nil }

// download

// clean up: delete zip file and unzip files in error

// unzip and copy to destination

// Revert for any errors

func (m *Manager) install(name, src string, shellParas []string) (resultErr error) {
	_ = "STUB: not implemented"
	return nil
}

// The map of install files. Used to check if all required files are installed and for reverting

// remove all installed files if err happens

// Parse json file

// file copying

// Check if all files installed

// run install script if there is

func (m *Manager) List() []*PluginInfo { _ = "STUB: not implemented"; return nil }

func (m *Manager) GetPluginMeta(pt plugin.PluginType, symbolName string) (*runtime.PluginMeta, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *Manager) GetPluginInfo(pluginName string) (*PluginInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *Manager) Delete(name string) error { _ = "STUB: not implemented"; return nil }

// unregister the plugin

// delete files and uninstall metas

// Kill the process in the end, and return error if it cannot be deleted

func (m *Manager) UninstallAllPlugins() { _ = "STUB: not implemented"; return }

func (m *Manager) GetAllPlugins() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) GetAllPluginsStatus() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) pluginRegisterForImport(k, v string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) PluginImport(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) PluginPartialImport(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
