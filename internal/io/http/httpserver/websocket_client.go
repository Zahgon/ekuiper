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

package httpserver

import (
	"context"
	"crypto/tls"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type WebsocketClient struct {
	RecvTopic string
	SendTopic string

	requestHeader map[string][]string
	scheme        string
	addr          string
	path          string
	tlsConfig     *tls.Config
	conn          *websocket.Conn
	wg            *sync.WaitGroup
	cancel        context.CancelFunc
}

func NewWebsocketClient(scheme, addr, path string, tlsConfig *tls.Config, requestHeader map[string][]string) *WebsocketClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *WebsocketClient) Connect() error { _ = "STUB: not implemented"; return nil }

func (c *WebsocketClient) Run(ctx api.StreamContext) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (c *WebsocketClient) handleProcess(parCtx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (c *WebsocketClient) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func extractPathAndQuery(rawURL string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
