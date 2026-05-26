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
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var (
	subTopoPool = make(map[string]*SrcSubTopo)
	lock        syncx.Mutex
)

func GetOrCreateSubTopo(ctx api.StreamContext, name string, isSliceMode bool, init func(*SrcSubTopo) error) (*SrcSubTopo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// init runs under the pool lock so a new subtopo is not visible before
// it is usable. Keep init lightweight; node Provision must not perform
// time-consuming work. TODO: revisit if any source init becomes slow.

func RemoveSubTopo(name string) { _ = "STUB: not implemented"; return }

// GetSubTopoPoolSize returns the number of entries in the subtopo pool.
// It is intended for use in tests to verify cleanup after planning errors.
func GetSubTopoPoolSize() int { _ = "STUB: not implemented"; return 0 }

// CloseSubTopo decrements the reference count for ctx's rule on s and, when
// the last reference is removed, atomically cancels the subtopo and evicts it
// from the pool. Both locks (pool first, then subtopo) are held for the entire
// operation, so the destroy decision and the pool delete are one atomic step —
// no zombie window is possible. This mirrors GetOrCreateSubTopo: the pool file
// is the sole owner of all mutations to lock and subTopoPool.
func CloseSubTopo(ctx api.StreamContext, s *SrcSubTopo) { _ = "STUB: not implemented"; return }

// Always detach the rule from the schema layer, even if the sub-topology is being destroyed.
// This ensures RemoveRuleSchema is called for the global registry.

// Only update surviving rules' operators if the sub-topology isn't being destroyed.

// Capture chained source subtopo pointer before releasing locks.

// If the destroyed subtopo's source was itself a subtopo (chained connection),
// close it now — after releasing all locks to avoid any nesting issue.

func (s *SrcSubTopo) AddSrc(src node.DataSourceNode) *SrcSubTopo {
	_ = "STUB: not implemented"
	return nil
}

// AddOperator adds an internal operator to the subtopo.
func (s *SrcSubTopo) AddOperator(inputs []node.Emitter, operator node.OperatorNode) *SrcSubTopo {
	_ = "STUB: not implemented"
	return nil
}

func (s *SrcSubTopo) addEdge(from node.TopNode, to node.TopNode, toType string) {
	_ = "STUB: not implemented"
	return
}

func (s *SrcSubTopo) MergeSrc(parentTopo *def.PrintableTopo) { _ = "STUB: not implemented"; return }

func (s *SrcSubTopo) LinkTopo(parentTopo *def.PrintableTopo, parentJointName string) {
	_ = "STUB: not implemented"
	return
}

var _ node.MergeableTopo = &SrcSubTopo{}
