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
	"text/template"

	"github.com/lf-edge/ekuiper/v2/pkg/store"
)

type templateSqlQuery struct {
	tp *template.Template
	*TemplateSqlQueryCfg
}

func NewTemplateSqlQuery(cfg *TemplateSqlQueryCfg) (SqlQueryGenerator, error) {
	_ = "STUB: not implemented"
	return *new(SqlQueryGenerator), nil
}

func (t *templateSqlQuery) init() error {
	tp, err := template.New("sql").Parse(t.TemplateSql)
	if err != nil {
		return err
	}
	t.tp = tp
	return nil
}

func (t *templateSqlQuery) SqlQueryStatement() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *templateSqlQuery) UpdateMaxIndexValue(row map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

type TemplateSqlQueryCfg struct {
	TemplateSql string              `json:"templateSql"`
	IndexFields []*store.IndexField `json:"indexFields"`
	store       *store.IndexFieldStoreWrap
}

func (t *TemplateSqlQueryCfg) InitIndexFieldStore() { _ = "STUB: not implemented"; return }

func (t *TemplateSqlQueryCfg) SetIndexValue(v interface{}) { _ = "STUB: not implemented"; return }

func (t *TemplateSqlQueryCfg) GetIndexValue() interface{} { _ = "STUB: not implemented"; return nil }

func (t *TemplateSqlQueryCfg) GetIndexValueWrap() *store.IndexFieldStoreWrap {
	_ = "STUB: not implemented"
	return nil
}
