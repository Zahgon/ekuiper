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

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
)

type BatchOp struct {
	*defaultSinkNode
	// configs
	batchSize      int
	lingerInterval time.Duration
	// state
	currIndex int
}

func NewBatchOp(name string, rOpt *def.RuleOption, batchSize int, lingerInterval time.Duration) (*BatchOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *BatchOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (b *BatchOp) runWithTickerAndBatchSize(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (b *BatchOp) ingest(ctx api.StreamContext, item any, checkSize bool) {
	_ = "STUB: not implemented"
	return
}

// If receive EOF, sendBatchEnd out the result immediately. Only work with single stream

func (b *BatchOp) sendBatchEnd(ctx api.StreamContext) { _ = "STUB: not implemented"; return }

// Reset buffer

func (b *BatchOp) runWithBatchSize(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (b *BatchOp) runWithTicker(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}
