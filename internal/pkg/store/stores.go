// Copyright 2021-2023 EMQ Technologies Co., Ltd.
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

package store

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/store/definition"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/store/sql"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type StoreCreator func(conf definition.Config, name string) (definition.StoreBuilder, definition.TsBuilder, error)

var (
	storeBuilders = map[string]StoreCreator{
		"sqlite": sql.BuildStores,
	}
	globalStores     *stores = nil
	cacheStores      *stores = nil
	extStateStores   *stores = nil
	checkpointStores *stores = nil

	TraceStores sql.Database
)

type stores struct {
	kv        map[string]kv.KeyValue
	ts        map[string]kv.Tskv
	mu        syncx.Mutex
	kvBuilder definition.StoreBuilder
	tsBuilder definition.TsBuilder
}

func newStores(c definition.Config, name string) (*stores, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newExtStateStores(c definition.Config, name string) (*stores, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *stores) GetKV(table string) (kv.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(kv.KeyValue), nil
}

func (s *stores) DropKV(table string) { _ = "STUB: not implemented"; return }

func (s *stores) DropRefKVs(tablePrefix string) { _ = "STUB: not implemented"; return }

func (s *stores) GetTS(table string) (kv.Tskv, error) {
	_ = "STUB: not implemented"
	return *new(kv.Tskv), nil
}

func (s *stores) DropTS(table string) { _ = "STUB: not implemented"; return }

func GetKV(table string) (kv.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(kv.KeyValue), nil
}

func GetTS(table string) (kv.Tskv, error) { _ = "STUB: not implemented"; return *new(kv.Tskv), nil }

func DropTS(table string) error { _ = "STUB: not implemented"; return nil }

func DropKV(table string) error { _ = "STUB: not implemented"; return nil }

func GetCacheKV(table string) (kv.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(kv.KeyValue), nil
}

func DropCacheKV(table string) error { _ = "STUB: not implemented"; return nil }

func DropCacheKVForRule(rule string) error { _ = "STUB: not implemented"; return nil }

func GetExtStateKV(table string) (kv.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(kv.KeyValue), nil
}
