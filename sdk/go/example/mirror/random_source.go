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

package main

import (
	"github.com/lf-edge/ekuiper/sdk/go/api"
)

const dedupStateKey = "input"

type randomSourceConfig struct {
	Interval int                    `json:"interval"`
	Seed     int                    `json:"seed"`
	Pattern  map[string]interface{} `json:"pattern"`
	// how long will the source trace for deduplication. If 0, deduplicate is disabled; if negative, deduplicate will be the whole life time
	Deduplicate int    `json:"deduplicate"`
	Format      string `json:"format"`
}

// Emit data randomly with only a string field
type randomSource struct {
	conf *randomSourceConfig
	list [][]byte
}

func (s *randomSource) Configure(_ string, props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *randomSource) Open(ctx api.StreamContext, consumer chan<- api.SourceTuple, _ chan<- error) {
	_ = "STUB: not implemented"
	return
}

// dedup not supported yet
//list, err := ctx.GetState(dedupStateKey)
//if err != nil {
//	errCh <- err
//	return
//}

func randomize(p map[string]interface{}, seed int) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// TODO other data types

func (s *randomSource) isDup(ctx api.StreamContext, next map[string]interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// State not supported yet
// ctx.PutState(dedupStateKey, s.list)

func (s *randomSource) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }
