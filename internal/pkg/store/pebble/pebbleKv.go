// Copyright 2025 EMQ Technologies Co., Ltd.
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

package pebble

import (
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

type pebbleKvStore struct {
	database KVDatabase
	table    string
}

func createPebbleKvStore(database KVDatabase, table string) (kv.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(kv.KeyValue), nil
}

func (p *pebbleKvStore) key(k string) []byte { _ = "STUB: not implemented"; return nil }

func (p *pebbleKvStore) Setnx(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pebbleKvStore) Set(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pebbleKvStore) Get(key string, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pebbleKvStore) GetKeyedState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pebbleKvStore) SetKeyedState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pebbleKvStore) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (p *pebbleKvStore) Keys() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pebbleKvStore) All() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pebbleKvStore) Clean() error { _ = "STUB: not implemented"; return nil }

func (p *pebbleKvStore) Drop() error { _ = "STUB: not implemented"; return nil }

func (p *pebbleKvStore) GetByPrefix(prefix string) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
