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

package conf

import (
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// ConfKeysOperator define interface to query/add/update/delete the configs in memory
type ConfKeysOperator interface {
	GetPluginName() string
	GetConfContentByte() ([]byte, error)
	// CopyConfContent get the configurations in etc and data folder
	CopyConfContent() map[string]map[string]interface{}
	// CopyReadOnlyConfContent get the configurations in etc folder
	CopyReadOnlyConfContent() map[string]map[string]interface{}
	// CopyUpdatableConfContent get the configurations in data folder
	CopyUpdatableConfContent() map[string]map[string]interface{}
	// CopyUpdatableConfContentFor get the configuration for the specific configKeys
	CopyUpdatableConfContentFor(configKeys []string) map[string]map[string]interface{}
	// LoadConfContent load the configurations into data configuration part
	LoadConfContent(cf map[string]map[string]interface{})
	GetConfKeys() (keys []string)
	GetReadOnlyConfKeys() (keys []string)
	GetUpdatableConfKeys() (keys []string)
	DeleteConfKey(confKey string)
	DeleteConfKeyField(confKey string, reqField map[string]interface{}) error
	AddConfKey(confKey string, reqField map[string]interface{}) error
	AddConfKeyField(confKey string, reqField map[string]interface{}) error
	ClearConfKeys()
}

// ConfigOperator define interface to query/add/update/delete the configs in disk
type ConfigOperator interface {
	ConfKeysOperator
	SaveCfgToStorage() error
}

// ConfigKeys implement ConfKeysOperator interface, load the configs from etc/sources/xx.yaml and et/connections/connection.yaml
// Hold the connection configs for each connection type in etcCfg field
// Provide method to query/add/update/delete the configs
type ConfigKeys struct {
	storageType string
	lock        syncx.RWMutex
	pluginName  string                            // source type, can be mqtt/edgex/httppull
	etcCfg      map[string]map[string]interface{} // configs defined in etc/sources/yaml
	dataCfg     map[string]map[string]interface{}
	// delCfgKey save the config key which needs to be deleted from the storage
	delCfgKey map[string]struct{}
	// saveCfgKey save the config key which needs to be saved or updated into the storage
	saveCfgKey map[string]struct{}
}

func (c *ConfigKeys) saveCfgKeysIntoKVStorage(cfgType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) GetPluginName() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigKeys) GetConfContentByte() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConfigKeys) CopyConfContent() map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// note: config keys in data directory will overwrite those in etc directory with same name

func (c *ConfigKeys) LoadConfContent(cf map[string]map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *ConfigKeys) CopyReadOnlyConfContent() map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) CopyUpdatableConfContent() map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) CopyUpdatableConfContentFor(configKeys []string) map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) GetConfKeys() (keys []string) { _ = "STUB: not implemented"; return nil }

func (c *ConfigKeys) GetReadOnlyConfKeys() (keys []string) { _ = "STUB: not implemented"; return nil }

func (c *ConfigKeys) GetUpdatableConfKeys() (keys []string) { _ = "STUB: not implemented"; return nil }

func (c *ConfigKeys) DeleteConfKey(confKey string) { _ = "STUB: not implemented"; return }

func (c *ConfigKeys) ClearConfKeys() { _ = "STUB: not implemented"; return }

func recursionDelMap(cf, fields map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) DeleteConfKeyField(confKey string, reqField map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) AddConfKey(confKey string, reqField map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigKeys) AddConfKeyField(confKey string, reqField map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// SourceConfigKeysOps implement ConfOperator interface, load the configs from etc/sources/xx.yaml
type SourceConfigKeysOps struct {
	*ConfigKeys
}

func (c *SourceConfigKeysOps) SaveCfgToStorage() error { _ = "STUB: not implemented"; return nil }

// SinkConfigKeysOps implement ConfOperator interface, load the configs from data/sinks/xx.yaml
type SinkConfigKeysOps struct {
	*ConfigKeys
}

func (c *SinkConfigKeysOps) SaveCfgToStorage() error { _ = "STUB: not implemented"; return nil }

// ConnectionConfigKeysOps implement ConfOperator interface, load the configs from et/connections/connection.yaml
type ConnectionConfigKeysOps struct {
	*ConfigKeys
}

func (p *ConnectionConfigKeysOps) SaveCfgToStorage() error { _ = "STUB: not implemented"; return nil }

// NewConfigOperatorForSource construct function
func NewConfigOperatorForSource(pluginName string) ConfigOperator {
	_ = "STUB: not implemented"
	return *new(ConfigOperator)
}

// NewConfigOperatorFromSourceStorage construct function, Load the configs from etc/sources/xx.yaml
func NewConfigOperatorFromSourceStorage(pluginName string) (ConfigOperator, error) {
	_ = "STUB: not implemented"
	return *new(ConfigOperator), nil
}

// Just ignore error if yaml not found

// NewConfigOperatorForSink construct function
func NewConfigOperatorForSink(pluginName string) ConfigOperator {
	_ = "STUB: not implemented"
	return *new(ConfigOperator)
}

// NewConfigOperatorFromSinkStorage construct function, Load the configs from etc/sources/xx.yaml
func NewConfigOperatorFromSinkStorage(pluginName string) (ConfigOperator, error) {
	_ = "STUB: not implemented"
	return *new(ConfigOperator), nil
}

// NewConfigOperatorForConnection construct function
func NewConfigOperatorForConnection(pluginName string) ConfigOperator {
	_ = "STUB: not implemented"
	return *new(ConfigOperator)
}

// NewConfigOperatorFromConnectionStorage construct function, Load the configs from et/connections/connection.yaml
func NewConfigOperatorFromConnectionStorage(pluginName string) (ConfigOperator, error) {
	_ = "STUB: not implemented"
	return *new(ConfigOperator), nil
}

func getStorageType() string { _ = "STUB: not implemented"; return "" }
