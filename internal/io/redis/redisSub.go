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

type redisSub struct {
	conf *redisSubConfig
	conn *redis.Client
}

type redisSubConfig struct {
	Address  string   `json:"address"`
	Db       int      `json:"db"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Channels []string `json:"channels"`
}

func (r *redisSub) Validate(props map[string]any) error { _ = "STUB: not implemented"; return nil }

func (r *redisSub) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisSub) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisSub) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *redisSub) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, _ api.ErrorIngest) error {
	_ = "STUB: not implemented"
	// Subscribe to Redis channels
	return nil
}

func (r *redisSub) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func RedisSub() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var _ util.PingableConn = &redisSub{}
