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

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type SQLConnection struct {
	syncx.RWMutex
	url    string
	db     *sql.DB
	id     string
	closed bool
}

func (s *SQLConnection) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLConnection) GetId(ctx api.StreamContext) string { _ = "STUB: not implemented"; return "" }

func (s *SQLConnection) Dial(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *SQLConnection) Reconnect() error { _ = "STUB: not implemented"; return nil }

func (s *SQLConnection) GetDB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (s *SQLConnection) Ping(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *SQLConnection) DetachSub(ctx api.StreamContext, props map[string]any) {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (s *SQLConnection) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func CreateConnection(ctx api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func (s *SQLConnection) dial(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }
