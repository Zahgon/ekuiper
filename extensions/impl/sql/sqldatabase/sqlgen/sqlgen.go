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

package sqlgen

import (
	"github.com/lf-edge/ekuiper/v2/pkg/store"
)

type SqlQueryGenerator interface {
	IndexValuer
	SqlQueryStatement() (string, error)
	UpdateMaxIndexValue(rows map[string]interface{})
}

type IndexValuer interface {
	SetIndexValue(interface{})
	GetIndexValue() interface{}
	GetIndexValueWrap() *store.IndexFieldStoreWrap
}

const DATETIME_TYPE = "DATETIME"

type InternalSqlQueryCfg struct {
	Table       string              `json:"table"`
	Limit       int                 `json:"limit"`
	IndexFields []*store.IndexField `json:"indexFields"`
	store       *store.IndexFieldStoreWrap
}

func (i *InternalSqlQueryCfg) InitIndexFieldStore() { _ = "STUB: not implemented"; return }

func (i *InternalSqlQueryCfg) SetIndexValue(v interface{}) { _ = "STUB: not implemented"; return }

func (i *InternalSqlQueryCfg) GetIndexValueWrap() *store.IndexFieldStoreWrap {
	_ = "STUB: not implemented"
	return nil
}

func (i *InternalSqlQueryCfg) GetIndexValue() interface{} { _ = "STUB: not implemented"; return nil }

type sqlConfig struct {
	TemplateSqlQueryCfg *TemplateSqlQueryCfg `json:"templateSqlQueryCfg"`
	InternalSqlQueryCfg *InternalSqlQueryCfg `json:"internalSqlQueryCfg"`
}

func (cfg *sqlConfig) Init(props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func formatIndexFieldsDatetime(indexFields []*store.IndexField) error {
	_ = "STUB: not implemented"
	return nil
}

func GetQueryGenerator(driver string, props map[string]interface{}) (SqlQueryGenerator, error) {
	_ = "STUB: not implemented"
	return *new(SqlQueryGenerator), nil
}
