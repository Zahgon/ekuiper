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

package service

import (
	"context"
	"sync"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/binder"
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var (
	once      sync.Once
	mutex     syncx.Mutex
	singleton *Manager           // Do not call this directly, use GetServiceManager
	_         binder.FuncFactory = singleton
)

type Manager struct {
	executorPool *sync.Map // The pool of executors
	loaded       bool
	serviceBuf   *sync.Map
	functionBuf  *sync.Map

	etcDir                 string
	serviceInstallKV       kv.KeyValue
	serviceStatusInstallKV kv.KeyValue
	serviceKV              kv.KeyValue
	functionKV             kv.KeyValue
}

func InitManager() (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

// To boost the testing perf

func GetManager() *Manager {
	_ = "STUB: not implemented"

	// InitByFiles
	/**
	 * This function will parse the service definition json files in etc/services.
	 * It will validate all json files and their schemaFiles. If invalid, it just prints
	 * an error log and ignore. So it is possible that only valid service definition are
	 * parsed and available.
	 *
	 * NOT threadsafe, must run in lock
	 */return nil
}

func (m *Manager) InitByFiles() error { _ = "STUB: not implemented"; return nil }

// Parse schemas in batch. So we have 2 loops. First loop to collect files and the second to save the result.

func (m *Manager) initFile(baseName string) error { _ = "STUB: not implemented"; return nil }

// TODO validate serviceConf

// setting function alias

// Start Implement FunctionFactory

func (m *Manager) HasFunctionSet(_ string) bool { _ = "STUB: not implemented"; return false }

func (m *Manager) Function(name string) (api.Function, error) {
	_ = "STUB: not implemented"
	return *new(api.Function), nil
}

// executor is gotten from pool, so all externalFuncs with the same interface share the same executor instance

func (m *Manager) FunctionPluginInfo(funcName string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func (m *Manager) ConvName(funcName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// End Implement FunctionFactory

func (m *Manager) HasService(name string) bool { _ = "STUB: not implemented"; return false }

func (m *Manager) getFunction(name string) (*functionContainer, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *Manager) getService(name string) (*serviceInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Each interface maps to an executor
func (m *Manager) getExecutor(name string, info *interfaceInfo) (executor, error) {
	_ = "STUB: not implemented"
	return *new(executor), nil
}

func (m *Manager) deleteServiceFuncs(service string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) deleteFunc(service, name string) error { _ = "STUB: not implemented"; return nil }

// ** CRUD of the service files **

type ServiceCreationRequest struct {
	Name string `json:"name" yaml:"name"`
	File string `json:"file" yaml:"file"`
}

func (s *ServiceCreationRequest) InstallScript() string { _ = "STUB: not implemented"; return "" }

func (m *Manager) List() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Manager) Create(r *ServiceCreationRequest) error { _ = "STUB: not implemented"; return nil }

// download

// clean up: delete zip file and unzip files in error

// unzip and copy to destination

// save the install script

// init file to serviceKV

func (m *Manager) Delete(name string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) Get(name string) (*serviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) Update(req *ServiceCreationRequest) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) unzip(name, src string) error { _ = "STUB: not implemented"; return nil }

// Try unzip

// unzip

func (m *Manager) ListFunctions() ([]*functionContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) GetFunction(name string) (*functionContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) GetAllServices() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) GetAllServicesStatus() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) UninstallAllServices() { _ = "STUB: not implemented"; return }

func (m *Manager) servicesRegisterForImport(_, v string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) ImportServices(ctx context.Context, services map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) ImportPartialServices(ctx context.Context, services map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
