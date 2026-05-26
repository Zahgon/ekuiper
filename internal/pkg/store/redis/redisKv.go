// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

//go:build redisdb || !core

package redis

import (
	"github.com/redis/go-redis/v9"
)

const KvPrefix = "KV:STORE"

type redisKvStore struct {
	database  *redis.Client
	table     string
	keyPrefix string
}

func createRedisKvStore(redis *redis.Client, table string) (*redisKvStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv redisKvStore) Setnx(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv redisKvStore) Set(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv redisKvStore) Get(key string, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (kv redisKvStore) GetKeyedState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv redisKvStore) SetKeyedState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv redisKvStore) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (kv redisKvStore) Keys() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv redisKvStore) GetByPrefix(prefix string) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv redisKvStore) All() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv redisKvStore) metaKeys() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv redisKvStore) Clean() error { _ = "STUB: not implemented"; return nil }

func (kv redisKvStore) Drop() error { _ = "STUB: not implemented"; return nil }

func (kv redisKvStore) tableKey(key string) string { _ = "STUB: not implemented"; return "" }

func (kv redisKvStore) trimPrefix(fullKey string) string { _ = "STUB: not implemented"; return "" }
