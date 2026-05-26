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

package store

import (
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type IndexField struct {
	IndexFieldName           string      `json:"indexField"`
	IndexFieldValue          interface{} `json:"indexValue"`
	IndexFieldDataType       string      `json:"indexFieldType"`
	IndexFieldDateTimeFormat string      `json:"dateTimeFormat"`
}

type IndexFieldStoreWrap struct {
	// use mutex to modify value in future
	syncx.RWMutex
	store *IndexFieldStore
}

type IndexFieldStore struct {
	IndexFieldValueList []*IndexField          `json:"indexFieldValueList"`
	IndexFieldValueMap  map[string]*IndexField `json:"indexFieldValueMap"`
}

func NewIndexFieldWrap(fields ...*IndexField) *IndexFieldStoreWrap {
	_ = "STUB: not implemented"
	return nil
}

func (wrap *IndexFieldStoreWrap) InitByStore(store *IndexFieldStore) {
	_ = "STUB: not implemented"
	return
}

func (wrap *IndexFieldStoreWrap) GetStore() *IndexFieldStore { _ = "STUB: not implemented"; return nil }

func (wrap *IndexFieldStoreWrap) Init(fields ...*IndexField) { _ = "STUB: not implemented"; return }

func (wrap *IndexFieldStoreWrap) GetFieldList() []*IndexField {
	_ = "STUB: not implemented"
	return nil
}

func (wrap *IndexFieldStoreWrap) GetFieldMap() map[string]*IndexField {
	_ = "STUB: not implemented"
	return nil
}

func (wrap *IndexFieldStoreWrap) UpdateFieldValue(name string, value interface{}) {
	_ = "STUB: not implemented"
	return
}

func (wrap *IndexFieldStoreWrap) UpdateByInput(input map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (wrap *IndexFieldStoreWrap) LoadFromList() { _ = "STUB: not implemented"; return }
