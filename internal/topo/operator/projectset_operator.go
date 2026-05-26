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

package operator

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

type ProjectSetOperator struct {
	SrfMapping  map[string]struct{}
	EnableLimit bool
	LimitCount  int
}

// Apply implement UnOperation
// ProjectSetOperator will extract the results from the set-returning-function into multi rows by aligning other columns
// For tuple, ProjectSetOperator will do the following transform:
// {"a":[1,2],"b":3} => {"a":1,"b":3},{"a":2,"b":3}
// For Collection, ProjectSetOperator will do the following transform:
// [{"a":[1,2],"b":3},{"a":[1,2],"b":4}] = > [{"a":"1","b":3},{"a":"2","b":3},{"a":"1","b":4},{"a":"2","b":4}]
func (ps *ProjectSetOperator) Apply(ctx api.StreamContext, data interface{}, _ *xsql.FunctionValuer, _ *xsql.AggregateFunctionValuer) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (ps *ProjectSetOperator) handleSRFRowForCollection(ctx api.StreamContext, data xsql.Collection) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *ProjectSetOperator) handleSRFRow(ctx api.StreamContext, row xsql.Row) (*resultWrapper, error) {
	_ = "STUB: not implemented"
	// for now we only support 1 srf function in the field
	return nil, nil
}

// clear original column value

type resultWrapper struct {
	joinTuples  []*xsql.JoinTuple
	groupTuples []*xsql.GroupedTuples
	rows        []xsql.Row
}

func newResultWrapper(len int, row xsql.Row) *resultWrapper { _ = "STUB: not implemented"; return nil }

func (r *resultWrapper) appendTuple(index int, newRow xsql.Row) { _ = "STUB: not implemented"; return }
