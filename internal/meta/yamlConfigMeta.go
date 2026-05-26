// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

package meta

import (
	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type configManager struct {
	lock                     syncx.RWMutex
	cfgOperators             map[string]conf.ConfigOperator
	sourceConfigStatusDb     kv.KeyValue
	sinkConfigStatusDb       kv.KeyValue
	connectionConfigStatusDb kv.KeyValue
}

// ConfigManager Hold the ConfigOperator for yaml configs defined in etc/sources/xxx.yaml and etc/connections/connection.yaml
// for configs in etc/sources/xxx.yaml, the map key is sources.xxx format, xxx will be mqtt/httppull and so on
// for configs in etc/connections/connection.yaml, the map key is connections.xxx format, xxx will be mqtt/edgex
var ConfigManager *configManager

func InitYamlConfigManager() { _ = "STUB: not implemented"; return }

const (
	SourceCfgOperatorKeyTemplate     = "sources.%s"
	SourceCfgOperatorKeyPrefix       = "sources."
	SinkCfgOperatorKeyTemplate       = "sinks.%s"
	SinkCfgOperatorKeyPrefix         = "sinks."
	ConnectionCfgOperatorKeyTemplate = "connections.%s"
	ConnectionCfgOperatorKeyPrefix   = "connections."
)

// loadConfigOperatorForSource
// Try to load ConfigOperator for plugin xxx from /etc/sources/xxx.yaml  /data/sources/xxx.yaml
// If plugin xxx not exist, no error response
func loadConfigOperatorForSource(pluginName string) { _ = "STUB: not implemented"; return }

// loadConfigOperatorForSink
// Try to load ConfigOperator for plugin xxx from /data/sinks/xxx.yaml
// If plugin xxx not exist, no error response
func loadConfigOperatorForSink(pluginName string) { _ = "STUB: not implemented"; return }

// loadConfigOperatorForConnection
// Try to load ConfigOperator for plugin from /etc/connections/connection.yaml /data/connections/connection.yaml
// If plugin not exist in /etc/connections/connection.yaml, no error response
func loadConfigOperatorForConnection(pluginName string) { _ = "STUB: not implemented"; return }

func delConfKey(configOperatorKey, confKey, language string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func DelSourceConfKey(plgName, confKey, language string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func DelSinkConfKey(plgName, confKey, language string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func DelConnectionConfKey(plgName, confKey, language string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func delYamlConf(configOperatorKey string) { _ = "STUB: not implemented"; return }

func GetConfOperator(configOperatorKey string) (conf.ConfigOperator, bool) {
	_ = "STUB: not implemented"
	return *new(conf.ConfigOperator), false
}

func GetSourceResourceConf(sourceType string) map[string]map[string]map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// specific sourceType

func appendConfKeyInResult(result map[string]map[string]map[string]interface{}, typ, confKey string, confValue map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func GetYamlConf(configOperatorKey, language string) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceConfigurations(plg string, cf YamlConfigurations) YamlConfigurations {
	_ = "STUB: not implemented"
	return *new(YamlConfigurations)
}

func addSourceConfKeys(plgName string, configurations YamlConfigurations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddSourceConfKey(plgName, confKey, language string, content []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateConf(pluginName string, props map[string]interface{}, isSource bool) error {
	_ = "STUB: not implemented"
	return nil
}

func AddSinkConfKey(plgName, confKey, language string, content []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func addSinkConfKeys(plgName string, cf YamlConfigurations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddConnectionConfKey(plgName, confKey, language string, reqField map[string]interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func addConnectionConfKeys(plgName string, cf YamlConfigurations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func GetResources(language string) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ResetConfigs() { _ = "STUB: not implemented"; return }

type YamlConfigurations map[string]map[string]interface{}

type YamlConfigurationSet struct {
	Sources     map[string]string `json:"sources"`
	Sinks       map[string]string `json:"sinks"`
	Connections map[string]string `json:"connections"`
}

func GetConfigurations() YamlConfigurationSet {
	_ = "STUB: not implemented"
	return *new(YamlConfigurationSet)
}

type YamlConfigurationKeys struct {
	Sources map[string][]string
	Sinks   map[string][]string
}

func GetConfigurationsFor(yaml YamlConfigurationKeys) YamlConfigurationSet {
	_ = "STUB: not implemented"
	return *new(YamlConfigurationSet)
}

func GetConfigurationStatus() YamlConfigurationSet {
	_ = "STUB: not implemented"
	return *new(YamlConfigurationSet)
}

func LoadConfigurations(configSets YamlConfigurationSet) YamlConfigurationSet {
	_ = "STUB: not implemented"
	return *new(YamlConfigurationSet)
}

func LoadConfigurationsPartial(configSets YamlConfigurationSet) YamlConfigurationSet {
	_ = "STUB: not implemented"
	return *new(YamlConfigurationSet)
}
