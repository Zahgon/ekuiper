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

package node

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.opentelemetry.io/otel/trace"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/node/cache"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

// CacheOp receives tuples and decide to send through or save to disk. Run right before sink
// Immutable: true
// Input: any (mostly MessageTuple/MessageTupleList, may receive RawTuple after transformOp)
// Special validation: one output only
type CacheOp struct {
	*defaultSinkNode
	// configs
	cacheConf *model.SinkConf
	// state
	cache    *cache.SyncCache
	currItem any
	hasCache bool
	// send timer, only enabled when there is cache. disable when all cache are sent
	resendTicker  *clock.Ticker
	resendTimerCh <-chan time.Time
	// trace span map. need to save it until it is sent because cache op does not send out one by one
	rowHandle map[any]trace.Span
}

func NewCacheOp(ctx api.StreamContext, name string, rOpt *def.RuleOption, sc *model.SinkConf) (*CacheOp, error) {
	_ = "STUB: not implemented"
	// use channel buffer as memory cache
	return nil, nil
}

// Exec ingest data and send through.
// If channel full, save data to disk cache and start send timer
// Once all cache sent, stop send timer
func (s *CacheOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// If already have the cache, append this to cache and send the currItem
// Otherwise, send out the new data. If blocked, make it currItem
// already have cache, add current data to cache and send out the cache

func (s *CacheOp) send() { _ = "STUB: not implemented"; return }

// current item sent out finally

// read

// should never happen

// cancel the timer since all cache are sent

// Send by custom broadcast, if successful, reset currItem to nil

func (s *CacheOp) doBroadcast(val interface{}) { _ = "STUB: not implemented"; return }

// send through. The sink must retry until successful

// rule stop so stop waiting

// Start the send interval
