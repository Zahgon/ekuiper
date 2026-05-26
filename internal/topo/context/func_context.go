// Copyright 2021 EMQ Technologies Co., Ltd.
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

package context

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type DefaultFuncContext struct {
	api.StreamContext
	funcId int
}

func NewDefaultFuncContext(ctx api.StreamContext, id int) *DefaultFuncContext {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultFuncContext) IncrCounter(key string, amount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultFuncContext) GetCounter(key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *DefaultFuncContext) PutState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultFuncContext) GetState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *DefaultFuncContext) DeleteState(key string) error { _ = "STUB: not implemented"; return nil }

func (c *DefaultFuncContext) GetFuncId() int { _ = "STUB: not implemented"; return 0 }

func (c *DefaultFuncContext) convertKey(key string) string { _ = "STUB: not implemented"; return "" }
