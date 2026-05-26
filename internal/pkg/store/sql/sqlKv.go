// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

package sql

import (
	"database/sql"
)

type sqlKvStore struct {
	database Database
	table    string

	preparedGetStmt         *sql.Stmt
	preparedSetStmt         *sql.Stmt
	preparedDeleteQueryStmt *sql.Stmt
	preparedDeleteStmt      *sql.Stmt
	preparedGetByPrefixStmt *sql.Stmt
}

func createSqlKvStore(database Database, table string) (*sqlKvStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv *sqlKvStore) initPreparedStmt() error { _ = "STUB: not implemented"; return nil }

func (kv *sqlKvStore) Setnx(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv *sqlKvStore) Set(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv *sqlKvStore) Get(key string, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (kv *sqlKvStore) GetKeyedState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv *sqlKvStore) SetKeyedState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kv *sqlKvStore) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (kv *sqlKvStore) GetByPrefix(prefix string) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv *sqlKvStore) Keys() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (kv *sqlKvStore) All() (all map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kv *sqlKvStore) Clean() error { _ = "STUB: not implemented"; return nil }

func (kv *sqlKvStore) Drop() error { _ = "STUB: not implemented"; return nil }
