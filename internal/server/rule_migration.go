// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/processor"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type RuleMigrationProcessor struct {
	r *processor.RuleProcessor
	s *processor.StreamProcessor
}

type InstallScriptGetter interface {
	InstallScript(s string) (string, string)
}

func NewRuleMigrationProcessor(r *processor.RuleProcessor, s *processor.StreamProcessor) *RuleMigrationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func newDependencies() *dependencies { _ = "STUB: not implemented"; return nil }

// dependencies copy all connections related configs by hardcode
type dependencies struct {
	rules            []string
	streams          []string
	tables           []string
	sources          []string
	sinks            []string
	sourceConfigKeys map[string][]string
	sinkConfigKeys   map[string][]string
	functions        []string
	schemas          []string
}

func ruleTraverse(rule *def.Rule, de *dependencies) { _ = "STUB: not implemented"; return }

// streams

// get streams

// get tables

// get source type

// get config key

// get schema id

// actions

// function

// Rules

// get config key

// get schema id

func (p *RuleMigrationProcessor) ConfigurationPartialExport(rules []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *RuleMigrationProcessor) exportRules(rules []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *RuleMigrationProcessor) exportStreams(streams []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *RuleMigrationProcessor) exportTables(tables []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *RuleMigrationProcessor) exportSelected(de *dependencies, config *Configuration) {
	_ = "STUB: not implemented"
	// get the stream and table
	return
}

// get the sources

// get sinks

// get functions

// get sourceCfg/sinkCfg

// get schema

// should never happen

func parsePick(props map[string]interface{}) (*ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFunc(props map[string]interface{}) (*ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFilter(props map[string]interface{}) (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func parseHaving(props map[string]interface{}) (ast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ast.Expr), nil
}

func parseSwitch(props map[string]interface{}) ([]ast.Expr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseOrderBy(props map[string]interface{}) (*ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseGroupBy(props map[string]interface{}) (*ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseJoin(props map[string]interface{}) (*ast.SelectStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
