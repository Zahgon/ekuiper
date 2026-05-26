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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type ProjectPlan struct {
	baseLogicalPlan
	isAggregate      bool
	allWildcard      bool
	sendMeta         bool
	sendNil          bool
	fields           ast.Fields
	fieldLen         int
	colNames         [][]string
	exceptNames      []string
	wildcardEmitters map[string]bool
	aliasFields      ast.Fields
	exprFields       ast.Fields
	// exprIndices holds the SinkContent output slot for each exprField entry
	// when UseSliceTuple is enabled.  Populated by createLogicalPlanFull.
	exprIndices []int
	enableLimit bool
	limitCount  int
}

func (p ProjectPlan) Init() *ProjectPlan { _ = "STUB: not implemented"; return nil }

// Invisible ExprFields must not appear in the output or in
// exprIndices; exclude them so both slices stay in sync.

func (p *ProjectPlan) BuildExplainInfo() { _ = "STUB: not implemented"; return }

func (p *ProjectPlan) PruneColumns(fields []ast.Expr) error { _ = "STUB: not implemented"; return nil }
