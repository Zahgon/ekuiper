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

package pubsub

import (
	"regexp"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const IdProperty = "topic"

type pubConsumers struct {
	count             int
	consumers         map[string]chan any // The consumer channel list [sourceId]chan, the value must be message or message list
	consumersReplaced map[string]int
}

type subChan struct {
	regex *regexp.Regexp
	ch    chan any
}

var (
	pubTopics = make(map[string]*pubConsumers)
	subExps   = make(map[string]*subChan)
	mu        = syncx.RWMutex{}
)

func CreatePub(topic string) { _ = "STUB: not implemented"; return }

func CreateSub(wildcard string, regex *regexp.Regexp, sourceId string, bufferLength int) chan any {
	_ = "STUB: not implemented"
	return nil
}

func CloseSourceConsumerChannel(topic string, sourceId string) { _ = "STUB: not implemented"; return }

func RemovePub(topic string) { _ = "STUB: not implemented"; return }

func ProduceAny(ctx api.StreamContext, topic string, data any) { _ = "STUB: not implemented"; return }

func Produce(ctx api.StreamContext, topic string, data MemTuple) { _ = "STUB: not implemented"; return }

func ProduceList(ctx api.StreamContext, topic string, list []MemTuple) {
	_ = "STUB: not implemented"
	return
}

func ProduceError(ctx api.StreamContext, topic string, err error) {
	_ = "STUB: not implemented"
	return
}

func doProduce(ctx api.StreamContext, topic string, data any) { _ = "STUB: not implemented"; return }

// broadcast to all consumers

// rule stop so stop waiting

func addPubConsumer(topic string, sourceId string, ch chan any) { _ = "STUB: not implemented"; return }

// If already exist, it is usually the rule is restarting and the previous handle is not released yet
// Just use the latest ch as the handle. Also record the replaced status so that it won't remove all handles during removal of the previous handle

func removePubConsumer(topic string, sourceId string, c *pubConsumers) {
	_ = "STUB: not implemented"
	return
}

// GetPubCount returns the number of producers for a topic. For testing only.
func GetPubCount(topic string) int { _ = "STUB: not implemented"; return 0 }

// Reset For testing only
func Reset() { _ = "STUB: not implemented"; return }
