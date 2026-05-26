// Copyright 2024 EMQ Technologies Co., Ltd.
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

package sse

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/connection"
)

type SseConfig struct {
	Endpoint string `json:"endpoint"`
}

type SSESink struct {
	cw    *connection.ConnWrapper
	cfg   *SseConfig
	props map[string]any
	topic string
}

func (s *SSESink) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SSESink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *SSESink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"

	// Connection pool will handle status change
	return nil
}

func (s *SSESink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SSESink) collect(ctx api.StreamContext, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var _ api.BytesCollector = &SSESink{}

func buildSseEpID(endpoint string) string { _ = "STUB: not implemented"; return "" }
