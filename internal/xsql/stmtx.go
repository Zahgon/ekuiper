// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

package xsql

import (
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

func GetStreams(stmt *ast.SelectStatement) (result []string) { _ = "STUB: not implemented"; return nil }

// TODO sources must be a stream

func GetStatementFromSql(sql string) (stmt *ast.SelectStatement, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type StreamInfo struct {
	StreamType ast.StreamType `json:"streamType"`
	StreamKind string         `json:"streamKind"`
	Statement  string         `json:"statement"`
	Temp       bool           `json:"temp,omitempty"`
}

func GetDataSourceStatement(m kv.KeyValue, name string) (*StreamInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetDataSource(m kv.KeyValue, name string) (stmt *ast.StreamStmt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
