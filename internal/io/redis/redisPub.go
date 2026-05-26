// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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
)

type redisPub struct {
	conf *redisPubConfig
	conn *redis.Client
}

type redisPubConfig struct {
	Address  string `json:"address"`
	Db       int    `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	Channel  string `json:"channel"`
}

func (r *redisPub) Validate(props map[string]any) error { _ = "STUB: not implemented"; return nil }

func (r *redisPub) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisPub) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisPub) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisPub) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	// Publish
	return nil
}

func (r *redisPub) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func RedisPub() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var (
	_ api.BytesCollector = &redisPub{}
	_ util.PingableConn  = &redisPub{}
)
