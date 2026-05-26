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

package runtime

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

// Error handling: wrap all error in a function to handle

type PortableSource struct {
	symbolName string
	reg        *PluginMeta
	clean      func() error
	dataCh     DataInChannel

	topic string
	props map[string]any
}

type messageWrapper struct {
	Message map[string]any `json:"message"`
	Meta    map[string]any `json:"meta"`
}

func (ps *PortableSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PortableSource) Connect(ctx api.StreamContext, _ api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// wait for plugin data

// Control: send message to plugin to ask starting symbol

func (ps *PortableSource) Subscribe(ctx api.StreamContext, ingest api.TupleIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil

	// make sure recv has timeout
}

// do nothing

func NewPortableSource(symbolName string, reg *PluginMeta) *PortableSource {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PortableSource) Configure(topic string, props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PortableSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

var _ api.TupleSource = &PortableSource{}
