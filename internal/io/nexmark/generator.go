// Copyright 2025 EMQ Technologies Co., Ltd.
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

package nexmark

import (
	"context"
	"math/rand"
)

type EventGenerator struct {
	ctx        context.Context
	cancel     context.CancelFunc
	eventChan  chan map[string]any
	bufferSize int
	qps        int
	startTS    uint64
	eventID    int64
	r          *rand.Rand
	GenOption
}

type GenOption struct {
	excludePerson  bool
	excludeAuction bool
	excludeBid     bool
}

type WithGenOption func(clientConf *GenOption)

func WithExcludePerson() WithGenOption { _ = "STUB: not implemented"; return *new(WithGenOption) }

func WithExcludeBid() WithGenOption { _ = "STUB: not implemented"; return *new(WithGenOption) }

func WithExcludeAuction() WithGenOption { _ = "STUB: not implemented"; return *new(WithGenOption) }

func NewEventGenerator(parCtx context.Context, qps, bufferSize int, opts ...WithGenOption) *EventGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *EventGenerator) inc() { _ = "STUB: not implemented"; return }

func (g *EventGenerator) genPerson() Person { _ = "STUB: not implemented"; return *new(Person) }

func (g *EventGenerator) genAuction() Auction { _ = "STUB: not implemented"; return *new(Auction) }

func (g *EventGenerator) genBid() Bid { _ = "STUB: not implemented"; return *new(Bid) }

func (g *EventGenerator) GenStream() { _ = "STUB: not implemented"; return }

func (g *EventGenerator) Close() { _ = "STUB: not implemented"; return }
