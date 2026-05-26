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

package sql

import (
	"database/sql"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	client2 "github.com/lf-edge/ekuiper/v2/extensions/impl/sql/client"
	"github.com/lf-edge/ekuiper/v2/extensions/impl/sql/sqldatabase/sqlgen"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type SQLSourceConnector struct {
	id            string
	conf          *SQLConf
	Query         sqlgen.SqlQueryGenerator
	conn          *client2.SQLConnection
	props         map[string]any
	needReconnect bool
	conId         string
	columns       []interface{}
	stats         *sqlSourceStats
	ruleID        string
	opID          string
}

type sqlSourceStats struct {
	totalScanIntoMapDuration time.Duration
	totalWaitDuration        time.Duration
}

func (s *SQLSourceConnector) Ping(ctx api.StreamContext, m map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSourceConnector) resetStats() { _ = "STUB: not implemented"; return }

func (s *SQLSourceConnector) updateMetrics() { _ = "STUB: not implemented"; return }

type SQLConf struct {
	Interval            cast.DurationConf           `json:"interval"`
	DBUrl               string                      `json:"dburl"`
	URL                 string                      `json:"url,omitempty"`
	Datasource          string                      `json:"datasource"`
	TemplateSqlQueryCfg *sqlgen.TemplateSqlQueryCfg `json:"templateSqlQueryCfg"`
}

func init() {
	modules.RegisterConnection("sql", client2.CreateConnection)
}

func (s *SQLSourceConnector) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSourceConnector) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSourceConnector) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSourceConnector) Pull(ctx api.StreamContext, recvTime time.Time, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

func (s *SQLSourceConnector) queryData(ctx api.StreamContext, rcvTime time.Time, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

func (s *SQLSourceConnector) GetOffset() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SQLSourceConnector) Rewind(offset interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLSourceConnector) ResetOffset(input map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func scanIntoMap(mapValue map[string]interface{}, values []interface{}, columns []string, stats *sqlSourceStats) {
	_ = "STUB: not implemented"
	return
}

func prepareValues(ctx api.StreamContext, values []interface{}, columnTypes []*sql.ColumnType, columns []string) {
	_ = "STUB: not implemented"
	return
}

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var (
	_ api.PullTupleSource = &SQLSourceConnector{}
	_ util.PingableConn   = &SQLSourceConnector{}
)

func (sc *SQLConf) resolveDBURL(props map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildScanValueByColumnType(ctx api.StreamContext, colName, colType string, nullable bool) interface{} {
	_ = "STUB: not implemented"
	return nil
}
