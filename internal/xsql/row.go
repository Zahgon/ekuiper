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

package xsql

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// The original message map may be big. Make sure it is immutable so that never make a copy of it.
// The tuple clone should be cheap.

/*
 * Interfaces definition
 */

type Wildcarder interface {
	// All Value returns the value and existence flag for a given key.
	All(table string) (map[string]any, bool)
}

type Event interface {
	GetTimestamp() time.Time
	IsWatermark() bool
}

type EventRow interface {
	Row
	Event
}

type ReadonlyRow interface {
	HasTracerCtx
	Valuer
	AliasValuer
	Wildcarder
}

// RawRow is the basic data type for logical row. It could be a row or a collection row.
type RawRow interface {
	ReadonlyRow
	// Del Only for some ops like functionOp * and Alias
	Del(col string)
	// Set Only for some ops like functionOp *
	Set(col string, value interface{})
	// ToMap converts the row to a map to export to other systems *
	ToMap() map[string]interface{}
	// Pick the columns and discard others. It replaces the underlying message with a new value. There are 3 types to pick: column, alias and anonymous expressions.
	// cols is a list [columnname, tablename]
	Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool)
}

type Row interface {
	RawRow
	Clone() Row
}

type HasTracerCtx interface {
	GetTracerCtx() api.StreamContext
	SetTracerCtx(ctx api.StreamContext)
}

type MetaData interface {
	MetaData() Metadata
}

// EmittedData is data that is produced by a specific source
type EmittedData interface {
	// GetEmitter returns the emitter of the row
	GetEmitter() string
}

// CollectionRow is the aggregation row of a non-grouped collection. Thinks of it as a single group.
// The row data is immutable
type CollectionRow interface {
	RawRow
	AggregateData
	// Clone when broadcast to make sure each row are dealt single threaded
	// Clone() CollectionRow
}

type ControlTuple interface {
	ControlType() string
}

// AffiliateRow part of other row types do help calculation of newly added cols
type AffiliateRow struct {
	lock     syncx.RWMutex
	CalCols  map[string]interface{} // mutable and must be cloned when broadcast
	AliasMap map[string]interface{}
}

func (d *AffiliateRow) AppendAlias(key string, value interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *AffiliateRow) AliasValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (d *AffiliateRow) aliasValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (d *AffiliateRow) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (d *AffiliateRow) Set(col string, value interface{}) { _ = "STUB: not implemented"; return }

func (d *AffiliateRow) Del(col string) { _ = "STUB: not implemented"; return }

func (d *AffiliateRow) Clone() AffiliateRow { _ = "STUB: not implemented"; return *new(AffiliateRow) }

//nolint:govet

