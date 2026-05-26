// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package memory

import (
	"regexp"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/memory/store"
)

type lc struct {
	Topic string `json:"datasource"`
	Key   string `json:"key"`
}

// lookupsource is a lookup source that reads data from memory
// The memory lookup table reads a global memory store for data
type lookupsource struct {
	topic      string
	topicRegex *regexp.Regexp
	table      *store.Table
	key        string
}

func (s *lookupsource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *lookupsource) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *lookupsource) Lookup(ctx api.StreamContext, _ []string, keys []string, values []interface{}) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *lookupsource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetLookupSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }
