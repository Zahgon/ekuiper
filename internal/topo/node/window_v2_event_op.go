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
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type EventSlidingWindowOp struct {
	*WindowV2Operator
	Delay            time.Duration
	Length           time.Duration
	stateFuncs       []*ast.Call
	triggerCondition ast.Expr
	delayTS          []time.Time
}

func NewEventSlidingWindowOp(o *WindowV2Operator) *EventSlidingWindowOp {
	_ = "STUB: not implemented"
	return nil
}

func (s *EventSlidingWindowOp) exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (o *WindowV2Operator) ingest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// watermark tuple should return
