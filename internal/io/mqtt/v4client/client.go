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

package v4client

import (
	"crypto/tls"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/mqtt/client"
)

type Client struct {
	cli                 pahoMqtt.Client
	EnableClientSession bool
}

type ConnectionConfig struct {
	Server              string `json:"server"`
	PVersion            string `json:"protocolVersion"`
	ClientId            string `json:"clientid"`
	Uname               string `json:"username"`
	Password            string `json:"password"`
	EnableClientSession bool   `json:"enableClientSession"`
	pversion            uint   // 3 or 4
	tls                 *tls.Config
}

func Provision(ctx api.StreamContext, props map[string]any, onConnect client.ConnectHandler, onConnectLost client.ConnectErrorHandler, onReconnect client.ConnectHandler) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Connect(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (c *Client) ParseMsg(ctx api.StreamContext, p any) ([]byte, map[string]any, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) Publish(_ api.StreamContext, topic string, qos byte, retained bool, payload []byte, _ map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Subscribe(ctx api.StreamContext, topic string, qos byte, callback client.MessageHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Unsubscribe(_ api.StreamContext, topic string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Disconnect(_ api.StreamContext) { _ = "STUB: not implemented"; return }

func ValidateConfig(ctx api.StreamContext, props map[string]any) (*ConnectionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default to MQTT 3.1.1 or NanoMQ cannot connect

func handleToken(token pahoMqtt.Token) error { _ = "STUB: not implemented"; return nil }

var _ client.Client = &Client{}
