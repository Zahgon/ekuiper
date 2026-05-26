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

package processor

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

type RuleProcessor struct {
	db           kv.KeyValue
	ruleStatusDb kv.KeyValue
}

func NewRuleProcessor() *RuleProcessor { _ = "STUB: not implemented"; return nil }

func (p *RuleProcessor) ExecCreateWithValidation(name, ruleJson string) (*def.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// old rule has newer version

func (p *RuleProcessor) ExecCreate(name, ruleJson string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *RuleProcessor) ExecUpsert(id, ruleJson string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *RuleProcessor) ExecReplaceRuleState(name string, triggered bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *RuleProcessor) GetRuleJson(id string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *RuleProcessor) GetRuleById(id string) (*def.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRuleByJsonValidated called when the json is getting from trusted source like db
func (p *RuleProcessor) GetRuleByJsonValidated(id, ruleJson string) (*def.Rule, error) {
	_ = "STUB: not implemented"
	return nil,

		// set default rule options
		nil
}

func (p *RuleProcessor) GetRuleByJson(id, ruleJson string) (*def.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validation

// CanReplace compare which version is newer, return true if new version is newer
// If both version are empty, need to replace to be backward compatible
func CanReplace(old, new string) bool { _ = "STUB: not implemented"; return false }

func clone(opt def.RuleOption) *def.RuleOption { _ = "STUB: not implemented"; return nil }

func (p *RuleProcessor) ExecExists(name string) bool { _ = "STUB: not implemented"; return false }

func (p *RuleProcessor) ExecDesc(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Parse into def.Rule and re-marshal to ensure consistent field ordering

func (p *RuleProcessor) GetAllRules() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *RuleProcessor) GetAllRulesJson() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *RuleProcessor) ExecDrop(name string) error { _ = "STUB: not implemented"; return nil }

func cleanCheckpoint(name string) error { _ = "STUB: not implemented"; return nil }

func cleanSinkCache(name string) error { _ = "STUB: not implemented"; return nil }
