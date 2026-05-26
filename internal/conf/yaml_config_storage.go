// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

const (
	cfgStoreKVStorage = "kv"
)

type cfgKVStorage interface {
	Set(string, map[string]interface{}) error
	Delete(string) error
	GetByPrefix(string) (map[string]map[string]interface{}, error)
}

type kvMemory struct {
	store map[string]map[string]interface{}
}

func (m *kvMemory) Set(key string, v map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *kvMemory) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (m *kvMemory) GetByPrefix(prefix string) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	mockMemoryKVStore *kvMemory
	kvStore           *sqlKVStore
)

// GetYamlConfigAllKeys get all plugin keys about sources/sinks/connections
func GetYamlConfigAllKeys(typ string) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getKVStorage() (s cfgKVStorage, err error) {
	_ = "STUB: not implemented"
	return *new(cfgKVStorage), nil
}

// SaveCfgKeyToKV ...
func SaveCfgKeyToKV(key string, cfg map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadCfgKeyKV(key string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func saveCfgKeyToKV(key string, cfg map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func delCfgKeyInStorage(key string) error { _ = "STUB: not implemented"; return nil }

func getCfgKeyFromStorageByPrefix(prefix string) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip data if not conf

func buildKey(confType string, pluginName string, confKey string) string {
	_ = "STUB: not implemented"
	return ""
}

type sqlKVStore struct {
	kv kv.KeyValue
}

func NewSqliteKVStore(table string) (*sqlKVStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlKVStore) Set(k string, v map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlKVStore) Delete(k string) error { _ = "STUB: not implemented"; return nil }

func (s *sqlKVStore) GetByPrefix(prefix string) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteCfgIntoKVStorage ...
func WriteCfgIntoKVStorage(typ string, plugin string, confKey string, confData map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DropCfgKeyFromStorage ...
func DropCfgKeyFromStorage(typ string, plugin string, confKey string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCfgFromKVStorage ...
func GetCfgFromKVStorage(typ string, plugin string, confKey string) (map[string]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClearKVStorage only used in unit test
func ClearKVStorage() error { _ = "STUB: not implemented"; return nil }

// GetAllConnConfigs return connections' plugin -> confKey -> props
func GetAllConnConfigs() (map[string]map[string]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitKey(key string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}
