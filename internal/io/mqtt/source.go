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

package mqtt

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
)

// SourceConnector is the connector for mqtt source
// When sharing the same connection, each topic will have one single sourceConnector as the shared source node
type SourceConnector struct {
	tpc   string
	cfg   *Conf
	props map[string]any

	cli        *Connection
	conId      string
	eof        api.EOFIngest
	eofPayload []byte
}

type Conf struct {
	Topic      string `json:"datasource"`
	Qos        int    `json:"qos"`
	SelId      string `json:"connectionSelector"`
	EofMessage string `json:"eofMessage"`
}

func (ms *SourceConnector) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *SourceConnector) Ping(ctx api.StreamContext, props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *SourceConnector) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for connection

// Subscribe is a one time only operation for source. It connects to the mqtt broker and subscribe to the topic
// Run open before subscribe
func (ms *SourceConnector) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, _ api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *SourceConnector) onMessage(ctx api.StreamContext, msg any, ingest api.BytesIngest) {
	_ = "STUB: not implemented"
	return
}

// extract trace id

func (ms *SourceConnector) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *SourceConnector) SetEofIngest(eof api.EOFIngest) { _ = "STUB: not implemented"; return }

func GetSource() api.Source {
	_ = "STUB: not implemented"
	return *

	// SubId the mqtt connection can only sub to a topic once
	new(api.Source)
}

func (ms *SourceConnector) SubId(props map[string]any) string { _ = "STUB: not implemented"; return "" }

var (
	_ api.BytesSource   = &SourceConnector{}
	_ api.Bounded       = &SourceConnector{}
	_ util.PingableConn = &SourceConnector{}
)
