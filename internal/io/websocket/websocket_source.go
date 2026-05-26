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

package websocket

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type WebsocketSource struct {
	topic         string
	cfg           *WebsocketConfig
	props         map[string]any
	connectionTyp string
	sourceID      string
}

type WebsocketConfig struct {
	Endpoint string `json:"datasource"`
}

func (w *WebsocketSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WebsocketSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (w *WebsocketSource) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WebsocketSource) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

func solveProps(props map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var _ api.BytesSource = &WebsocketSource{}
