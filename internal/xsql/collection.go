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

package xsql

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

/*
 *   Collection interfaces
 */

// AggregateData Could be a tuple or collection
type AggregateData interface {
	AggregateEval(expr ast.Expr, v CallValuer) []interface{}
}

type SortingData interface {
	HasTracerCtx
	Len() int
	Swap(i, j int)
	Index(i int) Row
}

// Collection A collection of rows as a table. It is used for window, join, group by, etc.
type Collection interface {
	HasTracerCtx
	api.MessageTupleList
	SortingData
	// GroupRange through each group. For non-grouped collection, the whole data is a single group
	GroupRange(func(i int, aggRow CollectionRow) (bool, error)) error
	// Range through each row. For grouped collection, each row is an aggregation of groups
	Range(func(i int, r ReadonlyRow) (bool, error)) error
	// RangeSet range through each row by cloning the row
	RangeSet(func(i int, r Row) (bool, error)) error
	Filter(indexes []int) Collection
	GetWindowRange() *WindowRange
	// ToMaps returns the data as a map
	ToMaps() []map[string]interface{}
	// SetIsAgg Set by project, indicate if the collection is used in an aggregate context which will affect ToMaps output
	SetIsAgg(isAgg bool)
	// ToAggMaps returns the aggregated data as a map
	ToAggMaps() []map[string]interface{}
	// ToRowMaps returns all the data in the collection
	ToRowMaps() []map[string]interface{}
	// GetBySrc returns the rows by the given emitter
	GetBySrc(emitter string) []Row
	// Clone the collection
	Clone() Collection
}

/*
 *   Collection types definitions
 */

type WindowTuples struct {
	Ctx     api.StreamContext
	Content []Row // immutable
	*WindowRange
	contentBySrc map[string][]Row // volatile, temporary cache]

	AffiliateRow
	cachedMap map[string]interface{}
	isAgg     bool
}

var (
	_ Collection    = &WindowTuples{}
	_ CollectionRow = &WindowTuples{}
)

type JoinTuples struct {
	Ctx     api.StreamContext
	Content []*JoinTuple
	*WindowRange

	AffiliateRow
	cachedMap map[string]interface{}
	isAgg     bool
}

func (s *JoinTuples) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (s *JoinTuples) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

var (
	_ Collection    = &JoinTuples{}
	_ CollectionRow = &JoinTuples{}
)

type GroupedTuplesSet struct {
	Ctx    api.StreamContext
	Groups []*GroupedTuples
	*WindowRange
}

func (s *GroupedTuplesSet) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (s *GroupedTuplesSet) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

var _ Collection = &GroupedTuplesSet{}

/*
 *   Collection implementations
 */

func (w *WindowTuples) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (w *WindowTuples) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (w *WindowTuples) Index(index int) Row { _ = "STUB: not implemented"; return *new(Row) }

func (w *WindowTuples) Len() int { _ = "STUB: not implemented"; return 0 }

func (w *WindowTuples) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (w *WindowTuples) GetBySrc(emitter string) []Row { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) GetWindowRange() *WindowRange { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) Range(f func(i int, r ReadonlyRow) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WindowTuples) RangeOfTuples(f func(index int, tuple api.MessageTuple) bool) {
	_ = "STUB: not implemented"
	return
}

func (w *WindowTuples) RangeSet(f func(i int, r Row) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WindowTuples) GroupRange(f func(i int, aggRow CollectionRow) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WindowTuples) AddTuple(tuple Row) *WindowTuples { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) AggregateEval(expr ast.Expr, v CallValuer) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Filter the tuples by the given predicate
func (w *WindowTuples) Filter(indexes []int) Collection {
	_ = "STUB: not implemented"
	return *new(Collection)
}

func (w *WindowTuples) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (w *WindowTuples) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (w *WindowTuples) All(_ string) (map[string]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (w *WindowTuples) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) Clone() Collection { _ = "STUB: not implemented"; return *new(Collection) }

func (w *WindowTuples) ToAggMaps() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) ToRowMaps() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) ToMaps() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (w *WindowTuples) Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool) {
	_ = "STUB: not implemented"
	return
}

func (w *WindowTuples) SetIsAgg(_ bool) { _ = "STUB: not implemented"; return }

func (s *JoinTuples) Len() int      { _ = "STUB: not implemented"; return 0 }
func (s *JoinTuples) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s *JoinTuples) Index(i int) Row { _ = "STUB: not implemented"; return *new(Row) }

func (s *JoinTuples) AggregateEval(expr ast.Expr, v CallValuer) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *JoinTuples) GetWindowRange() *WindowRange { _ = "STUB: not implemented"; return nil }

