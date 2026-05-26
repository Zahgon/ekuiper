// Copyright 2021 EMQ Technologies Co., Ltd.
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

package state

import (
	"encoding/gob"
	"sync"

	"github.com/lf-edge/ekuiper/v2/internal/topo/checkpoint"
	ts2 "github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/store"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register(checkpoint.BufferOrEvent{})
	gob.Register(&store.IndexFieldStore{})
}

// KVStore The manager for checkpoint storage.
//
// mapStore keys
//
//	{ "checkpoint1", "checkpoint2" ... "checkpointn" : The complete or incomplete snapshot
type KVStore struct {
	db          ts2.Tskv
	mapStore    *sync.Map // The current root store of a rule
	checkpoints []int64
	max         int
	ruleId      string
}

// Store in path ./data/checkpoint/$ruleId
// Store 2 things:
// "checkpoints":A queue for completed checkpoint id
// "$checkpointId":A map with key of checkpoint id and value of snapshot(gob serialized)
// Assume each operator only has one instance
func getKVStore(ruleId string) (*KVStore, error) { _ = "STUB: not implemented"; return nil, nil }

// read data from badger db

func (s *KVStore) restore() error { _ = "STUB: not implemented"; return nil }

func (s *KVStore) SaveState(checkpointId int64, opId string, state map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *KVStore) SaveCheckpoint(checkpointId int64) error { _ = "STUB: not implemented"; return nil }

// TODO is the order promised?

// GetOpState Only run in the initialization
func (s *KVStore) GetOpState(opId string) (*sync.Map, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *KVStore) Clean() error { _ = "STUB: not implemented"; return nil }
