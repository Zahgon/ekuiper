// Copyright 2022-2023 EMQ Technologies Co., Ltd.
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

package cache

import (
	"context"
	"time"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type item struct {
	data       []map[string]any
	expiration time.Time
}

type Cache struct {
	expireTime      time.Duration
	cacheMissingKey bool
	cancel          context.CancelFunc
	items           map[string]*item
	syncx.RWMutex
}

func NewCache(expireTime time.Duration, cacheMissingKey bool) *Cache {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Cache) deleteExpired() { _ = "STUB: not implemented"; return }

func (c *Cache) Set(key string, value []map[string]any) { _ = "STUB: not implemented"; return }

func (c *Cache) Get(key string) ([]map[string]any, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Cache) Close() { _ = "STUB: not implemented"; return }
