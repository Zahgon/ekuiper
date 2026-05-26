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

type pebbleTsStore struct {
	database KVDatabase
	table    string
	last     int64
}

func createPebbleTs(database KVDatabase, table string) (kv.Tskv, error) {
	_ = "STUB: not implemented"
	return *new(kv.Tskv), nil
}

func (t *pebbleTsStore) key(k int64) []byte { _ = "STUB: not implemented"; return nil }

func (t *pebbleTsStore) Set(key int64, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *pebbleTsStore) Get(key int64, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *pebbleTsStore) Last(value interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *pebbleTsStore) Delete(key int64) error { _ = "STUB: not implemented"; return nil }

func (t *pebbleTsStore) DeleteBefore(key int64) error { _ = "STUB: not implemented"; return nil }

func (t *pebbleTsStore) Close() error { _ = "STUB: not implemented"; return nil }

func (t *pebbleTsStore) Drop() error { _ = "STUB: not implemented"; return nil }

func getLastTs(d KVDatabase, table string) int64 { _ = "STUB: not implemented"; return 0 }
