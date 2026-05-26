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
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

const (
	WebsocketTopicPrefix = "$$websocket/"
)

func recvTopic(endpoint string, isServer bool) string { _ = "STUB: not implemented"; return "" }

func sendTopic(endpoint string, isServer bool) string { _ = "STUB: not implemented"; return "" }

type websocketEndpointContext struct {
	wg    *sync.WaitGroup
	conns map[*websocket.Conn]context.CancelFunc
}

func RegisterWebSocketEndpoint(ctx api.StreamContext, endpoint string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func UnRegisterWebSocketEndpoint(endpoint string) { _ = "STUB: not implemented"; return }

// wait all process exit

func (m *GlobalServerManager) handleProcess(ctx api.StreamContext, endpoint string, instanceID int, c *websocket.Conn, cancel context.CancelFunc, parWg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func sendProcess(ctx api.StreamContext, topic, sourceID string, c *websocket.Conn, cancel context.CancelFunc, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func recvProcess(ctx api.StreamContext, topic string, c *websocket.Conn, cancel context.CancelFunc, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func (m *GlobalServerManager) RegisterWebSocketEndpoint(ctx api.StreamContext, endpoint string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (m *GlobalServerManager) UnRegisterWebSocketEndpoint(endpoint string) *websocketEndpointContext {
	_ = "STUB: not implemented"
	return nil
}

func (m *GlobalServerManager) CloseEndpointConnection(endpoint string, c *websocket.Conn) {
	_ = "STUB: not implemented"
	return
}

func (m *GlobalServerManager) AddEndpointConnection(endpoint string, c *websocket.Conn, cancel context.CancelFunc) *sync.WaitGroup {
	_ = "STUB: not implemented"
	return nil
}

func (m *GlobalServerManager) FetchInstanceID() int { _ = "STUB: not implemented"; return 0 }

// getEndpointConnections only for unit test
func (m *GlobalServerManager) getEndpointConnections(endpoint string) *websocketEndpointContext {
	_ = "STUB: not implemented"
	return nil
}
