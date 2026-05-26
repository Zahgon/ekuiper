// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

// UnOperation interface represents unary operations (i.e. Map, Filter, etc)
type UnOperation interface {
	Apply(ctx api.StreamContext, data interface{}, fv *xsql.FunctionValuer, afv *xsql.AggregateFunctionValuer) interface{}
}

// UnFunc implements UnOperation as type func (context.Context, interface{})
type UnFunc func(api.StreamContext, interface{}) interface{}

// Apply implements UnOperation.Apply method
func (f UnFunc) Apply(ctx api.StreamContext, data interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

type UnaryOperator struct {
	*defaultSinkNode
	op        UnOperation
	cancelled bool
}

// New NewUnary creates *UnaryOperator value
func New(name string, options *def.RuleOption) *UnaryOperator {
	_ = "STUB: not implemented"
	return nil
}

// SetOperation sets the executor operation
func (o *UnaryOperator) SetOperation(op UnOperation) {
	_ = "STUB: not implemented"

	// Exec is the entry point for the executor
	return
}

func (o *UnaryOperator) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// validate p

func (o *UnaryOperator) doOp(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// process incoming item

// ends, do nothing

// is cancelling
