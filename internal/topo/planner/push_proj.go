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

package planner

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

type pushProjectionPlan struct{}

// pushProjectionPlan inject Projection Plan between the shared Datasource and its father only if the Plan have windowPlan
// We use Projection to remove the unused column before windowPlan in order to reduce memory consuming
func (pp *pushProjectionPlan) optimize(plan LogicalPlan, _ *def.RuleOption) (LogicalPlan, error) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), nil
}

func (pp *pushProjectionPlan) searchWindowPlan(plan LogicalPlan) bool {
	_ = "STUB: not implemented"
	return false
}

func (pp *pushProjectionPlan) searchJoinPlan(plan LogicalPlan) bool {
	_ = "STUB: not implemented"
	return false
}

func (pp *pushProjectionPlan) pushProjection(ctx *searchCtx) { _ = "STUB: not implemented"; return }

func buildFields(ds *DataSourcePlan) []ast.Field { _ = "STUB: not implemented"; return nil }

func (pp *pushProjectionPlan) name() string { _ = "STUB: not implemented"; return "" }

type pushAliasDecode struct{}

func (p *pushAliasDecode) optimize(plan LogicalPlan, option *def.RuleOption) (LogicalPlan, error) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), nil
}

func (p *pushAliasDecode) name() string { _ = "STUB: not implemented"; return "" }

type searchCtx struct {
	find               []*sharedSource
	noSharedDatasource []*DataSourcePlan
}

type sharedSource struct {
	ds     *DataSourcePlan
	father LogicalPlan
}

func searchSharedDataSource(ctx *searchCtx, plan, father LogicalPlan) {
	_ = "STUB: not implemented"
	return
}

func searchNoSharedDatasource(ctx *searchCtx, plan LogicalPlan) { _ = "STUB: not implemented"; return }
