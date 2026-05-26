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

package redis

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/redis/go-redis/v9"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
)

type config struct {
	// host:port address.
	Addr     string `json:"addr,omitempty"`
	Username string `json:"username,omitempty"`
	// Optional password. Must match the password specified in the
	Password string `json:"password,omitempty"`
	// Database to be selected after connecting to the server.
	Db int `json:"db,omitempty"`
	// key of field
	Field string `json:"field,omitempty"`
	// key define
	Key          string            `json:"key,omitempty"`
	KeyType      string            `json:"keyType,omitempty"`
	DataType     string            `json:"dataType,omitempty"`
	Expiration   cast.DurationConf `json:"expiration,omitempty"`
	RowkindField string            `json:"rowkindField"`
	DataTemplate string            `json:"dataTemplate"`
	Fields       []string          `json:"fields"`
	DataField    string            `json:"dataField"`
}

type RedisSink struct {
	c   *config
	cli *redis.Client
}

func (r *RedisSink) Provision(_ api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RedisSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// use default DB

func (r *RedisSink) Validate(props map[string]any) error { _ = "STUB: not implemented"; return nil }

func (r *RedisSink) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// use default DB

func (r *RedisSink) Collect(ctx api.StreamContext, item api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RedisSink) CollectList(ctx api.StreamContext, items api.MessageTupleList) error {
	_ = "STUB: not implemented"
	// TODO handle partial error
	return nil
}

func (r *RedisSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (r *RedisSink) save(ctx api.StreamContext, data map[string]any) error {
	_ = "STUB: not implemented"
	return nil

	// prepare key value pairs
}

// get action type

// set key value pairs

// never happen

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var (
	_ api.TupleCollector = &RedisSink{}
	_ util.PingableConn  = &RedisSink{}
)
