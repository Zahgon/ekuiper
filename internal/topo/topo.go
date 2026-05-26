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

package topo

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/checkpoint"
	kctx "github.com/lf-edge/ekuiper/v2/internal/topo/context"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

var uid atomic.Uint32

// State represents the lifecycle state
type State int

const (
	StateInitialized State = iota
	StateOpened
	StateClosed
)

// Topo is the runtime DAG for a rule
// It only runs once. If the rule restarts, another topo is created.
type Topo struct {
	streams      []string
	sources      []node.DataSourceNode
	sinks        []node.DataSinkNode
	ctx          api.StreamContext
	cancel       context.CancelFunc
	drain        chan error
	ops          []node.OperatorNode
	subSrcOpsMap map[string]struct{}
	name         string
	runId        int
	options      *def.RuleOption
	store        api.Store
	coordinator  *checkpoint.Coordinator
	topo         *def.PrintableTopo
	sinkSchema   map[string]*ast.JsonStreamField
	opsWg        *sync.WaitGroup
	spawnDone    chan struct{}
	// all other things are read only during lifecycle except state
	state atomic.Value
}

func NewWithNameAndOptions(name string, options *def.RuleOption) (*Topo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ensure context is set

func (s *Topo) Open() <-chan error { _ = "STUB: not implemented"; return nil }

// open stream sink, after log sink is ready.

// activate checkpoint

func (s *Topo) doClose() (closed bool) { _ = "STUB: not implemented"; return false }

func (s *Topo) doClean() {
	_ = "STUB: not implemented"
	// completion signal to inform topo.Open receiver
	return
}

// inform the operators to exit

// Cancel may be called multiple times so must be idempotent
func (s *Topo) Cancel() { _ = "STUB: not implemented"; return }

func (s *Topo) GracefulStop(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

// GetSourceNodes only for test
func (s *Topo) GetSourceNodes() []node.DataSourceNode { _ = "STUB: not implemented"; return nil }

func (s *Topo) GetRunId() int { _ = "STUB: not implemented"; return 0 }

func (s *Topo) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (s *Topo) SetStreams(streams []string) { _ = "STUB: not implemented"; return }

func (s *Topo) GetStreams() []string { _ = "STUB: not implemented"; return nil }

func (s *Topo) SetSinkSchema(sinkSchema map[string]*ast.JsonStreamField) {
	_ = "STUB: not implemented"
	return
}

func (s *Topo) GetSinkSchema() map[string]*ast.JsonStreamField {
	_ = "STUB: not implemented"
	return nil
}

func (s *Topo) GetContext() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (s *Topo) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *Topo) AddSrc(src node.DataSourceNode) *Topo { _ = "STUB: not implemented"; return nil }

func (s *Topo) AddSink(inputs []node.Emitter, snk node.DataSinkNode) *Topo {
	_ = "STUB: not implemented"
	return nil
}

func (s *Topo) AddSinkAlterOperator(sink *node.SinkNode, operator node.OperatorNode) *Topo {
	_ = "STUB: not implemented"
	return nil
}

func (s *Topo) AddOperator(inputs []node.Emitter, operator node.OperatorNode) *Topo {
	_ = "STUB: not implemented"
	return nil
}

// add rule id to make operator name unique

func (s *Topo) addEdge(from node.TopNode, to node.TopNode, toType string) {
	_ = "STUB: not implemented"
	return
}

// prepareContext setups internal context before
// Only run once when new topo
func (s *Topo) prepareContext() { _ = "STUB: not implemented"; return }

func (s *Topo) EnableTracer(isEnabled bool, strategy kctx.TraceStrategy) {
	_ = "STUB: not implemented"
	return
}

func (s *Topo) IsTraceEnabled() bool { _ = "STUB: not implemented"; return false }

func (s *Topo) enableCheckpoint(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// do nothing for now
// rt.EnableCheckpoint(&sources, &ops)
// should never happen

// If use shared stream sub topo, ignore subtopo for this rule's checkpoint.
// Send the barrier from the first op after subtopo

func (s *Topo) GetCoordinator() *checkpoint.Coordinator { _ = "STUB: not implemented"; return nil }

func (s *Topo) GetMetricsV2() map[string]map[string]any { _ = "STUB: not implemented"; return nil }

func (s *Topo) GetMetrics() (keys []string, values []any) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Topo) RemoveMetrics() { _ = "STUB: not implemented"; return }

func (s *Topo) GetTopo() *def.PrintableTopo { _ = "STUB: not implemented"; return nil }

func (s *Topo) ResetStreamOffset(name string, input map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Topo) waitClose() {
	_ = "STUB: not implemented"
	// wait all operators close and spawning finish
	return
}
