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

type CommonQueryGenerator struct {
	*InternalSqlQueryCfg
}

func (q *CommonQueryGenerator) quoteIdentifier(identifier string) string {
	_ = "STUB: not implemented"
	return ""
}

func (q *CommonQueryGenerator) getSelect() string { _ = "STUB: not implemented"; return "" }

func (q *CommonQueryGenerator) getCondition() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (q *CommonQueryGenerator) getOrderby() string { _ = "STUB: not implemented"; return "" }

func (q *CommonQueryGenerator) getLimit() string { _ = "STUB: not implemented"; return "" }

func NewCommonSqlQuery(cfg *InternalSqlQueryCfg) SqlQueryGenerator {
	_ = "STUB: not implemented"
	return *new(SqlQueryGenerator)
}

func (q *CommonQueryGenerator) SqlQueryStatement() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (q *CommonQueryGenerator) UpdateMaxIndexValue(row map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

type OracleQueryGenerate struct {
	*CommonQueryGenerator
}

func NewOracleQueryGenerate(cfg *InternalSqlQueryCfg) SqlQueryGenerator {
	_ = "STUB: not implemented"
	return *new(SqlQueryGenerator)
}

func (q *OracleQueryGenerate) SqlQueryStatement() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (q *OracleQueryGenerate) UpdateMaxIndexValue(row map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func getCondition(cfg *InternalSqlQueryCfg, quoteIdentifier func(string) string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func buildSingleIndexCondition(w *store.IndexField, quoteIdentifier func(string) string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getOrderBy(cfg *InternalSqlQueryCfg, quoteIdentifier func(string) string) string {
	_ = "STUB: not implemented"
	return ""
}

func buildSingleOrderBy(w *store.IndexField, quoteIdentifier func(string) string) string {
	_ = "STUB: not implemented"
	return ""
}

func updateMaxIndexValue(cfg *InternalSqlQueryCfg, row map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}
