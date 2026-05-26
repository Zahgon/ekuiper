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

package edgex

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/edgex/client"
)

type Source struct {
	cli *client.Client

	config      map[string]any
	topic       string
	messageType messageType
	buflen      int
	conId       string
}

type SourceConf struct {
	Topic       string      `json:"topic"`
	MessageType messageType `json:"messageType"`
	BufferLen   int         `json:"bufferLength"`
}

type SubConf struct {
	Topic string `json:"topic"`
}

type messageType string

const (
	MessageTypeEvent   messageType = "event"
	MessageTypeRequest messageType = "request"
)

func (es *Source) Provision(_ api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Source) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Source) SubId(props map[string]any) string { _ = "STUB: not implemented"; return "" }

func (es *Source) Subscribe(ctx api.StreamContext, ingest api.TupleIngest, ingestErr api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

// the source is closed

// r_meta["created"] = r.Created
// r_meta["modified"] = r.Modified

// r_meta["pushed"] = r.Pushed

// meta["pushed"] = eve.Pushed

// meta["created"] = eve.Created
// meta["modified"] = eve.Modified

func (es *Source) getValue(r dtos.BaseReading, logger api.Logger) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func convertFloatArray(v string, bitSize int) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (es *Source) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }
