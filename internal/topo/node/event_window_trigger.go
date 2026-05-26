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
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

// EventTimeTrigger scans the input tuples and find out the tuples in the current window
// The inputs are sorted by watermark op
type EventTimeTrigger struct {
	window   *WindowConfig
	interval time.Duration
}

func NewEventTimeTrigger(window *WindowConfig) (*EventTimeTrigger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use timeout to update watermark

// If the window end cannot be determined yet, return max int64 so that it can be recalculated for the next watermark
func (w *EventTimeTrigger) getNextWindow(inputs []xsql.EventRow, current time.Time, watermark time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// first run without a previous window

func (w *EventTimeTrigger) getNextSessionWindow(inputs []xsql.EventRow, now time.Time) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (o *WindowOperator) execEventWindow(ctx api.StreamContext, inputs []xsql.EventRow, _ chan<- error) {
	_ = "STUB: not implemented"
	return
}

// process incoming item

// send the last part

// Session window needs a recalculation of window because its window end depends on the inputs

// scan all events and find out the event in the current window

// send the first part

// first tuple, set the window start time, which will set to triggerTime

// is cancelling

func getEarliestEventTs(inputs []xsql.EventRow, startTs time.Time, endTs time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (o *WindowOperator) ingest(ctx api.StreamContext, item any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// watermark tuple should return
