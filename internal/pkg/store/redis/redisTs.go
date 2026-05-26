// Copyright 2021-2023 EMQ Technologies Co., Ltd.
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

//go:build redisdb || !core

package redis

import (
	"github.com/redis/go-redis/v9"
)

const (
	TsPrefix             = "KV:TS"
	AddToSortedSet       = "ZADD"
	ReversedRangeByScore = "ZREVRANGEBYSCORE"
	RemoveRangeByScore   = "ZREMRANGEBYSCORE"
	Delete               = "DEL"
	ReversedRange        = "ZREVRANGE"
)

type ts struct {
	db    *redis.Client
	table string
	last  int64
	key   string
}

func createRedisTs(redis *redis.Client, table string) (*ts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ts) Set(key int64, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *ts) Get(key int64, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *ts) Last(value interface{}) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (t *ts) Delete(key int64) error { _ = "STUB: not implemented"; return nil }

func (t *ts) DeleteBefore(key int64) error { _ = "STUB: not implemented"; return nil }

func (t *ts) Close() error { _ = "STUB: not implemented"; return nil }

func (t *ts) Drop() error { _ = "STUB: not implemented"; return nil }

func getLast(db *redis.Client, key string, value interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
