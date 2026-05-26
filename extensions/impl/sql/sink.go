// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/extensions/impl/sql/client"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/pkg/connection"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

const (
	LblInsert = "insert"
	LblUpdate = "update"
	LblDel    = "del"
)

type SQLSinkConnector struct {
	config        *sqlSinkConfig
	cw            *connection.ConnWrapper
	conn          *client.SQLConnection
	props         map[string]any
	needReconnect bool
}

type sqlSinkConfig struct {
	*SQLConf
	Table        string   `json:"table"`
	Fields       []string `json:"fields"`
	RowKindField string   `json:"rowKindField"`
	KeyField     string   `json:"keyField"`
}

func (c *sqlSinkConfig) buildInsertSql(ctx api.StreamContext, mapData map[string]interface{}, keys []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *sqlSinkConfig) getValuesByKeys(ctx api.StreamContext, mapData map[string]interface{}, keys []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Escape single quotes by doubling them (SQL standard) to avoid breaking the literal.

func quoteSQLString(s string) string {
	_ = "STUB: not implemented"
	// SQL string literal escaping: ' -> ”.
	return ""
}

func (s *SQLSinkConnector) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Consume This is run after provision. Discard common confs that will only be handled in sink itself
func (s *SQLSinkConnector) Consume(props map[string]any) { _ = "STUB: not implemented"; return }

func (s *SQLSinkConnector) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) Collect(ctx api.StreamContext, item api.MessageTuple) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) collect(ctx api.StreamContext, item map[string]any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) CollectList(ctx api.StreamContext, items api.MessageTupleList) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) collectList(ctx api.StreamContext, items []map[string]any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// save save updatable data only to db
func (s *SQLSinkConnector) save(ctx api.StreamContext, table string, data map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) writeToDB(ctx api.StreamContext, sqlStr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSinkConnector) extractKeys(item map[string]any) []string {
	_ = "STUB: not implemented"
	return nil
}

func buildInsertSQL(table string, keys []string, values []string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var (
	_ api.TupleCollector  = &SQLSinkConnector{}
	_ util.PingableConn   = &SQLSinkConnector{}
	_ model.PropsConsumer = &SQLSinkConnector{}
)
