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
	"github.com/lf-edge/ekuiper/v2/pkg/connection"
)

type SinkConf struct {
	Topic        string      `json:"topic"`
	TopicPrefix  string      `json:"topicPrefix"`
	MessageType  messageType `json:"messageType"`
	ContentType  string      `json:"contentType"`
	DeviceName   string      `json:"deviceName"`
	ProfileName  string      `json:"profileName"`
	SourceName   string      `json:"sourceName"`
	Metadata     string      `json:"metadata"`
	DataTemplate string      `json:"dataTemplate"`
	Fields       []string    `json:"fields"`
	DataField    string      `json:"dataField"`
}

type EdgexMsgBusSink struct {
	c *SinkConf

	config map[string]any
	topic  string

	id         string
	cw         *connection.ConnWrapper
	cli        *client.Client
	sendParams map[string]any
}

func (ems *EdgexMsgBusSink) Provision(ctx api.StreamContext, ps map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// If meta data are static, the "dynamic" topic is static

// calculate dynamically

func (ems *EdgexMsgBusSink) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (ems *EdgexMsgBusSink) produceEvents(ctx api.StreamContext, item any) (*dtos.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// impossible

// Override the devicename if user specified the value

// Ignore nil values

// default media type

func getValueType(v any) (string, any, error) { _ = "STUB: not implemented"; return "", *new(any), nil }

// default to string array

func getValueByType(v any, vt string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (ems *EdgexMsgBusSink) getMeta(result []map[string]any) *meta {
	_ = "STUB: not implemented"
	return nil
}

// Try to get the meta field

func (ems *EdgexMsgBusSink) Collect(ctx api.StreamContext, data api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (ems *EdgexMsgBusSink) CollectList(ctx api.StreamContext, data api.MessageTupleList) error {
	_ = "STUB: not implemented"
	return nil
}

func (ems *EdgexMsgBusSink) doCollect(ctx api.StreamContext, item any) error {
	_ = "STUB: not implemented"
	return nil
}

// dynamic topic

func (ems *EdgexMsgBusSink) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

type eventMeta struct {
	id          *string
	deviceName  string
	profileName string
	sourceName  string
	origin      *int64
	tags        map[string]any
}

type readingMeta struct {
	id           *string
	deviceName   *string
	profileName  *string
	resourceName *string
	origin       *int64
	valueType    *string
	mediaType    *string
}

func (m *readingMeta) decorate(r *dtos.BaseReading) dtos.BaseReading {
	_ = "STUB: not implemented"
	return *new(dtos.BaseReading)
}

type meta struct {
	eventMeta
	readingMetas map[string]any
}

func newMetaFromMap(m1 map[string]any) *meta { _ = "STUB: not implemented"; return nil }

func (m *meta) readingMeta(ctx api.StreamContext, readingName string) *readingMeta {
	_ = "STUB: not implemented"
	return nil
}

func (m *meta) createEvent() *dtos.Event { _ = "STUB: not implemented"; return nil }

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }
