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

package store

import (
	"regexp"

	"github.com/lf-edge/ekuiper/v2/internal/io/memory/pubsub"
)

// Reg registers a topic to save it to memory store
// Create a new go routine to listen to the topic and save the data to memory
func Reg(topic string, topicRegex *regexp.Regexp, key string) (*Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runTable should only run in a single instance.
// This go routine is used to accumulate data in memory
// If the go routine close, the go routine exits but the data will be kept until table dropped
func runTable(topic string, topicRegex *regexp.Regexp, t *Table) { _ = "STUB: not implemented"; return }

// should never happen

func ingestMemTuple(t *Table, tuple pubsub.MemTuple) { _ = "STUB: not implemented"; return }

// Unreg unregisters a topic to remove it from memory store
func Unreg(topic string, key string) error {
	_ = "STUB: not implemented"
	// Must be an atomic operation
	return nil
}
