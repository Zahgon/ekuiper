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

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

const (
	SseTopicPrefix = "$$sse/"
)

type sseEndpointContext struct {
	wg    *sync.WaitGroup
	conns map[int64]context.CancelFunc
}

func recvSseTopic(endpoint string) string { _ = "STUB: not implemented"; return "" }

func sendSseTopic(endpoint string) string { _ = "STUB: not implemented"; return "" }

func RegisterSSEEndpoint(ctx api.StreamContext, endpoint string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func UnRegisterSSEEndpoint(endpoint string) { _ = "STUB: not implemented"; return }

// wait all connections to close

func (m *GlobalServerManager) RegisterSSEEndpoint(ctx api.StreamContext, endpoint string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Create a cancel context for this specific connection

// Create a subscription to the send topic
// The sourceID must be unique for each connection to ensure all clients receive the message

func (m *GlobalServerManager) AddSSEConnection(endpoint string, connID int64, cancel context.CancelFunc) (*sync.WaitGroup, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *GlobalServerManager) CloseSSEConnection(endpoint string, connID int64) {
	_ = "STUB: not implemented"
	return
}

func (m *GlobalServerManager) UnRegisterSSEEndpoint(endpoint string) *sseEndpointContext {
	_ = "STUB: not implemented"
	return nil
}

// Cancel all active connections
