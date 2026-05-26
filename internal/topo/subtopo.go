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

package topo

import (
	"context"
	"sync/atomic"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/checkpoint"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/internal/topo/schema"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// SrcSubTopo Implements node.SourceNode
type SrcSubTopo struct {
	name        string
	isSliceMode bool

	// creation state
	source node.DataSourceNode
	// May be empty
	ops         []node.OperatorNode
	tail        node.Emitter
	topo        *def.PrintableTopo
	schemaLayer *schema.SharedLayer
	// runtime state
	// Ref state, affect the pool. Update when rule created or stopped
	syncx.RWMutex
	refRules map[string]map[int]chan<- error // map[ruleId][runId]errCh, notify the rule for errors
	// Runtime state, affect the running loop. Update when any rule opened or all rules stopped
	opened           atomic.Int32 // 0 is init, 1 is open, -1 is close
	cancel           context.CancelFunc
	enableCheckpoint bool
}

const (
	InitState  int32 = 0
	OpenState  int32 = 1
	CloseState int32 = -1
)

func (s *SrcSubTopo) Init(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// Open is different from main topo because this will run multiple times.
// Each new ref rule will run open subtopo
func (s *SrcSubTopo) Open(ctx api.StreamContext, parentErrCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// Update the ref count

// Attach schemas

// If not opened yet, open it. Each ref rule start will try to run this.

// Close satisfies the node.DataSourceNode interface. All pool-lifecycle logic
// (reference counting, pool eviction, chained close) lives in CloseSubTopo so
// that subtopo_pool.go is the single owner of the pool lock and map.
func (s *SrcSubTopo) Close(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// IsSliceMode this is a constant set when creating new subtopo
func (s *SrcSubTopo) IsSliceMode() bool { _ = "STUB: not implemented"; return false }

func (s *SrcSubTopo) AddOutput(output chan interface{}, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SrcSubTopo) RemoveOutput(name string) error { _ = "STUB: not implemented"; return nil }

func (s *SrcSubTopo) updateRef(ctx api.StreamContext, parentErrCh chan<- error) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SrcSubTopo) removeRef(ctx api.StreamContext) (ruleClose bool, destroy bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Run inside ok to make sure clean up only run once

func (s *SrcSubTopo) notifyError(poe error) {
	_ = "STUB: not implemented"
	// Notify error to all ref rules
	return
}

func (s *SrcSubTopo) GetSource() node.DataSourceNode {
	_ = "STUB: not implemented"
	return *new(node.DataSourceNode)
}

func (s *SrcSubTopo) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *SrcSubTopo) SubMetrics() (keys []string, values []any) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SrcSubTopo) GetMetrics() []any { _ = "STUB: not implemented"; return nil }

func (s *SrcSubTopo) OpsCount() int {
	_ = "STUB: not implemented"

	// RefCount returns the number of rules currently referencing this sub-topology.
	// It is intended for use in tests to verify cleanup after planning errors.
	return 0
}

func (s *SrcSubTopo) RefCount() int { _ = "STUB: not implemented"; return 0 }

func (s *SrcSubTopo) StoreSchema(ruleID, dataSource string, schema map[string]*ast.JsonStreamField, isWildCard bool) {
	_ = "STUB: not implemented"
	return
}

// RemoveMetrics is called when the rule is deleted
func (s *SrcSubTopo) RemoveMetrics(ruleId string) { _ = "STUB: not implemented"; return }

func (s *SrcSubTopo) EnableCheckpoint(sources *[]checkpoint.StreamTask, ops *[]checkpoint.NonSourceTask) {
	_ = "STUB: not implemented"
	return
}

func prepareSharedContext(parCtx api.StreamContext, k string, qos def.Qos) (api.StreamContext, context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(api.StreamContext), *new(context.CancelFunc), nil
}

var (
	_ node.DataSourceNode          = &SrcSubTopo{}
	_ checkpoint.SourceSubTopoTask = &SrcSubTopo{}
)
