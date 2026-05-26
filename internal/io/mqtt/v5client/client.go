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

package v5client

import (
	"crypto/tls"
	"net/url"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/mqtt/client"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type Client struct {
	cm *autopaho.ConnectionManager
	syncx.Mutex
	// subscription route info
	router paho.Router
	// record if already have subscription for a topic
	subs                map[string]struct{}
	EnableClientSession bool
}

type ConnectionConfig struct {
	Server                       string `json:"server"`
	ClientId                     string `json:"clientid"`
	Uname                        string `json:"username"`
	Password                     string `json:"password"`
	EnableClientSession          bool   `json:"enableClientSession"`
	ClientStatePath              string `json:"clientStatePath"`
	SessionExpiryIntervalSeconds int    `json:"sessionExpiryIntervalSeconds"`
	serverUrl                    *url.URL
	tls                          *tls.Config
}

func Provision(ctx api.StreamContext, props map[string]any, onConnect client.ConnectHandler, onConnectLost client.ConnectErrorHandler, _ client.ConnectHandler) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO backoff?

// Keepalive message should be sent every 20 seconds
// CleanStartOnInitialConnection defaults to false. Setting this to true will clear the session on the first connection.

// SessionExpiryInterval - Seconds that a session will survive after disconnection.
// It is important to set this because otherwise, any queued messages will be lost if the connection drops and
// the server will not queue messages while it is down. The specific setting will depend upon your needs
// (60 = 1 minute, 3600 = 1 hour, 86400 = one day, 0xFFFFFFFE = 136 years, 0xFFFFFFFF = don't expire)

// eclipse/paho.golang/paho provides base mqtt functionality, the below config will be passed in for each connection

// If you are using QOS 1/2, then it's important to specify a client id (which must be unique)

// starts process; will reconnect until context cancelled

func (c *Client) Connect(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (c *Client) Subscribe(ctx api.StreamContext, topic string, qos byte, callback client.MessageHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// register router first

func (c *Client) Publish(ctx api.StreamContext, topic string, qos byte, retained bool, payload []byte, properties map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Unsubscribe(ctx api.StreamContext, topic string) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not exit immediately when unsub error. Just remove unsub handler

func (c *Client) Disconnect(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (c *Client) ParseMsg(ctx api.StreamContext, msg any) ([]byte, map[string]any, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ValidateConfig(ctx api.StreamContext, props map[string]any) (*ConnectionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ client.Client = &Client{}
