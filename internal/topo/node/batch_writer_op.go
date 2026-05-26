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
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

// BatchWriterOp is a streaming writer to convert batch data into bytes in streaming way
// Immutable: false
// Input: any (mostly MessageTuple/SinkTupleList, may receive RawTuple after transformOp). Batch EOF is a signal to flush the buffer.
// Output: RawTuple
type BatchWriterOp struct {
	*defaultSinkNode
	writer message.ConvertWriter
	// save lastRow to get the props
	lastRow any
}

func NewBatchWriterOp(ctx api.StreamContext, name string, rOpt *def.RuleOption, schema map[string]*ast.JsonStreamField, sc *SinkConf) (*BatchWriterOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec decode op receives map/[]map and converts it to bytes.
// If receiving bytes, just return it.
func (o *BatchWriterOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// if batch EOF, flush the buffer

// TODO trace for batch

// sendBatchEnd out raw bytes
// create a new file

func (o *BatchWriterOp) ingest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}
