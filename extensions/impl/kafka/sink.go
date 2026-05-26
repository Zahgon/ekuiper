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
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node/metric"
	"github.com/lf-edge/ekuiper/v2/pkg/connection"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

const (
	lblBuild     = "build"
	LblUnmarshal = "unmarshal"
	LblCollect   = "collect"
	LblReq       = "req"
	LblKafka     = "kafka"
	LblMsg       = "msg"
	LblQueueIn   = "queue-in"
	LblIngest    = "ingest"
	LblSend      = "send"
	LblBytes     = "bytes"
	LblOffset    = "offset"
)

type KafkaSink struct {
	props          map[string]any
	writer         *kafkago.Writer
	transport      *kafkago.Transport
	kc             *kafkaConf
	tlsConfig      *tls.Config
	headersMap     map[string]string
	headerTemplate string
	cw             *connection.ConnWrapper
	saslConf       *saslConf
	mechanism      sasl.Mechanism
	LastStats      kafkago.WriterStats
	msgQ           chan *kafkago.Message
	messages       []kafkago.Message
	currIndex      int
	ruleID         string
	opID           string
	statManager    metric.StatManager
	connected      bool
	sch            api.StatusChangeHandler
}

func (k *KafkaSink) setStatManager(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (k *KafkaSink) Info() model.SinkInfo { _ = "STUB: not implemented"; return *new(model.SinkInfo) }

type KafkaCollectStats struct {
	TotalBuildMsgDuration     time.Duration
	TotalUnmarshalMsgDuration time.Duration
	TotalCollectMsgDuration   time.Duration
}

type kafkaConf struct {
	kafkaWriterConf
	Brokers        string        `json:"brokers"`
	Topic          string        `json:"topic"`
	SelId          string        `json:"connectionSelector"`
	MaxAttempts    int           `json:"maxAttempts"`
	RequiredACKs   int           `json:"requiredACKs"`
	Key            string        `json:"key"`
	Headers        interface{}   `json:"headers"`
	LingerInterval time.Duration `json:"lingerInterval"`

	// write config
	Compression string `json:"compression"`
}

type kafkaWriterConf struct {
	BatchSize    int           `json:"batchSize"`
	BatchTimeout time.Duration `json:"-"`
	BatchBytes   int64         `json:"batchBytes"`
}

func (c *kafkaConf) validate() error { _ = "STUB: not implemented"; return nil }

func (k *KafkaSink) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// run batch

func (k *KafkaSink) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *kafkaConnectionConf) pingBrokerRaw(address string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSink) buildKafkaWriter(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// kafka java-client default balancer

func (k *KafkaSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (k *KafkaSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSink) runWithTickerAndBatchSize(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (k *KafkaSink) runWithBatchSize(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (k *KafkaSink) runWithTicker(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (k *KafkaSink) ingest(ctx api.StreamContext, d *kafkago.Message, checkSize bool) {
	_ = "STUB: not implemented"
	return
}

func (k *KafkaSink) send(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (k *KafkaSink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSink) collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KafkaSink) buildMsg(ctx api.StreamContext, item api.RawTuple) (kafkago.Message, error) {
	_ = "STUB: not implemented"
	return *new(kafkago.Message), nil
}

func (k *KafkaSink) setHeaders() error { _ = "STUB: not implemented"; return nil }

func (k *KafkaSink) parseHeaders(ctx api.StreamContext, item api.RawTuple) ([]kafkago.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KafkaSink) handleErrMsgs(ctx api.StreamContext, err error, count int) {
	_ = "STUB: not implemented"
	return
}

func (k *KafkaSink) handleConnectedSch(err error) { _ = "STUB: not implemented"; return }

func toCompression(c string) kafkago.Compression {
	_ = "STUB: not implemented"
	return *new(kafkago.Compression)
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var (
	_ api.BytesCollector = &KafkaSink{}
	_ util.PingableConn  = &KafkaSink{}
	_ model.SinkInfoNode = &KafkaSink{}
)

func getDefaultKafkaConf() *kafkaConf { _ = "STUB: not implemented"; return nil }

// 1MB

func (kc *kafkaConf) configure(props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
