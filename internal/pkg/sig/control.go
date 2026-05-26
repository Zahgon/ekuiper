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

package sig

import (
	"context"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type MqttControl struct {
	// make sure this is never null
	cli      mqtt.Client
	interval time.Duration
	topic    string
	// signals
	lock syncx.RWMutex
	sigs map[string]struct{}
	// func to stop
	cancel context.CancelFunc
}

var Ctrl *MqttControl

const (
	CtrlTopic    = "ctrl/subready"
	CtrlAckTopic = "ctrl/suback"
)

// InitMQTTControl Should only called once
func InitMQTTControl() { _ = "STUB: not implemented"; return }

func NewMQTTControl(server string, cid string) *MqttControl {
	_ = "STUB: not implemented"
	// connect to MQTT
	return nil
}

// Connect to MQTT

// subscribe?

func handleToken(token mqtt.Token) error { _ = "STUB: not implemented"; return nil }

func (c *MqttControl) Add(name string) { _ = "STUB: not implemented"; return }

func (c *MqttControl) Rem(name string) { _ = "STUB: not implemented"; return }

// start run when there are topics and stop when no topics needed
func (c *MqttControl) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *MqttControl) scan() { _ = "STUB: not implemented"; return }

func (c *MqttControl) pub(message string) { _ = "STUB: not implemented"; return }
