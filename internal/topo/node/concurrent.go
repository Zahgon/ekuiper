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
)

// WorkerFunc is the function to process the data
// The function do not need to process error and control messages
// The function must return a slice of data for each input. To omit the data, return nil
type workerFunc func(ctx api.StreamContext, item any) []any

func runWithOrder(ctx api.StreamContext, node *defaultSinkNode, numWorkers int, wf workerFunc) {
	_ = "STUB: not implemented"
	return
}

func runWithOrderAndInterval(ctx api.StreamContext, node *defaultSinkNode, numWorkers int, wf workerFunc, sendInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Start worker goroutines

// start merger goroutine

// Distribute input data to workers

// Merge multiple channels into one preserving the order
func merge(ctx api.StreamContext, node *defaultSinkNode, sendInterval time.Duration, output chan any, channels ...chan []any) {
	_ = "STUB: not implemented"

	// Start a goroutine for each input channel
	return
}

func distribute(ctx api.StreamContext, node *defaultSinkNode, numWorkers int, workerChans []chan any) {
	_ = "STUB: not implemented"
	return
}

// Round-robin

// Just send out all inputs even they are control tuples

func worker(ctx api.StreamContext, node *defaultSinkNode, i int, wf workerFunc, inputRaw chan any, output chan []any) {
	_ = "STUB: not implemented"
	return
}
