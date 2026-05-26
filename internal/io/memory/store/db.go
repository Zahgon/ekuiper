// Copyright 2022-2023 EMQ Technologies Co., Ltd.
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
	"context"

	"github.com/lf-edge/ekuiper/v2/internal/io/memory/pubsub"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type tableCount struct {
	syncx.RWMutex
	count int
	t     *Table
}

func (tc *tableCount) Increase() int { _ = "STUB: not implemented"; return 0 }

func (tc *tableCount) Decrease() int { _ = "STUB: not implemented"; return 0 }

type database struct {
	syncx.RWMutex
	tables map[string]*tableCount // topic_key: table
}

// getTable return the table of the topic.
func (db *database) getTable(topic string, key string) (*Table, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// addTable add a table to the database
// If the table already exists, return the existing table;
// otherwise, create a new table and return it.
// The second argument is to indicate if the table is newly created
func (db *database) addTable(topic string, key string) (*Table, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// dropTable drop the table of the topic/values
// stops to accumulate job
// deletes the cache data
func (db *database) dropTable(topic string, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Table has one writer and multiple reader
type Table struct {
	syncx.RWMutex
	topic string
	key   string
	// datamap is the overall data indexed by primary key
	datamap map[any]pubsub.MemTuple
	cancel  context.CancelFunc
}

func createTable(topic string, key string) *Table { _ = "STUB: not implemented"; return nil }

func (t *Table) add(value pubsub.MemTuple) { _ = "STUB: not implemented"; return }

func (t *Table) delete(key interface{}) { _ = "STUB: not implemented"; return }

func (t *Table) setCancel(cancel context.CancelFunc) { _ = "STUB: not implemented"; return }

func (t *Table) callCancel() { _ = "STUB: not implemented"; return }

func (t *Table) Read(keys []string, values []interface{}) ([]pubsub.MemTuple, error) {
	_ = "STUB: not implemented"
	return nil,

		// Find the primary key
		nil
}

var db = &database{
	tables: make(map[string]*tableCount),
}
