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

package kafka

import (
	"crypto/tls"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
)

type KafkaSource struct {
	reader    *kafkago.Reader
	offset    int64
	tlsConfig *tls.Config
	sc        *kafkaSourceConf
	saslConf  *saslConf
	mechanism sasl.Mechanism
	connected bool
	sch       api.StatusChangeHandler
}

type kafkaSourceConf struct {
	Topic       string `json:"datasource"`
	Brokers     string `json:"brokers"`
	GroupID     string `json:"groupID"`
	Partition   int    `json:"partition"`
	MaxAttempts int    `json:"maxAttempts"`
	MaxBytes    int    `json:"maxBytes"`
}

func (c *kafkaSourceConf) validate() error { _ = "STUB: not implemented"; return nil }

func (c *kafkaSourceConf) GetReaderConfig() *kafkago.ReaderConfig {
	_ = "STUB: not implemented"
	return nil
}

func getSourceConf(props map[string]interface{}) (*kafkaSourceConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KafkaSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSource) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSource) ping(address string) error { _ = "STUB: not implemented"; return nil }

func (k *KafkaSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (k *KafkaSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSource) handleConnectedSch(err error) { _ = "STUB: not implemented"; return }

func (k *KafkaSource) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSource) Rewind(offset interface{}) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

func (k *KafkaSource) ResetOffset(input map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSource) GetOffset() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

const (
	SASL_NONE  = "none"
	SASL_PLAIN = "plain"
	SASL_SCRAM = "scram"
)

type saslConf struct {
	SaslAuthType string `json:"saslAuthType"`
	SaslUserName string `json:"saslUserName"`
	SaslPassword string `json:"password"`
	OldPassword  string `json:"saslPassword,omitempty"`
}

func getSaslConf(props map[string]interface{}) (*saslConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *saslConf) resolvePassword() { _ = "STUB: not implemented"; return }

func (c *saslConf) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *saslConf) GetMechanism() (sasl.Mechanism, error) {
	_ = "STUB: not implemented"
	return *new(sasl.Mechanism), nil
}

// sasl authentication type

const (
	mockErrStart int = iota
	castConfErr
	saslConfErr
	mechanismErr
	mockErrEnd
)

func mockKakfaSourceErr(v, exp int) error { _ = "STUB: not implemented"; return nil }

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var (
	_ api.BytesSource   = &KafkaSource{}
	_ util.PingableConn = &KafkaSource{}
)
