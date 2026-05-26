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

type SqlServerQueryGenerator struct {
	*InternalSqlQueryCfg
}

func (q *SqlServerQueryGenerator) quoteIdentifier(identifier string) string {
	_ = "STUB: not implemented"
	return ""
}

func (q *SqlServerQueryGenerator) getSelect() string { _ = "STUB: not implemented"; return "" }

func (q *SqlServerQueryGenerator) getCondition() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (q *SqlServerQueryGenerator) getOrderby() string { _ = "STUB: not implemented"; return "" }

func NewSqlServerQuery(cfg *InternalSqlQueryCfg) SqlQueryGenerator {
	_ = "STUB: not implemented"
	return *new(SqlQueryGenerator)
}

func (q *SqlServerQueryGenerator) SqlQueryStatement() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (q *SqlServerQueryGenerator) UpdateMaxIndexValue(row map[string]interface{}) {
	_ = "STUB: not implemented"
	// since internal sql have asc clause, so the last element is largest
	return
}
