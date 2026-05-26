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

package tdengine3

import (
	"database/sql"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	_ "github.com/taosdata/driver-go/v3/taosWS"
)

type TaosConfig struct {
	ProvideTs    bool     `json:"provideTs"`
	Port         int      `json:"port"`
	Host         string   `json:"host"`
	User         string   `json:"user"`
	Password     string   `json:"password"`
	Database     string   `json:"database"`
	Table        string   `json:"table"`
	TsFieldName  string   `json:"tsFieldName"`
	Fields       []string `json:"fields"`
	STable       string   `json:"sTable"`
	TagFields    []string `json:"tagFields"`
	DataTemplate string   `json:"dataTemplate"`
	DataField    string   `json:"dataField"`
}

type tdengineSink3 struct {
	cfg *TaosConfig
	cli *sql.DB
}

func (t *tdengineSink3) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tdengineSink3) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tdengineSink3) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tdengineSink3) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (t *tdengineSink3) Collect(ctx api.StreamContext, item api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tdengineSink3) CollectList(ctx api.StreamContext, items api.MessageTupleList) error {
	_ = "STUB: not implemented"
	return nil
}

func (cfg *TaosConfig) buildSql(item api.MessageTuple) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func contains(slice []string, target string) bool { _ = "STUB: not implemented"; return false }

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }
