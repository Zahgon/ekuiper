// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

package client

import (
	"github.com/edgexfoundry/go-mod-messaging/v4/messaging"
	"github.com/edgexfoundry/go-mod-messaging/v4/pkg/types"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type Client struct {
	mbconf types.MessageBusConfig
	client messaging.MessageClient
	id     string
}

var optKeys = map[string]string{
	"clientid": "ClientId", "username": "Username", "password": "Password", "qos": "Qos", "keepalive": "KeepAlive", "retained": "Retained", "connectionpayload": "ConnectionPayload", "certfile": "CertFile", "keyfile": "KeyFile", "certpemblock": "CertPEMBlock", "keypemblock": "KeyPEMBlock", "skipcertverify": "SkipCertVerify",
}

func GetConnection(_ api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func (es *Client) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Client) GetId(ctx api.StreamContext) string { _ = "STUB: not implemented"; return "" }

func (es *Client) Dial(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (es *Client) Ping(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (es *Client) DetachSub(ctx api.StreamContext, props map[string]any) {
	_ = "STUB: not implemented"
	return
}

func (es *Client) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

type EdgexConf struct {
	Protocol string            `json:"protocol"`
	Server   string            `json:"server"`
	Host     string            `json:"host"`
	Port     int               `json:"port"`
	Type     string            `json:"type"`
	Optional map[string]string `json:"optional"`
}

// Modify the copied conf to print no password.
func printConf(mbconf types.MessageBusConfig) { _ = "STUB: not implemented"; return }

func (es *Client) CfgValidate(props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Client) Publish(env types.MessageEnvelope, topic string) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Client) Subscribe(msg chan types.MessageEnvelope, topic string, err chan error) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Client) Disconnect() error { _ = "STUB: not implemented"; return nil }
