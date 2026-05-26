// Copyright 2025 EMQ Technologies Co., Ltd.
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

package xsql

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

type SliceTuple struct {
	// This is immutable, it is shared by all rules with the shared source. It is accessed by source schema
	SourceContent model.SliceVal
	// After project, this is set by combining selected source field, alias field and expression field. It is accessed by sink schema
	SinkContent model.SliceVal
	// Save the calculated fields which will not sink. Currently, these are analytic result
	TempCalContent model.SliceVal
	Timestamp      time.Time
	ctx            api.StreamContext
	// TODO remove later?
	schemaMap map[string]int
	Props     map[string]string
}

func (s *SliceTuple) GetTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *SliceTuple) IsWatermark() bool { _ = "STUB: not implemented"; return false }

func (s *SliceTuple) DynamicProps(template string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *SliceTuple) AllProps() map[string]string { _ = "STUB: not implemented"; return nil }

func (s *SliceTuple) ValueByIndex(index, sourceIndex int) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// SetByIndex set sink result
func (s *SliceTuple) SetByIndex(index int, value any) { _ = "STUB: not implemented"; return }

// SetTempByIndex set analytic result. Separate it from sink to save memory in window
func (s *SliceTuple) SetTempByIndex(index int, value any) { _ = "STUB: not implemented"; return }

func (s *SliceTuple) TempByIndex(index int) any { _ = "STUB: not implemented"; return *new(any) }

func (s *SliceTuple) Compact(len int) { _ = "STUB: not implemented"; return }

func (s *SliceTuple) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (s *SliceTuple) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (s *SliceTuple) Value(key, _ string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (s *SliceTuple) Meta(key, table string) (any, bool) {
	_ = "STUB: not implemented"
	// TODO implement me
	return *new(any), false
}

func (s *SliceTuple) AliasValue(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (s *SliceTuple) AppendAlias(key string, value any) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SliceTuple) All(_ string) (map[string]any, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// do nothing

func (s *SliceTuple) Del(_ string) {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (s *SliceTuple) Set(col string, value any) { _ = "STUB: not implemented"; return }

func (s *SliceTuple) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

func (s *SliceTuple) Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool) {
	_ = "STUB: not implemented"
	return
}

func (s *SliceTuple) Clone() Row { _ = "STUB: not implemented"; return *new(Row) }

func (s *SliceTuple) FuncValue(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

var (
	_ Row                 = &SliceTuple{}
	_ Event               = &SliceTuple{}
	_ model.IndexValuer   = &SliceTuple{}
	_ api.HasDynamicProps = &SliceTuple{}
)
