// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

	"github.com/lf-edge/ekuiper/v2/pkg/nng"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const (
	DefaultNeuronUrl = "ipc:///tmp/neuron-ekuiper.ipc"
	PROTOCOL         = "pair"
)

var (
	NeuronTraceHeader           = []byte{0x0A, 0xCE}
	NeuronTraceIDStartIndex     = len(NeuronTraceHeader)
	NeuronTraceIDEndIndex       = NeuronTraceIDStartIndex + 16
	NeuronTraceSpanIDStartIndex = NeuronTraceIDEndIndex
	NeuronTraceSpanIDEndIndex   = NeuronTraceSpanIDStartIndex + 8
	NeuronTraceHeaderLen        = 2 + 16 + 8
)

type source struct {
	c     *nng.SockConf
	cli   *nng.Sock
	props map[string]any
	conId string
	mu    syncx.RWMutex
}

func (s *source) Provision(_ api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *source) ConnId(props map[string]any) string { _ = "STUB: not implemented"; return "" }

func (s *source) SubId(_ map[string]any) string { _ = "STUB: not implemented"; return "" }

func (s *source) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *source) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, ingestErr api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

// no receiving deadline, will wait until the socket closed

func (s *source) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

func extractTraceMeta(ctx api.StreamContext, data []byte) ([]byte, map[string]interface{}) {
	_ = "STUB: not implemented"

	// extract rawData
	return nil, nil
}

// by setting traceId meta, source node knows how to construct a trace