func (d *AffiliateRow) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (d *AffiliateRow) MergeMap(cachedMap map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

// Do not write out the internal fields

func (d *AffiliateRow) Pick(cols [][]string) [][]string { _ = "STUB: not implemented"; return nil }

/*
 *  Message definition
 */

// Message is a valuer that substitutes values for the mapped interface. It is the basic type for data events.
type Message map[string]interface{}

func (m Message) Get(key string) (value any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (m Message) Range(f func(key string, value any) bool) { _ = "STUB: not implemented"; return }

func (m Message) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

var _ Valuer = Message{}

type Metadata Message

// Alias will not need to convert cases
type Alias struct {
	AliasMap map[string]interface{}
}

/*
 * All row types definitions, watermark, barrier
 */

type RawTuple struct {
	Ctx       api.StreamContext
	Emitter   string
	Timestamp time.Time
	Rawdata   []byte
	Metadata  Metadata // immutable
	Props     map[string]string
}

func (r *RawTuple) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (r *RawTuple) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (r *RawTuple) Replace(new []byte) { _ = "STUB: not implemented"; return }

func (r *RawTuple) DynamicProps(template string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *RawTuple) AllProps() map[string]string { _ = "STUB: not implemented"; return nil }

func (r *RawTuple) Raw() []byte { _ = "STUB: not implemented"; return nil }

func (r *RawTuple) Meta(key, table string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *RawTuple) Clone() *RawTuple { _ = "STUB: not implemented"; return nil }

var (
	_ api.RawTuple        = &RawTuple{}
	_ api.HasDynamicProps = &RawTuple{}
)

// Tuple The input row, produced by the source
type Tuple struct {
	Ctx       api.StreamContext
	Emitter   string
	Message   Message // the original pointer is immutable & big; may be cloned.
	Timestamp time.Time
	Metadata  Metadata // immutable
	Props     map[string]string

	AffiliateRow
	lock      syncx.Mutex            // lock for the cachedMap, because it is possible to access by multiple sinks
	cachedMap map[string]interface{} // clone of the row and cached for performance
}

func (t *Tuple) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (t *Tuple) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

func (t *Tuple) Created() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (t *Tuple) AllMeta() map[string]any { _ = "STUB: not implemented"; return nil }

var (
	_ Row          = &Tuple{}
	_ MetaData     = &Tuple{}
	_ api.MetaInfo = &Tuple{}
)

// JoinTuple is a row produced by a join operation
type JoinTuple struct {
	Ctx    api.StreamContext
	Tuples []Row // The content is immutable, but the slice may be added or removed
	AffiliateRow
	lock      syncx.Mutex
	cachedMap map[string]interface{} // clone of the row and cached for performance of toMap
}

func (jt *JoinTuple) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (jt *JoinTuple) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

var _ Row = &JoinTuple{}

// GroupedTuples is a collection of tuples grouped by a key
type GroupedTuples struct {
	Ctx     api.StreamContext
	Content []Row
	*WindowRange
	AffiliateRow
	lock      syncx.Mutex
	cachedMap map[string]interface{} // clone of the row and cached for performance of toMap
}

func (s *GroupedTuples) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (s *GroupedTuples) SetTracerCtx(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

var _ CollectionRow = &GroupedTuples{}

/*
 *   Implementations
 */

func ToMessage(input interface{}) (Message, bool) {
	_ = "STUB: not implemented"
	return *new(Message), false
}

func (m Message) Value(key, _ string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Only when with 'SELECT * FROM ...'  and 'schemaless', the key in map is not convert to lower case.
// So all keys in map should be converted to lowercase and then compare them.

func (m Message) getIgnoreCase(key interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m Message) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// MetaData implementation

func (m Metadata) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m Metadata) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Tuple implementation

func (t *Tuple) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *Tuple) All(string) (map[string]any, bool) { _ = "STUB: not implemented"; return nil, false }

func (t *Tuple) Clone() Row { _ = "STUB: not implemented"; return *new(Row) }

// ToMap should only use in sink.
func (t *Tuple) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// clone the message

func (t *Tuple) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *Tuple) MetaData() Metadata { _ = "STUB: not implemented"; return *new(Metadata) }

func (t *Tuple) GetEmitter() string { _ = "STUB: not implemented"; return "" }

func (t *Tuple) DynamicProps(template string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (t *Tuple) AllProps() map[string]string { _ = "STUB: not implemented"; return nil }

func (t *Tuple) AggregateEval(expr ast.Expr, v CallValuer) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tuple) GetTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (t *Tuple) IsWatermark() bool { _ = "STUB: not implemented"; return false }

func (t *Tuple) FuncValue(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *Tuple) Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool) {
	_ = "STUB: not implemented"
	// invalidate cache, will calculate again
	return
}

// JoinTuple implementation

func (jt *JoinTuple) AddTuple(tuple Row) { _ = "STUB: not implemented"; return }

func (jt *JoinTuple) AddTuples(tuples []Row) { _ = "STUB: not implemented"; return }

func (jt *JoinTuple) doGetValue(key, table string, isVal bool) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TODO support key without modifier?

// TODO should use hash here

func getTupleValue(tuple Row, key string, isVal bool) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (jt *JoinTuple) GetEmitter() string { _ = "STUB: not implemented"; return "" }

func (jt *JoinTuple) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (jt *JoinTuple) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (jt *JoinTuple) All(stream string) (map[string]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (jt *JoinTuple) Clone() Row { _ = "STUB: not implemented"; return *new(Row) }

func (jt *JoinTuple) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// clone the message

func (jt *JoinTuple) Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool) {
	_ = "STUB: not implemented"
	return
}

func (jt *JoinTuple) AggregateEval(expr ast.Expr, v CallValuer) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// GroupedTuple implementation

func (s *GroupedTuples) AggregateEval(expr ast.Expr, v CallValuer) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *GroupedTuples) Value(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *GroupedTuples) Meta(key, table string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *GroupedTuples) All(_ string) (map[string]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *GroupedTuples) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (s *GroupedTuples) Clone() Row { _ = "STUB: not implemented"; return *new(Row) }

func (s *GroupedTuples) Pick(allWildcard bool, cols [][]string, wildcardEmitters map[string]bool, except []string, sendNil bool) {
	_ = "STUB: not implemented"
	return
}
