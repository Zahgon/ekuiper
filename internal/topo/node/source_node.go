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

package node

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
)

// SourceNode is a node that connects to an external source
// The SourceNode is an all-in-one source node that support connect and decode and more.
// The SourceConnectorNode is a node that only connects to external source and does not decode.
type SourceNode struct {
	*defaultNode

	s         api.Source
	interval  time.Duration
	notifySub bool
}

type sourceConf struct {
	Interval cast.DurationConf `json:"interval"`
}

// NewSourceNode creates a SourceConnectorNode
func NewSourceNode(ctx api.StreamContext, name string, ss api.Source, props map[string]any, rOpt *def.RuleOption) (*SourceNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open will be invoked by topo. It starts reading data.
func (m *SourceNode) Open(ctx api.StreamContext, ctrlCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (m *SourceNode) ingestBytes(ctx api.StreamContext, data []byte, meta map[string]any, ts time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *SourceNode) traceStart(ctx api.StreamContext, meta map[string]any, tuple xsql.HasTracerCtx) {
	_ = "STUB: not implemented"
	return
}

// If read from parent trace

func (m *SourceNode) ingestAnyTuple(ctx api.StreamContext, data any, meta map[string]any, ts time.Time) {
	_ = "STUB: not implemented"
	return
}

// Maps are expected from user extension

// expected from file which send out any tuple type

// Source tuples are expected from memory

// should never happen

func (m *SourceNode) connectionStatusChange(status string, message string) {
	_ = "STUB: not implemented"
	// TODO only send out error when status change from connected?
	return
}

func (m *SourceNode) ingestMap(t map[string]any, meta map[string]any, ts time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *SourceNode) ingestTuple(t *xsql.Tuple, ts time.Time) { _ = "STUB: not implemented"; return }

// If receiving tuple, its source is still in the system. So continue tracing

func (m *SourceNode) ingestError(ctx api.StreamContext, err error) {
	_ = "STUB: not implemented"
	return
}

func (m *SourceNode) ingestEof(ctx api.StreamContext, msg string) {
	_ = "STUB: not implemented"
	return
}

// GetSource only used for test
func (m *SourceNode) GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

const (
	OffsetKey = "$$offset"
)

func (m *SourceNode) Rewind(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (m *SourceNode) updateState(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Run Subscribe could be a long-running function
func (m *SourceNode) Run(ctx api.StreamContext, ctrlCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// Blocking and wait for connection. The connect will call the dial and retry if fails

func (m *SourceNode) runPull(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (m *SourceNode) doPull(ctx api.StreamContext, tc time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
