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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

type BatchMergerOp struct {
	*defaultSinkNode
	// save lastRow to get the props
	lastRow any
	wt      *xsql.WindowTuples
}

func NewBatchMergerOp(name string, rOpt *def.RuleOption) (*BatchMergerOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec decode op receives map/[]map and converts it to []map.
func (o *BatchMergerOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// TODO: find a way to avoid using ToMaps

func (o *BatchMergerOp) ingest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (o *BatchMergerOp) appendWindowTuples(row xsql.Row) { _ = "STUB: not implemented"; return }
