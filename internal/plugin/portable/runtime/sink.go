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

type sinkConf struct {
	RequireAck bool `json:"requireAck"`
}

type PortableSink struct {
	symbolName string
	reg        *PluginMeta
	props      map[string]interface{}
	dataCh     DataOutChannel
	ackCh      DataInChannel
	c          *sinkConf
	clean      func() error
}

func (ps *PortableSink) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PortableSink) Connect(ctx api.StreamContext, _ api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Control: send message to plugin to ask starting symbol

// must start symbol firstly

func (ps *PortableSink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPortableSink(symbolName string, reg *PluginMeta) *PortableSink {
	_ = "STUB: not implemented"
	return nil
}

type ackResponse struct {
	Error string `json:"error"`
}

func (ps *PortableSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func recvAck(ctx api.StreamContext, dataCh DataInChannel) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// make sure recv has timeout
		nil
}
