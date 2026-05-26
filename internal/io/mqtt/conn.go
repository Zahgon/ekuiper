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

package mqtt

import (
	"sync"
	"sync/atomic"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/mqtt/client"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type Connection struct {
	mu syncx.Mutex
	client.Client
	id        string
	server    string
	connected atomic.Bool
	status    atomic.Value
	scHandler api.StatusChangeHandler
	// key is the topic. Each topic will have only one connector map[string]*client.SubscriptionInfo
	subscriptions sync.Map
}

func CreateConnection(_ api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func ValidateConfig(props map[string]any) error { _ = "STUB: not implemented"; return nil }

func (conn *Connection) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *Connection) GetId(_ api.StreamContext) string { _ = "STUB: not implemented"; return "" }

func (conn *Connection) Dial(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// store connected status immediately to avoid publish error due to onConnect is called slower

func (conn *Connection) Status(_ api.StreamContext) modules.ConnectionStatus {
	_ = "STUB: not implemented"
	return *new(modules.ConnectionStatus)
}

func (conn *Connection) SetStatusChangeHandler(ctx api.StreamContext, sch api.StatusChangeHandler) {
	_ = "STUB: not implemented"
	return
}

func (conn *Connection) onConnect(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// should never happen. If happens because of connection, it will retry later

func (conn *Connection) onConnectLost(ctx api.StreamContext, err error) {
	_ = "STUB: not implemented"
	return
}

func (conn *Connection) onReconnecting(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (conn *Connection) DetachSub(ctx api.StreamContext, props map[string]any) {
	_ = "STUB: not implemented"
	return
}

func (conn *Connection) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (conn *Connection) Ping(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// MQTT features

func (conn *Connection) Publish(ctx api.StreamContext, topic string, qos byte, retained bool, payload []byte, properties map[string]string) error {
	_ = "STUB: not implemented"
	// Need to return error immediately so that we can enable cache immediately
	return nil
}

func (conn *Connection) Subscribe(ctx api.StreamContext, topic string, qos byte, callback client.MessageHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (conn *Connection) ParseMsg(ctx api.StreamContext, msg any) ([]byte, map[string]any, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

const (
	dataSourceProp = "datasource"
)

func getTopicFromProps(props map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var _ modules.StatefulDialer = &Connection{}
