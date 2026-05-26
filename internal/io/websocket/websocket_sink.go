// Copyright 2024-2024 EMQ Technologies Co., Ltd.
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

	"github.com/lf-edge/ekuiper/v2/pkg/connection"
)

type WebsocketSink struct {
	cw    *connection.ConnWrapper
	cfg   *WebsocketConfig
	props map[string]any
	topic string
}

func (w *WebsocketSink) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WebsocketSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (w *WebsocketSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"

	// Connection pool will handle status change
	return nil
}

func (w *WebsocketSink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WebsocketSink) collect(ctx api.StreamContext, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var _ api.BytesCollector = &WebsocketSink{}

func buildWebsocketEpID(endpoint string) string { _ = "STUB: not implemented"; return "" }
