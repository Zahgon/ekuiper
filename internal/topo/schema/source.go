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

package schema

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type SharedLayer struct {
	syncx.RWMutex
	schema map[string]*ast.JsonStreamField
	// ruleID -> schema
	// Save the schemainfo for each rule only to use when need to attach schema when the rule is starting.
	// Get updated if the rule is updated. Never delete it until the subtopo is deleted.
	reg map[string]schemainfo
	// ruleID -> StreamName
	streamMap map[string]string
	// ruleID -> wildcard
	wildcardMap map[string]struct{}
	// field name -> index; append only, do not detach. Use this to assign source index for fields
	indexMap map[string]int
}

type schemainfo struct {
	datasource string
	schema     map[string]*ast.JsonStreamField
	isWildcard bool
}

func newSharedLayer() *SharedLayer { _ = "STUB: not implemented"; return nil }

func (s *SharedLayer) RegSchema(ruleID, dataSource string, schema map[string]*ast.JsonStreamField, isWildCard bool) {
	_ = "STUB: not implemented"
	return
}

func (s *SharedLayer) updateReg() { _ = "STUB: not implemented"; return }

func (s *SharedLayer) Attach(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// If it is wildcard, the schema is already the biggest, no need to recalculate

func (s *SharedLayer) Detach(ctx api.StreamContext, isClose bool) error {
	_ = "STUB: not implemented"
	return nil
}

// If any remaining rule is still wildcard, keep schema as nil (schemaless).
// Merging nil wildcard schemas would produce {} which causes the decoder to
// reject every field and emit empty tuples.

func (s *SharedLayer) GetSchema() map[string]*ast.JsonStreamField {
	_ = "STUB: not implemented"
	return nil

	//if len(s.wildcardMap) > 0 {
	//	return nil
	//}
}

func (s *SharedLayer) GetSchemaIndex() map[string]int { _ = "STUB: not implemented"; return nil }

func (s *SharedLayer) merge(originSchema, newSchema map[string]*ast.JsonStreamField) (map[string]*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update index map

func mergeSchema(originSchema, newSchema map[string]*ast.JsonStreamField) (map[string]*ast.JsonStreamField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
