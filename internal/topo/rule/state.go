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

package rule

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	kctx "github.com/lf-edge/ekuiper/v2/internal/topo/context"
	"github.com/lf-edge/ekuiper/v2/internal/topo/rule/machine"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// State control the Rule RunState
// Created when loading from DB or creating. Deleted when Rule deleting
// May be accessed by multiple go routines, receiving concurrent request to change the RunState
type State struct {
	// A singleton for state, create at new and never change
	ruleLock syncx.RWMutex
	// Nearly constant, only change when update the Rule
	// It is used to construct topo
	Rule          *def.Rule
	logger        api.Logger
	updateTrigger func(string, bool)
	// The physical rule instance for each **run**. control the lifecycle in State.
	topology *topo.Topo
	// temporary storage for topo graph to make sure even Rule close, the graph is still available
	topoGraph *def.PrintableTopo
	// Metric RunState
	stoppedMetrics []any
	// State machine
	sm machine.StateMachine
}

// NewState provision a state instance only.
// Do not plan or run as before. If the Rule is not triggered, do not plan or run.
// When called by recover Rule, expect
func NewState(rule *def.Rule, updateTriggerFunc func(string, bool)) *State {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) GetRule() *def.Rule { _ = "STUB: not implemented"; return nil }

func (s *State) SetRule(r *def.Rule) { _ = "STUB: not implemented"; return }

// ValidateAndRun tries to set up the rule in an atomic way
// It is the only way to update the state rule.
// 1. validate the new rule
// 2. set the state rule property and create the new topo
// 3. stop and clean the old topo if any
// 4. run the new topo
// Notice that, the return err is VALIDATION error only. Run error is async and checked from rule status
// Topo side effect: 1.This function will create and store a new topo if no validation error
// 2. If there is validation error, this function will destroy the new topo
func (s *State) ValidateAndRun(newRule *def.Rule) error { _ = "STUB: not implemented"; return nil }

func (s *State) Bootstrap() error { _ = "STUB: not implemented"; return nil }

// Start run start or add the start action to queue
// By check state, it assures only one Start function is running at any time. (thread safe)
// regSchedule: whether need to handle scheduler. If call externally, set it to true
func (s *State) Start() error { _ = "STUB: not implemented"; return nil }

// delegate to rule patrol checker

func (s *State) ScheduleStart() error { _ = "STUB: not implemented"; return nil }

// doStart trigger the Rule run. If no trigger error, the Rule will run async and control the state by itself

// Stop run stop action or add the stop action to queue
// regSchedule: whether need to handle scheduler. If call externally, set it to true
func (s *State) Stop() { _ = "STUB: not implemented"; return }

func (s *State) ScheduleStop() { _ = "STUB: not implemented"; return }

// do stop, stopping action and starting action are mutual exclusive. No concurrent problem here

func (s *State) StopWithLastWill(msg string) { _ = "STUB: not implemented"; return }

func (s *State) Delete() { _ = "STUB: not implemented"; return }

func (s *State) GetState() machine.RunState {
	_ = "STUB: not implemented"
	return *new(machine.RunState)
}

func (s *State) GetStartTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *State) GetSchema() (map[string]*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStatusMessage return the current RunState of the Rule
// No set is provided, RunState are changed according to the action (start, stop)
func (s *State) GetStatusMessage() string { _ = "STUB: not implemented"; return "" }

// Compose status line

// Compose run timing metrics

// Compose metrics

func (s *State) GetStatusMap() map[string]any { _ = "STUB: not implemented"; return nil }

// Compose metrics

func (s *State) GetTopoGraph() *def.PrintableTopo { _ = "STUB: not implemented"; return nil }

func (s *State) SetIsTraceEnabled(isEnabled bool, stra kctx.TraceStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *State) IsTraceEnabled() bool { _ = "STUB: not implemented"; return false }

func (s *State) GetMetrics() ([]string, []any) { _ = "STUB: not implemented"; return nil, nil }

func (s *State) GetStreams() []string { _ = "STUB: not implemented"; return nil }

func (s *State) GetLastWill() string { _ = "STUB: not implemented"; return "" }

func (s *State) ResetStreamOffset(name string, input map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}
