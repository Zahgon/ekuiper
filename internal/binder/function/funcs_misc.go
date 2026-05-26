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

package function

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

func registerMiscFunc() { _ = "STUB: not implemented"; return }

// directly return in the valuer

// directly return in the valuer

// directly return in the valuer

// directly return in the valuer

func round(num float64) int { _ = "STUB: not implemented"; return 0 }

func toFixed(num float64, precision int) float64 { _ = "STUB: not implemented"; return 0 }

func jsonCall(ctx api.StreamContext, args []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// page Rotate storage for in memory cache
// Not thread safe!
type ringqueue struct {
	Data []any
	H    int
	T    int
	L    int
	Size int
}

func newRingqueue(size int) *ringqueue { _ = "STUB: not implemented"; return nil }

// When deleting, head++, if tail == head, it is empty
// When append, tail++, if tail== head, it is full

// fill item will fill the queue with item value
func (p *ringqueue) fill(item interface{}) { _ = "STUB: not implemented"; return }

// append item if list is not full and return true; otherwise return false
func (p *ringqueue) append(item interface{}) bool {
	_ = "STUB: not implemented"
	// full
	return false
}

// fetch get the first item in the cache and remove
func (p *ringqueue) fetch() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// peek get the first item in the cache but keep it
func (p *ringqueue) peek() (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

func (p *ringqueue) isFull() bool { _ = "STUB: not implemented"; return false }
