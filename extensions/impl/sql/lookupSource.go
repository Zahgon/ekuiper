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
	"github.com/lf-edge/ekuiper/contract/v2/api"

	client2 "github.com/lf-edge/ekuiper/v2/extensions/impl/sql/client"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
)

type SqlLookupSource struct {
	conf          *SQLConf
	conn          *client2.SQLConnection
	props         map[string]any
	driver        string
	table         string
	needReconnect bool
	gen           sqlQueryGen
	conId         string
}

func (s *SqlLookupSource) Ping(ctx api.StreamContext, m map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SqlLookupSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SqlLookupSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *SqlLookupSource) Connect(ctx api.StreamContext, sc api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SqlLookupSource) Lookup(ctx api.StreamContext, fields []string, keys []string, values []any) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type sqlQueryGen interface {
	buildQuery(fields []string, keys []string, values []interface{}) string
}

type defaultSQLGen struct{ table string }

func (g defaultSQLGen) buildQuery(fields []string, keys []string, values []interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

type noQuoteSQLGen struct{ table string }

func (g noQuoteSQLGen) buildQuery(fields []string, keys []string, values []interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SqlLookupSource) buildGen() sqlQueryGen {
	_ = "STUB: not implemented"
	return *new(sqlQueryGen)
}

func GetLookupSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var (
	_ api.LookupSource  = &SqlLookupSource{}
	_ util.PingableConn = &SqlLookupSource{}
)
