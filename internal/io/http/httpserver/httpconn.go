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

package httpserver

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type HttpPushConnection struct {
	topic    string
	cfg      *connectionCfg
	endpoint string
	method   string
	id       string
}

func (h *HttpPushConnection) GetId(ctx api.StreamContext) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *HttpPushConnection) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpPushConnection) Dial(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

type connectionCfg struct {
	Datasource string `json:"datasource"`
	Method     string `json:"method"`
}

func CreateConnection(_ api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func (h *HttpPushConnection) Ping(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpPushConnection) DetachSub(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (h *HttpPushConnection) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpPushConnection) GetTopic() string { _ = "STUB: not implemented"; return "" }
