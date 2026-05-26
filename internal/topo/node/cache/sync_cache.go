// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

// SyncCache is the struct to handle cache saving and read
// The data are sink tuples: MessageTuple, MessageTupleList or RawTuple

// page Rotates storage for in memory cache
// Not thread safe!
type page struct {
	Data []any
	H    int
	T    int
	L    int
	Size int
}

// newPage create a new cache page
func newPage(size int) *page { _ = "STUB: not implemented"; return nil }

// When deleting, head++, if tail == head, it is empty
// When append, tail++, if tail== head, it is full

// append item if list is not full and return true; otherwise return false
func (p *page) append(item any) bool {
	_ = "STUB: not implemented"
	// full
	return false
}

// peak get the first item in the cache
func (p *page) peak() (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (p *page) delete() bool { _ = "STUB: not implemented"; return false }

func (p *page) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (p *page) reset() { _ = "STUB: not implemented"; return }

const (
	syncCacheLength = "length"
	syncCacheAdd    = "add"
	syncCachePop    = "pop"
	syncCacheFlush  = "flush"
	syncCacheDrop   = "drop"
	syncCacheLoad   = "load"
)

type SyncCache struct {
	RuleID string
	OpID   string
	// cache config
	cacheConf   *model.SinkConf
	maxDiskPage int
	// cache storage
	writeBufferPage *page
	readBufferPage  *page
	// status
	diskSize     int // the count of pages has been saved
	CacheLength  int // readonly, for metrics only to save calculation
	diskPageTail int // init from the database
	diskPageHead int
	// serialize
	store kv.KeyValue
}

func NewSyncCache(ctx api.StreamContext, cacheConf *model.SinkConf) (*SyncCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The maximum pages in disk. This includes readBuffer, all disk page and write buffer. When flush, all save into disk

// add one more slot so that there will be at least one slot between head and tail to find out the head/tail id

func (c *SyncCache) InitStore(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (c *SyncCache) SetupMeta(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// AddCache not thread safe!
func (c *SyncCache) AddCache(ctx api.StreamContext, item any) error {
	_ = "STUB: not implemented"
	return nil
}

// cool page full, save to disk

func (c *SyncCache) appendWriteCache(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

// disk full, replace read buffer page

// also delete read buffer which is even older

// rotate

func (c *SyncCache) insertReadCache(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

// insert before current head

// If not full

// PopCache not thread safe!
func (c *SyncCache) PopCache(ctx api.StreamContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// read from disk or cool list

// use cool page as the new page

// loaded means whether load the page to memory or just drop
func (c *SyncCache) deleteDiskPage(ctx api.StreamContext, loaded bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SyncCache) loadFromDisk(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

// load page from the disk

// caution, must create a new page instance

func (c *SyncCache) initStore(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// restore the sink cache from disk

// Save 0 when init and save 1 when close. Wait for close for newly started sink node

// may be saving

// restore the disk cache

// no disk cache

// Flush save memory states to disk.
func (c *SyncCache) Flush(ctx api.StreamContext) { _ = "STUB: not implemented"; return }
