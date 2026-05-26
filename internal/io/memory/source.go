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

package memory

import (
	"regexp"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type c struct {
	Topic        string `json:"datasource"`
	BufferLength int    `json:"bufferLength"`
}

type source struct {
	topicRegex *regexp.Regexp
	c          *c
}

func (s *source) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *source) Connect(_ api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe For memory source, it can receive a source tuple directly. So just pass it through
func (s *source) Subscribe(ctx api.StreamContext, ingest api.TupleIngest, ingestErr api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

func getRegexp(topic string) (*regexp.Regexp, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *source) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSource() api.TupleSource { _ = "STUB: not implemented"; return *new(api.TupleSource) }
