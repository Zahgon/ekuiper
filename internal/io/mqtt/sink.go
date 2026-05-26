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

package mqtt

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/pkg/connection"
)

// AdConf is the advanced configuration for the mqtt sink
type AdConf struct {
	Tpc      string            `json:"topic"`
	Qos      byte              `json:"qos"`
	Retained bool              `json:"retained"`
	SelId    string            `json:"connectionSelector"`
	Props    map[string]string `json:"properties"`
	PVersion string            `json:"protocolVersion"`
}

type Sink struct {
	id     string
	cw     *connection.ConnWrapper
	adconf *AdConf
	config map[string]interface{}
	cli    *Connection
}

func (ms *Sink) Provision(ctx api.StreamContext, ps map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *Sink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMQTTSinkTopic(topic string) error { _ = "STUB: not implemented"; return nil }

func (ms *Sink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

// If tpc supports dynamic props(template), planner will guarantee the result has the parsed dynamic props

func (ms *Sink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (ms *Sink) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var (
	_ api.BytesCollector = &Sink{}
	_ util.PingableConn  = &Sink{}
)
