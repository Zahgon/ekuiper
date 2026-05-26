// Copyright 2024 EMQ Technologies Co., Ltd.
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

package model

import (
	"time"
)

type DefaultSourceTuple struct {
	message map[string]any
	meta    map[string]any
	time    time.Time
	raw     []byte
}

// NewDefaultRawTuple creates a new DefaultSourceTuple with raw data. Use this when extend source connector
func NewDefaultRawTuple(raw []byte, meta map[string]any, ts time.Time) *DefaultSourceTuple {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultRawTupleIgnoreTs(raw []byte, meta map[string]any) *DefaultSourceTuple {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultSourceTuple(message map[string]any, meta map[string]any, timestamp time.Time) *DefaultSourceTuple {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultSourceTuple) Value(key, table string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (t *DefaultSourceTuple) Range(f func(key string, value any) bool) {
	_ = "STUB: not implemented"
	return
}

func (t *DefaultSourceTuple) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

func (t *DefaultSourceTuple) Meta(key, table string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (t *DefaultSourceTuple) AllMeta() map[string]any { _ = "STUB: not implemented"; return nil }

func (t *DefaultSourceTuple) Timestamp() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (t *DefaultSourceTuple) Raw() []byte { _ = "STUB: not implemented"; return nil }
