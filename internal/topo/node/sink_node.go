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

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

// SinkNode represents a sink node that collects data from the stream
// It typically only do connect and send. It does not do any processing.
// This node is the skeleton. It will refer to a sink instance to do the real work.
type SinkNode struct {
	*defaultSinkNode
	sink           api.Sink
	eoflimit       int
	currentEof     int
	resendInterval time.Duration
	doCollect      func(ctx api.StreamContext, sink api.Sink, data any) error
	// channel for resend
	resendOut chan<- any
}

// Caching:
// 1. Set cache settings to enable diskCache
// 2. Set resendInterval and bufferLength will use bufferLength as the memory cache
// 3. By default, drop if it cannot sends out.
func newSinkNode(ctx api.StreamContext, name string, rOpt def.RuleOption, eoflimit int, sc *SinkConf, isRetry bool) *SinkNode {
	_ = "STUB: not implemented"
	// set collect retry according to cache setting
	return nil
}

// default retry interval to 100ms

// Sink input channel as buffer

func (s *SinkNode) setKafkaSinkStatsManager(ctx api.StreamContext) {
	_ = "STUB: not implemented"
	return
}

func (s *SinkNode) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// resend handling when enabling cache. Two cases: 1. send to alter queue with resendOUt. 2. retry (blocking) until success or unrecoverable error if resendInterval is set

// do nothing

// rule stop so stop waiting

func (s *SinkNode) SetResendOutput(output chan<- any) { _ = "STUB: not implemented"; return }

func (s *SinkNode) connectionStatusChange(status string, message string) {
	_ = "STUB: not implemented"
	return
}

func (s *SinkNode) ingest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// NewBytesSinkNode creates a sink node that collects data from the stream. Do some static validation
func NewBytesSinkNode(ctx api.StreamContext, name string, sink api.BytesCollector, rOpt def.RuleOption, eoflimit int, sc *SinkConf, isRetry bool) (*SinkNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bytesCollect(ctx api.StreamContext, sink api.Sink, data any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// NewTupleSinkNode creates a sink node that collects data from the stream. Do some static validation
func NewTupleSinkNode(ctx api.StreamContext, name string, sink api.TupleCollector, rOpt def.RuleOption, eoflimit int, sc *SinkConf, isRetry bool) (*SinkNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create converter for decoding RawTuple

// createTupleCollect creates a tupleCollect function with the given converter
func createTupleCollect(conv message.Converter) func(ctx api.StreamContext, sink api.Sink, data any) error {
	_ = "STUB: not implemented"
	return nil
}

// Some tuple list type also implements tuple. So need to handle list firstly

// may receive raw tuple from data template

// decodeAndCollect decodes RawTuple and calls Collect or CollectList based on result
func decodeAndCollect(ctx api.StreamContext, sink api.TupleCollector, d *xsql.RawTuple, conv message.Converter) error {
	_ = "STUB: not implemented"
	return nil
}

var _ DataSinkNode = (*SinkNode)(nil)
