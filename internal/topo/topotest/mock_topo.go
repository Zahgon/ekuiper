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

package topotest

import (
	"testing"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/testx"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

func init() {
	testx.InitEnv("topotest")
}

const POSTLEAP = 1000 // Time change after all data sends out
type RuleTest struct {
	Name string
	Sql  string
	R    [][]map[string]any // The result
	M    map[string]any     // final metrics
	T    *def.PrintableTopo // printable topo, an optional field
	W    int                // wait time for each data sending, in milli
	TL   int                // table load wait time before first data, in milli (for tests using lookup tables)
}

// CommonResultFunc A function to convert memory sink result to map slice
func CommonResultFunc(result []any) [][]map[string]any { _ = "STUB: not implemented"; return nil }

func DoRuleTest(t *testing.T, tests []RuleTest, opt *def.RuleOption, w int) {
	_ = "STUB: not implemented"
	return
}

func DoRuleTestWithResultFunc(t *testing.T, tests []RuleTest, opt *def.RuleOption, w int, resultFunc func(result []any) [][]map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Drop any stale checkpoint state from a previous run of the same rule ID

// Create the rule which sink to memory topic

// Send data with leaps

// do nothing

// the done signal is sent when all sources are EOF. The mock source is bounded so this will be triggered.

// Fast drain: collect everything immediately available

// One more attempt to catch shutdown-flushed results

func CompareMetrics(tp *topo.Topo, m map[string]interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// do not find

func sendData(dataLength int, datas [][]*xsql.Tuple, tp *topo.Topo, postleap int, wait int, tableLoadWait int) {
	_ = "STUB: not implemented"
	// TODO assume multiple data source send the data in order and has the same length
	return
}

// wait for table to load (only for tests using lookup tables)

// Make sure time is going forward only
// gradually add uptime to ensure checkpoint is triggered before the data send

// create a test rule with memory sink
func createTestRule(t *testing.T, id string, tt RuleTest, opt *def.RuleOption) ([][]*xsql.Tuple, int, *topo.Topo, <-chan error) {
	_ = "STUB: not implemented"
	// Create stream
	return nil, 0, nil, nil
}

// HandleStream Create or drop streams
func HandleStream(createOrDrop bool, names []string, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

type RuleCheckpointTest struct {
	RuleTest
	PauseSize int // Stop stream after sending pauseSize source to test checkpoint resume
	Cc        int // checkpoint count when paused
	// PauseMetric map[string]interface{} // The metric to check when paused
}

func DoCheckpointRuleTest(t *testing.T, tests []RuleCheckpointTest, opt *def.RuleOption, w int) {
	_ = "STUB: not implemented"
	return
}

// Drop any stale checkpoint state from a previous run (but NOT during restart after checkpoint)

// Create the rule which sink to memory topic

// Send data with leaps

// do nothing

// Send async

// compare checkpoint count
// Wait longer than checkpoint interval (2s) to ensure at least one checkpoint completes

// resume stream

// Receive data

// Fast drain: collect everything immediately available

// One more attempt to catch shutdown-flushed results

func waitTopoReady(t *testing.T, tp *topo.Topo, id string) { _ = "STUB: not implemented"; return }
