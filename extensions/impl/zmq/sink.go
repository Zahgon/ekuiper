// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

//go:build !windows

package zmq

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
	zmq "github.com/pebbe/zmq4"
)

type zmqSink struct {
	publisher *zmq.Socket
	sc        *c
}

func (m *zmqSink) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *zmqSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *zmqSink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *zmqSink) sendToZmq(ctx api.StreamContext, v []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *zmqSink) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var _ api.BytesCollector = &zmqSink{}
