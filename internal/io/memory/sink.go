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
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/memory/pubsub"
)

type config struct {
	Topic        string `json:"topic"`
	RowkindField string `json:"rowkindField"`
	KeyField     string `json:"keyField"`
}

type sink struct {
	topic        string
	keyField     string
	rowkindField string
	meta         map[string]any
}

func (s *sink) Provision(_ api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) Collect(ctx api.StreamContext, data api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) wrapUpdatable(el pubsub.MemTuple) (pubsub.MemTuple, error) {
	_ = "STUB: not implemented"
	return *new(pubsub.MemTuple), nil
}

func (s *sink) CollectList(ctx api.StreamContext, tuples api.MessageTupleList) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSink() api.TupleCollector { _ = "STUB: not implemented"; return *new(api.TupleCollector) }
