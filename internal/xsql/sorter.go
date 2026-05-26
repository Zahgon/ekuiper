// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

package xsql

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
)

// MultiSorter implements the Sort interface, sorting the changes within.
type MultiSorter struct {
	Ctx api.StreamContext
	SortingData
	fields    ast.SortFields
	valuer    *FunctionValuer
	aggValuer *AggregateFunctionValuer
	values    []map[string]interface{}
}

func (ms *MultiSorter) GetTracerCtx() api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (ms *MultiSorter) SetTracerCtx(ctx api.StreamContext) {
	_ = "STUB: not implemented"

	// OrderedBy returns a Sorter that sorts using the less functions, in order.
	// Call its Sort method to sort the data.
	return
}

func OrderedBy(fields ast.SortFields, fv *FunctionValuer, afv *AggregateFunctionValuer) *MultiSorter {
	_ = "STUB: not implemented"
	return nil
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could change the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (ms *MultiSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ms *MultiSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort sorts the argument slice according to the fewer functions passed to OrderedBy.
func (ms *MultiSorter) Sort(data SortingData) error { _ = "STUB: not implemented"; return nil }

func validate(t string, v interface{}) error { _ = "STUB: not implemented"; return nil }
