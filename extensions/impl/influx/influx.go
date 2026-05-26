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

package influx

import (
	"crypto/tls"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/extensions/impl/tspoint"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

// c is the configuration for influx2 sink
type c struct {
	// connection
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	// http connection
	// tls conf in cert.go
	// write options
	Database    string `json:"database"`
	Measurement string `json:"measurement"`
	tspoint.WriteOptions
}

type influxSink struct {
	conf    c
	tlsconf *tls.Config
	// temp variables
	bp  client.BatchPoints
	cli client.Client
}

func (m *influxSink) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink) Ping(ctx api.StreamContext, props map[string]any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Test connection. Put it here to avoid server connection when running test in Configure

func (m *influxSink) Collect(ctx api.StreamContext, item api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink) CollectList(ctx api.StreamContext, items api.MessageTupleList) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink) collect(ctx api.StreamContext, data any) error {
	_ = "STUB: not implemented"
	return nil
}

// Write the batch

func (m *influxSink) transformPoints(ctx api.StreamContext, data any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *influxSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

func (m *influxSink) Info() model.SinkInfo { _ = "STUB: not implemented"; return *new(model.SinkInfo) }

var (
	_ api.TupleCollector = &influxSink{}
	_ util.PingableConn  = &influxSink{}
	_ model.SinkInfoNode = &influxSink{}
)
