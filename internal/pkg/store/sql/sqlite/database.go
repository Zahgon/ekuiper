// Copyright 2022-2023 EMQ Technologies Co., Ltd.
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

package sqlite

import (
	"database/sql"

	// introduce sqlite
	_ "modernc.org/sqlite"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/store/definition"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type Database struct {
	db   *sql.DB
	Path string
	mu   syncx.Mutex
}

func NewSqliteDatabase(c definition.Config, name string) (definition.Database, error) {
	_ = "STUB: not implemented"
	return *new(definition.Database), nil
}

func (d *Database) Connect() error { _ = "STUB: not implemented"; return nil }

func connectionString(dpath string) string { _ = "STUB: not implemented"; return "" }

func (d *Database) Disconnect() error { _ = "STUB: not implemented"; return nil }

func (d *Database) Apply(f func(db *sql.DB) error) error { _ = "STUB: not implemented"; return nil }
