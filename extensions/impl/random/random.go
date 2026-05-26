// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

package random

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

const dedupStateKey = "input"

type randomSourceConfig struct {
	Seed    int                    `json:"seed"`
	Pattern map[string]interface{} `json:"pattern"`
	// how long will the source trace for deduplication. If 0, deduplicate is disabled; if negative, deduplicate will be the whole lifetime
	Deduplicate int    `json:"deduplicate"`
	Format      string `json:"format"`
}

// Emit data randomly with only a string field
type randomSource struct {
	conf *randomSourceConfig
	list [][]byte
}

func (s *randomSource) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *randomSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *randomSource) Pull(ctx api.StreamContext, trigger time.Time, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

func randomize(p map[string]interface{}, seed int) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// TODO other data types

func (s *randomSource) isDup(ctx api.StreamContext, next map[string]interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *randomSource) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var _ api.PullTupleSource = &randomSource{}
