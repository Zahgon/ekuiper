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

package runtime

import (
	"os"
	"sync"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var (
	once sync.Once
	pm   *pluginInsManager
)

// TODO setting configuration
var PortbleConf = &PortableConfig{
	SendTimeout: 1000,
}

// PluginIns created at two scenarios
// 1. At runtime, plugin is created/updated: in order to be able to reload rules that already uses previous ins
// 2. At system start/restart
// Once created, never deleted until system shutdown
type PluginIns struct {
	syncx.RWMutex
	name     string
	ctrlChan ControlChannel // the same lifecycle as pluginIns, once created keep listening
	// audit the commands, so that when restarting the plugin, we can replay the commands
	commands map[Meta][]byte
	process  *os.Process // created when used by rule and deleted when delete the plugin
	Status   *PluginStatus
}

func NewPluginIns(name string, ctrlChan ControlChannel, process *os.Process) *PluginIns {
	_ = "STUB: not implemented"
	return nil
}

func NewPluginInsForTest(name string, ctrlChan ControlChannel) *PluginIns {
	_ = "STUB: not implemented"
	return nil
}

func (i *PluginIns) sendCmd(jsonArg []byte) error { _ = "STUB: not implemented"; return nil }

func (i *PluginIns) StartSymbol(ctx api.StreamContext, ctrl *Control) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *PluginIns) addRef(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (i *PluginIns) deRef(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (i *PluginIns) StopSymbol(ctx api.StreamContext, ctrl *Control) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop intentionally
func (i *PluginIns) Stop() error { _ = "STUB: not implemented"; return nil }

func (i *PluginIns) GetStatus() *PluginStatus { _ = "STUB: not implemented"; return nil }

// Manager plugin process and control socket
type pluginInsManager struct {
	instances map[string]*PluginIns
	syncx.RWMutex
}

func GetPluginInsManager() *pluginInsManager { _ = "STUB: not implemented"; return nil }

func (p *pluginInsManager) GetPluginInsStatus(name string) (*PluginStatus, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *pluginInsManager) getPluginIns(name string) (*PluginIns, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// AddPluginIns For mock only
func (p *pluginInsManager) AddPluginIns(name string, ins *PluginIns) {
	_ = "STUB: not implemented"
	return
}

// CreateIns Run when plugin is created/updated
func (p *pluginInsManager) CreateIns(pluginMeta *PluginMeta) (*PluginIns, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetOrStartProcess Control the plugin process lifecycle.
// Need to manage the resources: instances map, control socket, plugin process
// May be called at plugin creation or restart with previous state(ctrlCh, commands)
// PluginIns is created by plugin manager and started immediately or restart by rule/funcop.
// The ins is long running. Even for plugin delete/update, the ins will continue. So there is no delete.
// 1. During creation, clean up those resources for any errors in defer immediately after the resource is created.
// 2. During plugin running, when detecting plugin process exit, clean up those resources for the current ins.
func (p *pluginInsManager) GetOrStartProcess(pluginMeta *PluginMeta, pconf *PortableConfig) (_ *PluginIns, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// run initialization for firstly creating plugin instance

// ins has run

// should only happen for first start, then the ctrl channel will keep running

// init or restart all need to run the process

// just print out error inside

// must make sure the plugin ins is not cleaned up yet by checking the process identity
// clean up for stop unintentionally

// restore symbols by sending commands when restarting plugin

func (p *pluginInsManager) Kill(name string) error { _ = "STUB: not implemented"; return nil }

func (p *pluginInsManager) KillAll() error { _ = "STUB: not implemented"; return nil }

type PluginMeta struct {
	Name        string  `json:"name"`
	Version     string  `json:"version"`
	Language    string  `json:"language"`
	Executable  string  `json:"executable"`
	VirtualType *string `json:"virtualEnvType,omitempty"`
	Env         *string `json:"env,omitempty"`
}

const (
	PluginStatusRunning = "running"
	PluginStatusInit    = "initializing"
	PluginStatusErr     = "error"
	PluginStatusStop    = "stop"
)

type PluginStatus struct {
	RefCount map[string]int `json:"refCount"`
	Status   string         `json:"status"`
	ErrMsg   string         `json:"errMsg"`
}

func NewPluginStatus() *PluginStatus { _ = "STUB: not implemented"; return nil }

func (s *PluginStatus) StatusErr(err error) { _ = "STUB: not implemented"; return }

func (s *PluginStatus) StartRunning() { _ = "STUB: not implemented"; return }

func (s *PluginStatus) Stop() { _ = "STUB: not implemented"; return }

func (s *PluginStatus) GetRuleRefCount(rule string) int { _ = "STUB: not implemented"; return 0 }
