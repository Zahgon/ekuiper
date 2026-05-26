// Copyright 2023 EMQ Technologies Co., Ltd.
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

//go:build amd64 && fdb

package fdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

type fdbKvStore struct {
	database *fdb.Database
	subspace directory.DirectorySubspace
}

func createFdbKvStore(fdb *fdb.Database, db string, table string) (*fdbKvStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv fdbKvStore) Setnx(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv fdbKvStore) Set(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv fdbKvStore) GetByPrefix(prefix string) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv fdbKvStore) Get(key string, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (kv fdbKvStore) GetKeyedState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv fdbKvStore) SetKeyedState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv fdbKvStore) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (kv fdbKvStore) Keys() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv fdbKvStore) All() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv fdbKvStore) Clean() error { _ = "STUB: not implemented"; return nil }

func (kv fdbKvStore) Drop() error { _ = "STUB: not implemented"; return nil }
