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

package server

import (
	"context"
	"time"

	"github.com/Rookiecom/cpuprofile"
)

func initRuleset() {
	_ = "STUB: not implemented"
	// init data firstly, so that same version can take precedence
	return
}

func initFromLoc(loc string) error { _ = "STUB: not implemented"; return nil }

// Only leave one initialized file each time. Due to the time shift in some system, compare time is not a good idea

// delete all signal files

// create the unique file

// findInitializedTime finds one files starting with "initialized" and returns
// the int64 suffix value according to the rules:
// - No matching files: -1
// - Matching file with no numeric suffix: 0
// - Otherwise, the int64 suffix value
func findInitializedTime(root string) int64 { _ = "STUB: not implemented"; return 0 }

// Walk through the directory tree
// Default: no files found

// Check if it's a file and starts with "initialized"

// Extract the suffix after "initialized"

// No suffix, return 0

// Try to parse the suffix as an int64

// Valid suffix, return it

// Invalid suffix treated as no suffix

// Stop walking after first match

func resetAllRules() error { _ = "STUB: not implemented"; return nil }

func resetAllStreams() error { _ = "STUB: not implemented"; return nil }

func runScheduleRuleCheckerByInterval(d time.Duration, ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func runScheduleRuleChecker(ctx context.Context) { _ = "STUB: not implemented"; return }

type RuleStatusMetricsValue int

const (
	RuleStoppedByError RuleStatusMetricsValue = -1
	RuleStopped        RuleStatusMetricsValue = 0
	RuleRunning        RuleStatusMetricsValue = 1
)

func handleAllRuleStatusMetrics(rs []ruleWrapper) { _ = "STUB: not implemented"; return }

func handleAllScheduleRuleState(now time.Time, rs []ruleWrapper) { _ = "STUB: not implemented"; return }

// handle auto restart rules

func handleScheduleRuleState(now time.Time, rw ruleWrapper) error {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

type scheduleRuleAction int

const (
	scheduleRuleActionDoNothing scheduleRuleAction = iota
	scheduleRuleActionStart
	scheduleRuleActionStop
	doStop
)

func handleScheduleRule(now time.Time, rw ruleWrapper) scheduleRuleAction {
	_ = "STUB: not implemented"
	return *new(scheduleRuleAction)
}

func scheduleCronRuleAction(now time.Time, rw ruleWrapper) scheduleRuleAction {
	_ = "STUB: not implemented"
	return *new(scheduleRuleAction)
}

type Profiler interface {
	StartCPUProfiler(context.Context, time.Duration) error
	EnableWindowAggregator(int)
	GetWindowData() cpuprofile.DataSetAggregateMap
	RegisterTag(string, chan *cpuprofile.DataSetAggregate)
}

type ekuiperProfile struct{}

func (e *ekuiperProfile) StartCPUProfiler(ctx context.Context, t time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ekuiperProfile) EnableWindowAggregator(window int) { _ = "STUB: not implemented"; return }

func (e *ekuiperProfile) GetWindowData() cpuprofile.DataSetAggregateMap {
	_ = "STUB: not implemented"
	return *new(cpuprofile.DataSetAggregateMap)
}

func (e *ekuiperProfile) RegisterTag(tag string, receiveChan chan *cpuprofile.DataSetAggregate) {
	_ = "STUB: not implemented"
	return
}

func StartCPUProfiling(ctx context.Context, cpuProfile Profiler, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Enable window aggregator to allow GetWindowData() to work without panic.
// Window size of 5 means aggregating data over 5 profiling intervals.

func waitAllRuleStop() { _ = "STUB: not implemented"; return }
