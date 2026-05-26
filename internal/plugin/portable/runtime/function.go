// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

// PortableFunc each function symbol only has a singleton
// Each singleton are long-running go routine
// Currently, it is cached and never ended once created
// It is actually a wrapper of the data channel and can be fit to any plugin instance
// Thus, it is possible to hot reload, which is simply attach a new nng client to the same channel
// without changing the server(plugin runtime) side
// TODO think about ending a portable func when needed.
type PortableFunc struct {
	symbolName string
	reg        *PluginMeta // initial plugin meta, only used for initialize the function instance
	dataCh     DataReqChannel
	isAgg      int // 0 - not calculate yet, 1 - no, 2 - yes
}

func NewPortableFunc(symbolName string, reg *PluginMeta) (_ *PortableFunc, e error) {
	_ = "STUB: not implemented"
	// Setup channel and route the data
	return nil, nil
}

// Create function channel

// Start symbol

func (f *PortableFunc) Validate(args []interface{}) error {
	_ = "STUB: not implemented"
	// TODO function arg encoding
	return nil
}

func (f *PortableFunc) Exec(ctx api.FunctionContext, args []any) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func handleTimeout(err error, pname string) error { _ = "STUB: not implemented"; return nil }

func (f *PortableFunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (f *PortableFunc) Close() error { _ = "STUB: not implemented"; return nil }

// Symbol must be closed by instance manager
//		ins.StopSymbol(ctx, c)

func encode(funcName string, arg interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeCtx(ctx api.FunctionContext) (string, error) { _ = "STUB: not implemented"; return "", nil }
