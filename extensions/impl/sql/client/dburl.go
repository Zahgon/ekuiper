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

package client

import (
	"database/sql"

	_ "github.com/lf-edge/ekuiper/v2/extensions/impl/sql/sqldatabase/driver"
)

func ParseDBUrl(urlstr string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Open returns *sql.DB from urlstr
// As we use modernc.org/sqlite with `sqlite` as driver name and dburl use `sqlite3` as driver name, we need to fix it before open sql.DB

func ParseDriver(url string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func openDB(url string) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// sql.Open won't check connection, we need ping it later
