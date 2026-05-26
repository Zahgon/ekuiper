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

package node

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

type DedupTriggerNode struct {
	*defaultSinkNode
	// config
	aliasName  string
	startField string
	endField   string
	nowField   string
	expire     int64
	// state
	requests      PriorityQueue // All the cached events in order
	timeoutTicker *clock.Timer
	timeout       <-chan time.Time
}

func NewDedupTriggerNode(name string, options *def.RuleOption, aliasName string, startField string, endField string, nowField string, expire int64) *DedupTriggerNode {
	_ = "STUB: not implemented"
	return nil
}

func (w *DedupTriggerNode) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// future trigger event

func (w *DedupTriggerNode) trigger(ctx api.StreamContext, now int64) {
	_ = "STUB: not implemented"
	return
}

// trigger by event with timestamp, keep triggering until all history events are triggered

func (w *DedupTriggerNode) rowToReq(d xsql.Row) (*TriggerRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func doTrigger(ctx api.StreamContext, start int64, end int64, now int64, exp int64) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// histogram state, the timeslots which have been taken [{start, end}, {start, end}]

// clean up the expired timeslots

// Find the timeslots which have been taken
// Default to the rightest slot

// calculate timeslots and update histogram for each cases

// In a continuous empty slot

// do nothing

// left empty slot, right empty slot
// left slot + multiple middle empty slots + right slot

// left empty slot, right not empty slot
// left slot + multiple middle empty slots

// left not empty slot, right empty slot
// multiple middle empty slots + right slot

// left not empty slot, right not empty slot

// multiple middle empty slots

type TriggerRequest struct {
	start int64
	end   int64
	now   int64
	exp   int64
	tuple xsql.Row
}

type PriorityQueue []*TriggerRequest

// Push adds an item to the priority queue
func (pq *PriorityQueue) Push(x *TriggerRequest) { _ = "STUB: not implemented"; return }

// Pop removes and returns the item with the highest priority from the priority queue
func (pq *PriorityQueue) Pop() *TriggerRequest { _ = "STUB: not implemented"; return nil }

func (pq *PriorityQueue) Peek() *TriggerRequest { _ = "STUB: not implemented"; return nil }
