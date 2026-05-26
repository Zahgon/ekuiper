// Copyright 2023 EMQ Technologies Co., Ltd.
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

//go:build msgpack

package service

import (
	"fmt"
	"net/rpc"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

func init() {
	executors[MSGPACK] = func(desc descriptor, opt *interfaceOpt, _ *interfaceInfo) (executor, error) {
		d, ok := desc.(interfaceDescriptor)
		if !ok {
			return nil, fmt.Errorf("invalid descriptor type for msgpack-rpc")
		}
		exe := &msgpackExecutor{
			descriptor:   d,
			interfaceOpt: opt,
		}
		return exe, nil
	}
}

type msgpackExecutor struct {
	descriptor interfaceDescriptor
	*interfaceOpt

	syncx.Mutex
	connected bool
	conn      *rpc.Client
}

// InvokeFunction flat the params and result
func (m *msgpackExecutor) InvokeFunction(_ api.FunctionContext, name string, params []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO argument flat

// do nothing
