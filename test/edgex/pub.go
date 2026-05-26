// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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

package main

import (
	"os"

	"github.com/edgexfoundry/go-mod-messaging/v4/messaging"
	"github.com/edgexfoundry/go-mod-messaging/v4/pkg/types"
)

var msgConfig1 = types.MessageBusConfig{
	Broker: types.HostInfo{
		Host:     "127.0.0.1",
		Port:     1883,
		Protocol: "tcp",
	},
	Optional: map[string]string{
		"ClientId": "0001_client_id",
	},
	Type: messaging.MQTT,
}

func pubDefault() { _ = "STUB: not implemented"; return }

// r := rand.New(rand.NewSource(time.Now().UnixNano()))

// temp := r.Intn(100)
// humd := r.Intn(100)

func pubArrayMessage() { _ = "STUB: not implemented"; return }

func pubToMQTT(host string) { _ = "STUB: not implemented"; return }

func pubMetaSource() { _ = "STUB: not implemented"; return }

func main() {
	if len(os.Args) == 1 {
		pubDefault()
	} else if len(os.Args) == 2 {
		if v := os.Args[1]; v == "meta" {
			pubMetaSource()
		} else if v == "array" {
			pubArrayMessage()
		}
	} else if len(os.Args) == 3 {
		if v := os.Args[1]; v == "mqtt" {
			// The 2nd parameter is MQTT broker server address
			pubToMQTT(os.Args[2])
		}
	}
}
