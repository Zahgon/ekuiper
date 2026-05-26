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

package schema

import (
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type SchemaStore struct {
	syncx.RWMutex
	// rule -> datasource -> schema
	schemaMap map[string]map[string]map[string]*ast.JsonStreamField
	// rule -> datasource -> schema
	wildcardMap map[string]map[string]bool
	// shared stream schema reg, will be used by planner; stream -> sharelayer
	streamMap map[string]SchemaContainer
}

type SchemaContainer interface {
	GetSchema() map[string]*ast.JsonStreamField
	GetSchemaIndex() map[string]int
}

type schemaWrapper map[string]*ast.JsonStreamField

func (s schemaWrapper) GetSchemaIndex() map[string]int { _ = "STUB: not implemented"; return nil }

func (s schemaWrapper) GetSchema() map[string]*ast.JsonStreamField {
	_ = "STUB: not implemented"
	return nil
}

func initStore() *SchemaStore { _ = "STUB: not implemented"; return nil }

var GlobalSchemaStore = initStore()

type RuleSchemaResponse struct {
	// streamName -> schema
	Schema map[string]map[string]*ast.JsonStreamField `json:"schema"`
	// streamName -> wildcard
	Wildcard map[string]bool `json:"wildcard"`
}

func GetRuleSchema(ruleID string) RuleSchemaResponse {
	_ = "STUB: not implemented"
	return *new(RuleSchemaResponse)
}

func AddRuleSchema(ruleID, dataSource string, schema map[string]*ast.JsonStreamField, isWildcard bool) {
	_ = "STUB: not implemented"
	return
}

func RemoveRuleSchema(ruleID string) { _ = "STUB: not implemented"; return }

func GetStream(name string) SchemaContainer {
	_ = "STUB: not implemented"
	return *new(SchemaContainer)
}

func GetStreamSchema(name string) (map[string]*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetStreamSchemaIndex(streamName string) map[string]int { _ = "STUB: not implemented"; return nil }

func AddStaticStream(streamName string, schema map[string]*ast.JsonStreamField) {
	_ = "STUB: not implemented"
	return
}

func RemoveStreamSchema(streamName string) { _ = "STUB: not implemented"; return }
