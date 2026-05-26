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

package node

import (
	"bytes"
	"text/template"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
)

// TransformOp transforms the row/collection to sink tuples
// Immutable: false
// Change trigger frequency: true, by sendSingle property
// Input: Row/Collection
// Output: MessageTuple, SinkTupleList, RawTuple
type TransformOp struct {
	*defaultSinkNode
	dataField     string
	fields        []string
	excludeFields []string
	sendSingle    bool
	omitIfEmpty   bool
	// If the result format is text, the dataTemplate should be used to format the data and skip the encode step. Otherwise, the text must be unmarshall back to map
	isTextFormat bool
	dt           *template.Template
	templates    map[string]*template.Template
	isSliceMode  bool
	// temp state
	output bytes.Buffer
}

// NewTransformOp creates a transform node
// sink conf should have been validated before
func NewTransformOp(name string, rOpt *def.RuleOption, sc *SinkConf, templates []string) (*TransformOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TransformOp) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

// Worker do not need to process error and control messages
func (t *TransformOp) Worker(ctx api.StreamContext, item any) []any {
	_ = "STUB: not implemented"
	return nil
}

// MessageTuple or SinkTupleList

func (t *TransformOp) transformSlice(ctx api.StreamContext, item any) []any {
	_ = "STUB: not implemented"
	return nil
}

// TODO keep the tuple meta etc.
func toSinkTuple(_, spanCtx api.StreamContext, bs any, props map[string]string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// doTransform transforms the data according to the dataTemplate and fields
// If the dataTemplate is the last action and the result is text, the data will be returned as []byte
// Otherwise, the data will be return as a map or []map
func (t *TransformOp) doTransform(d any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// if only do data template

func (t *TransformOp) calculateProps(data any) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func itemToMap(item interface{}) []map[string]any { _ = "STUB: not implemented"; return nil }

// The order is important here, because some element is both a collection and a row, such as WindowTuples, JoinTuples, etc.
