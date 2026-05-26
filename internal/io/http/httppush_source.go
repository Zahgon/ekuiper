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

package http

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type HttpPushSource struct {
	topic    string
	sourceID string
	ch       <-chan any
	conf     *PushConf
	props    map[string]any
}

type PushConf struct {
	Method       string `json:"method"`
	BufferLength int    `json:"bufferLength"`
	DataSource   string `json:"datasource"`
}

func (h *HttpPushSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpPushSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// TODO if supports to be resource, this should change to the unique conn id

func (h *HttpPushSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpPushSource) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

var _ api.BytesSource = &HttpPushSource{}
