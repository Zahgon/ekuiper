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

package node

import (
	"sync"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.opentelemetry.io/otel/trace"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/checkpoint"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node/metric"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type defaultNode struct {
	name        string
	concurrency int
	sendError   bool
	metricMu    syncx.RWMutex
	statManager metric.StatManager
	ctx         api.StreamContext
	ctrlCh      chan<- error
	mu          syncx.Mutex
	qos         def.Qos
	outputMu    syncx.RWMutex
	outputs     map[string]chan any
	opsWg       *sync.WaitGroup
	// tracing state
	span                     trace.Span
	spanCtx                  api.StreamContext
	disableBufferFullDiscard bool
	isStatManagerHostBySink  bool
}

func newDefaultNode(name string, options *def.RuleOption) *defaultNode {
	_ = "STUB: not implemented"
	return nil
}

func (o *defaultNode) AddOutput(output chan any, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *defaultNode) RemoveOutput(name string) error { _ = "STUB: not implemented"; return nil }

func (o *defaultNode) GetName() string { _ = "STUB: not implemented"; return "" }

func (o *defaultNode) SetQos(qos def.Qos) { _ = "STUB: not implemented"; return }

func (o *defaultNode) GetMetrics() []any { _ = "STUB: not implemented"; return nil }

func (o *defaultNode) RemoveMetrics(ruleId string) { _ = "STUB: not implemented"; return }

func (o *defaultNode) Broadcast(val any) { _ = "STUB: not implemented"; return }

func (o *defaultNode) BroadcastCustomized(val any, broadcastFunc func(val any)) {
	_ = "STUB: not implemented"
	return
}

func (o *defaultNode) doBroadcast(val any) { _ = "STUB: not implemented"; return }

// Only copy tuples except the last one(copy previous one may change val, so copy the last) when there are many outputs to save one copy time

// Fallback to set the context when sending out so that all children have the same parent ctx
// If has set ctx in the node impl, do not override it

// wait buffer consume if buffer full

// Try to send the latest one. If full, read the oldest one and retry

// read the oldest to drop.

// record the error and stop propagating to avoid infinite loop

func (o *defaultNode) GetStreamContext() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

type defaultSinkNode struct {
	*defaultNode
	input          chan any
	barrierHandler checkpoint.BarrierHandler
	inputCount     int
}

func newDefaultSinkNode(name string, options *def.RuleOption) *defaultSinkNode {
	_ = "STUB: not implemented"
	return nil
}

func (o *defaultSinkNode) GetInput() (chan any, string) { _ = "STUB: not implemented"; return nil, "" }

func (o *defaultSinkNode) GetInputCount() int { _ = "STUB: not implemented"; return 0 }

func (o *defaultSinkNode) AddInputCount() { _ = "STUB: not implemented"; return }

func (o *defaultSinkNode) SetBarrierHandler(bh checkpoint.BarrierHandler) {
	_ = "STUB: not implemented"
	return
}

func (o *defaultNode) prepareExec(ctx api.StreamContext, errCh chan<- error, opType string) {
	_ = "STUB: not implemented"
	return
}

// If the context is already cancelled, do not join the wait group

func (o *defaultNode) finishExec() { _ = "STUB: not implemented"; return }

func (o *defaultNode) Close() { _ = "STUB: not implemented"; return }

// o.ctx might be nil if not started or already closed?
// But in prepareExec o.ctx is set.
// If Close is called before prepareExec, o.ctx is nil.
// If prepareExec sets o.ctx and o.opsWg is nil?
// The original code used o.ctx.GetLogger().
// If o.ctx is nil, this would panic.
// Assuming prepareExec sets o.ctx.

func (o *defaultSinkNode) preprocess(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// if it is a barrier, return true and ignore the further processing
// if it is blocked(align handler), return true and then write back to the channel later

func (o *defaultSinkNode) commonIngest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (o *defaultSinkNode) handleEof(ctx api.StreamContext, d xsql.EOFTuple) {
	_ = "STUB: not implemented"
	return
}

// onProcessStart do the common works(metric, trace) when receiving a message from upstream
func (o *defaultNode) onProcessStart(ctx api.StreamContext, val any) {
	_ = "STUB: not implemented"
	return
}

// Source just pass nil val so that no trace. The trace will start after extracting trace id

// onProcessEnd do the common works(metric, trace) after processing a message from upstream
func (o *defaultNode) onProcessEnd(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// onSend do the common works(metric, trace) after sending a message to downstream
func (o *defaultNode) onSend(ctx api.StreamContext, val any) { _ = "STUB: not implemented"; return }

// onError do the common works(metric, trace) after throwing an error
func (o *defaultNode) onError(ctx api.StreamContext, err error) { _ = "STUB: not implemented"; return }

// onError do the common works(metric, trace) after throwing an error
func (o *defaultNode) onErrorOpt(ctx api.StreamContext, err error, sendOut bool) {
	_ = "STUB: not implemented"
	return
}

func SourcePing(sourceType string, config map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func SinkPing(sinkType string, config map[string]any) error { _ = "STUB: not implemented"; return nil }

func LookupPing(lookupType string, config map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}
