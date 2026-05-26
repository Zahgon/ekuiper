// Copyright 2024-2026 EMQ Technologies Co., Ltd.
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

package kafka

import (
	"crypto/tls"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/segmentio/kafka-go/sasl"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type kafkaConnectionConf struct {
	Brokers string `json:"brokers"`

	tlsConfig *tls.Config
	mechanism sasl.Mechanism
}

func newKafkaConnectionConf(ctx api.StreamContext, props map[string]any) (*kafkaConnectionConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *kafkaConnectionConf) validate() error { _ = "STUB: not implemented"; return nil }

func (c *kafkaConnectionConf) ping() error { _ = "STUB: not implemented"; return nil }

func (c *kafkaConnectionConf) pingBroker(address string) error {
	_ = "STUB: not implemented"
	return nil
}

type KafkaConnection struct {
	id   string
	conf *kafkaConnectionConf
}

func init() {
	modules.RegisterConnection("kafka", CreateConnection)
}

func CreateConnection(_ api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func (k *KafkaConnection) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaConnection) Dial(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (k *KafkaConnection) GetId(_ api.StreamContext) string { _ = "STUB: not implemented"; return "" }

func (k *KafkaConnection) Ping(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (k *KafkaConnection) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

var _ modules.Connection = &KafkaConnection{}