func (s *JoinTuples) Range(f func(i int, r ReadonlyRow) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JoinTuples) RangeOfTuples(f func(index int, tuple api.MessageTuple) bool) {
	_ = "STUB: not implemented"
	return
}

func (s *JoinTuples) RangeSet(f func(i int, r Row) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JoinTuples) GroupRange(f func(i int, aggRow CollectionRow) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// Filter the tuples by the given predicate
func (s *JoinTuples) Filter(indexes []int) Collection {
	_ = "STUB: not implemented"
	return *new(Collection)
}

func (s *JoinTuples) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *JoinTuples) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *JoinTuples) All(_ string) (map[string]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *JoinTuples) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *JoinTuples) Clone() Collection { _ = "STUB: not implemented"; return *new(Collection) }

func (s *JoinTuples) ToAggMaps() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *JoinTuples) ToRowMaps() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *JoinTuples) ToMaps() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *JoinTuples) Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool) {
	_ = "STUB: not implemented"
	return
}

func (s *JoinTuples) SetIsAgg(_ bool) {
	_ = "STUB: not implemented"

	// GetBySrc to be implemented to support join after join
	return
}

func (s *JoinTuples) GetBySrc(_ string) []Row { _ = "STUB: not implemented"; return nil }

func (s *GroupedTuplesSet) Len() int        { _ = "STUB: not implemented"; return 0 }
func (s *GroupedTuplesSet) Swap(i, j int)   { _ = "STUB: not implemented"; return }
func (s *GroupedTuplesSet) Index(i int) Row { _ = "STUB: not implemented"; return *new(Row) }

func (s *GroupedTuplesSet) GetWindowRange() *WindowRange { _ = "STUB: not implemented"; return nil }

func (s *GroupedTuplesSet) Range(f func(i int, r ReadonlyRow) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *GroupedTuplesSet) RangeOfTuples(f func(index int, tuple api.MessageTuple) bool) {
	_ = "STUB: not implemented"
	return
}

func (s *GroupedTuplesSet) RangeSet(f func(i int, r Row) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *GroupedTuplesSet) GroupRange(f func(i int, aggRow CollectionRow) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// Filter clone and return the filtered set
func (s *GroupedTuplesSet) Filter(groups []int) Collection {
	_ = "STUB: not implemented"
	return *new(Collection)
}

func (s *GroupedTuplesSet) Clone() Collection { _ = "STUB: not implemented"; return *new(Collection) }

func (s *GroupedTuplesSet) ToMaps() []map[string]any { _ = "STUB: not implemented"; return nil }

func (s *GroupedTuplesSet) SetIsAgg(_ bool) {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (s *GroupedTuplesSet) ToAggMaps() []map[string]any { _ = "STUB: not implemented"; return nil }

func (s *GroupedTuplesSet) ToRowMaps() []map[string]any {
	_ = "STUB: not implemented"

	// GetBySrc to be implemented to support join after join
	return nil
}

func (s *GroupedTuplesSet) GetBySrc(_ string) []Row {
	_ = "STUB: not implemented"

	/*
	 *  WindowRange definitions. It should be immutable
	 */return nil
}

type WindowRangeValuer struct {
	*WindowRange
}

func (w WindowRangeValuer) Value(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (w WindowRangeValuer) Meta(_, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type WindowRange struct {
	windowStart   int64
	windowEnd     int64
	windowTrigger int64
}

func NewWindowRange(windowStart int64, windowEnd int64, windowTrigger int64) *WindowRange {
	_ = "STUB: not implemented"
	return nil
}

func (r *WindowRange) FuncValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type TransformedTupleList struct {
	Ctx     api.StreamContext
	Content []api.MessageTuple
	Maps    []map[string]any
	Props   map[string]string
}

func (l *TransformedTupleList) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (l *TransformedTupleList) SetTracerCtx(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (l *TransformedTupleList) DynamicProps(template string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (l *TransformedTupleList) AllProps() map[string]string { _ = "STUB: not implemented"; return nil }

func (l *TransformedTupleList) ToMaps() []map[string]any { _ = "STUB: not implemented"; return nil }

func (l *TransformedTupleList) Clone() *TransformedTupleList { _ = "STUB: not implemented"; return nil }

func (l *TransformedTupleList) RangeOfTuples(f func(index int, tuple api.MessageTuple) bool) {
	_ = "STUB: not implemented"
	return
}

func (l *TransformedTupleList) Len() int { _ = "STUB: not implemented"; return 0 }

var (
	_ api.MessageTupleList = &TransformedTupleList{}
	_ api.HasDynamicProps  = &TransformedTupleList{}
)
