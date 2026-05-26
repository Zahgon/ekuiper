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

package influx2

import (
	"crypto/tls"
	"strings"
	"time"

	client "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/extensions/impl/tspoint"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

// c is the configuration for influx2 sink
type c struct {
	// connection
	Addr         string        `json:"addr"`
	Token        string        `json:"token"`
	Org          string        `json:"org"`
	Bucket       string        `json:"bucket"`
	PrecisionStr string        `json:"precision"`
	Precision    time.Duration `json:"-"`
	// http connection
	// tls conf in cert.go
	// write options
	UseLineProtocol bool   `json:"useLineProtocol"` // 0: json, 1: line protocol
	Measurement     string `json:"measurement"`
	tspoint.WriteOptions
	BatchSize int `json:"batchSize"`
}

// influxSink2 is the sink for influx2.
// To ensure exact order, it uses blocking write api to write data to influxdb2.
type influxSink2 struct {
	conf    c
	tlsconf *tls.Config
	// save the token privately
	cli client.Client
}

func (m *influxSink2) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink2) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink2) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Test connection

func (m *influxSink2) Collect(ctx api.StreamContext, item api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink2) CollectList(ctx api.StreamContext, items api.MessageTupleList) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink2) collect(ctx api.StreamContext, data any) error {
	_ = "STUB: not implemented"
	return nil

	// Write out with blocking API to keep order. Batch is done by sink node side
}

func (m *influxSink2) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (m *influxSink2) transformPoints(ctx api.StreamContext, data any) ([]*write.Point, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *influxSink2) transformLines(ctx api.StreamContext, data any) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *influxSink2) rawPtToLine(rawPt *tspoint.RawPoint) string {
	_ = "STUB: not implemented"
	return ""
}

func writeLine(c int, builder *strings.Builder, k string, v any) int {
	_ = "STUB: not implemented"
	return 0
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

func (m *influxSink2) Info() model.SinkInfo { _ = "STUB: not implemented"; return *new(model.SinkInfo) }

var (
	_ api.TupleCollector = &influxSink2{}
	_ util.PingableConn  = &influxSink2{}
	_ model.SinkInfoNode = &influxSink2{}
)
