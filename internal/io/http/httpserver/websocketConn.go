// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

type WebsocketConnection struct {
	RecvTopic string
	SendTopic string
	id        string
	props     map[string]any
	cfg       *wscConfig
	isServer  bool
	client    *WebsocketClient
}

func (w *WebsocketConnection) GetId(ctx api.StreamContext) string {
	_ = "STUB: not implemented"
	return ""
}

func (w *WebsocketConnection) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WebsocketConnection) Dial(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

type wscConfig struct {
	Path          string              `json:"path"`
	Datasource    string              `json:"datasource"`
	Addr          string              `json:"addr"`
	Scheme        string              `json:"scheme"`
	RequestHeader map[string][]string `json:"requestHeader"`
}

func (w *WebsocketConnection) Ping(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WebsocketConnection) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateWebsocketConnection(ctx api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func getWsType(cfg *wscConfig) bool { _ = "STUB: not implemented"; return false }
