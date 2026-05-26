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

package nexmark

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type NexmarkSourceConfig struct {
	Qps            int  `json:"qps"`
	BufferSize     int  `json:"bufferSize"`
	ExcludePerson  bool `json:"excludePerson"`
	ExcludeAuction bool `json:"excludeAuction"`
	ExcludeBid     bool `json:"excludeBid"`
}

type NexmarkSource struct {
	config    NexmarkSourceConfig
	generator *EventGenerator
}

func (n *NexmarkSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NexmarkSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (n *NexmarkSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NexmarkSource) Subscribe(ctx api.StreamContext, ingest api.TupleIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }
