// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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
	"time"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo/rule"
	"github.com/lf-edge/ekuiper/v2/internal/topo/rule/machine"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// Rule storage includes kv and in memory registry
// Kv stores the rule text with *expected* status so that the rule can be restored after restart
// Registry stores the current rule state in runtime
// Here registry is the in memory registry
var registry *RuleRegistry

type RuleRegistry struct {
	syncx.RWMutex
	internal map[string]*rule.State
}

//// registry and db level state change functions

func (rr *RuleRegistry) update(key string, ruleJson string, value *rule.State) error {
	_ = "STUB: not implemented"
	return nil
}

// load the entry of a rule by id. It is used to get the current rule state
// or send command to a running rule
func (rr *RuleRegistry) load(key string) (value *rule.State, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (rr *RuleRegistry) keys() (keys []string) { _ = "STUB: not implemented"; return nil }

// save registers rule to in-memory registry and persists to DB atomically.
// It fails if the rule already exists in DB.
func (rr *RuleRegistry) save(key string, ruleJson string, value *rule.State) error {
	_ = "STUB: not implemented"
	return nil

	// Persist to DB first - ExecCreate fails if already exists
}

// Update registry only after successful DB write

// only register. It is called when recover from db
func (rr *RuleRegistry) register(key string, value *rule.State) { _ = "STUB: not implemented"; return }

func (rr *RuleRegistry) updateTrigger(id string, trigger bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (rr *RuleRegistry) delete(key string) (*rule.State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//// APIs for REST service
//// Keep consistent by DB. Rollback when db errors happen

func (rr *RuleRegistry) CreateRule(name, ruleJson string) (id string, err error) {
	_ = "STUB: not implemented"
	// Validate the rule json
	return "", nil
}

// create state and save

// Validate the topo

// Store to registry and KV

// rollback clean up

// RecoverRule loads in imported rule.
// Unlike creation, 1. it supposes the rule is valid thus, it will always create the rule state in registry
// 2. It does not handle rule saving to db.
func (rr *RuleRegistry) RecoverRule(r *def.Rule) string { _ = "STUB: not implemented"; return "" }

// Start the rule which runs async

// UpsertRule validates the new rule, then update the db, then restart the rule
// The entire operation is protected by a lock to ensure atomic version checking.
func (rr *RuleRegistry) UpsertRule(ruleId, ruleJson string) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate the rule json (can be done outside lock - no state change)

// Hold lock for entire operation to ensure atomic version check

// do upsert.

// if not exist, create it

// Version check is now atomic with the rest of the operation

// Persist directly - we already hold the lock

// Temp rule, just register in memory

// rollback clean up

func (rr *RuleRegistry) DeleteRule(name string) error {
	_ = "STUB: not implemented"
	// lock registry and db. rs level has its own lock
	return nil
}

func (rr *RuleRegistry) StartRule(name string) error { _ = "STUB: not implemented"; return nil }

func (rr *RuleRegistry) StopRule(name string) error { _ = "STUB: not implemented"; return nil }

func (rr *RuleRegistry) RestartRule(name string) error { _ = "STUB: not implemented"; return nil }

func (rr *RuleRegistry) GetAllRuleStatus() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rr *RuleRegistry) GetAllRulesWithStatus() ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mergeAndSortStrings merges two string slices, removes duplicates, and sorts the result.
func mergeAndSortStrings(slice1, slice2 []string) []string {
	_ = "STUB: not implemented"
	// Step 1: Merge the two slices
	return nil
}

// Pre-allocate capacity

// Step 2: Remove duplicates using a map
// Using struct{} for a memory-efficient "set"
// Pre-allocate capacity based on merged length (upper bound)

// If element not seen yet
// Mark as seen
// Add to unique slice

// Step 3: Sort the unique slice
// sort.Strings sorts a slice of strings in ascending order

func (rr *RuleRegistry) GetRuleStatus(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rr *RuleRegistry) GetRuleStatusV2(name string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rr *RuleRegistry) GetRuleTopo(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rr *RuleRegistry) GetRuleSinkSchema(name string) (map[string]*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rr *RuleRegistry) ValidateRule(name, ruleJson string) ([]string, bool, error) {
	_ = "STUB: not implemented"
	// Validate the ruleDef json
	return nil, false, nil
}

/// Rule Scheduler internal API

func (rr *RuleRegistry) scheduledStart(name string) error { _ = "STUB: not implemented"; return nil }

func (rr *RuleRegistry) scheduledStop(name string) error { _ = "STUB: not implemented"; return nil }

func (rr *RuleRegistry) stopAtExit(name string, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

//// Util functions

func getRuleExceptionStatus(name string) (ruleExceptionStatus, error) {
	_ = "STUB: not implemented"
	return *new(ruleExceptionStatus), nil
}

func getTargetException(keys []string, values []any, prefix string) (int64, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

type ruleExceptionStatus struct {
	Status            string `json:"status"`
	LastException     string `json:"last_exception"`
	ExceptionsTotal   int64  `json:"exceptions_total"`
	lastExceptionTime int64
}

type ruleWrapper struct {
	rule      *def.Rule
	state     machine.RunState
	startTime time.Time
}

func getAllRulesWithState() ([]ruleWrapper, error) { _ = "STUB: not implemented"; return nil, nil }

func getRuleState(name string) (machine.RunState, error) {
	_ = "STUB: not implemented"
	return *new(machine.RunState), nil
}

func deleteRuleData(name string) { _ = "STUB: not implemented"; return }
