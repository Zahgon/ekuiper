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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

// JoinAlignNode will block the stream and buffer all the table tuples. Once buffered, it will combine the later input with the buffer
// The input for batch table MUST be *WindowTuples
type JoinAlignNode struct {
	*defaultSinkNode
	// table states
	batch map[string][]*xsql.Tuple
	size  map[string]int
}

const BatchKey = "$$batchInputs"

func NewJoinAlignNode(name string, emitters []string, sizes []int, options *def.RuleOption) (*JoinAlignNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *JoinAlignNode) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// restore batch state

// process incoming item from both streams(transformed) and tables

func (n *JoinAlignNode) alignBatch(ctx api.StreamContext, input any) {
	_ = "STUB: not implemented"
	return
}
