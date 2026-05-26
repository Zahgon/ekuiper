// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package neuron

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/connection"
	"github.com/lf-edge/ekuiper/v2/pkg/nng"
)

type c struct {
	NodeName  string   `json:"nodeName"`
	GroupName string   `json:"groupName"`
	Tags      []string `json:"tags"`
	// If sent with the raw converted string or let us range over the result map
	Raw bool `json:"raw"`
}

type sink struct {
	cw    *connection.ConnWrapper
	c     *c
	cc    *nng.SockConf
	cli   *nng.Sock
	props map[string]any
}

type neuronTemplate struct {
	GroupName string      `json:"group_name"`
	NodeName  string      `json:"node_name"`
	Tags      []neuronTag `json:"tags"`
}

type neuronTag struct {
	Name  string `json:"tag_name"`
	Value any    `json:"value"`
}

func (s *sink) Provision(_ api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) Collect(ctx api.StreamContext, data api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

// CollectList sends all data at best effort
// It never return error, so it is not supported for cache and retry
func (s *sink) CollectList(ctx api.StreamContext, data api.MessageTupleList) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *sink) SendMapToNeuron(ctx api.StreamContext, tuple api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

// Send as many tags as possible in order and drop the tag if it is invalid

func doPublish(ctx api.StreamContext, cli *nng.Sock, tuple api.MessageTuple, t *neuronTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

func extractSpanContextIntoData(ctx api.StreamContext, data any, sendBytes []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }
