// Copyright 2021-2023 EMQ Technologies Co., Ltd.
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
	"context"
	"time"

	"github.com/lf-edge/ekuiper/sdk/go/api"
)

const LoggerKey = "$$logger"

type DefaultContext struct {
	ruleId     string
	opId       string
	instanceId int
	ctx        context.Context
	// Only initialized after withMeta set
	logger api.Logger
}

func Background() *DefaultContext { _ = "STUB: not implemented"; return nil }

func WithValue(parent *DefaultContext, key, val interface{}) *DefaultContext {
	_ = "STUB: not implemented"
	return nil
}

// Deadline Implement context interface
func (c *DefaultContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *DefaultContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) Err() error { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// Stream metas
func (c *DefaultContext) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *DefaultContext) GetLogger() api.Logger { _ = "STUB: not implemented"; return *new(api.Logger) }

func (c *DefaultContext) GetRuleId() string { _ = "STUB: not implemented"; return "" }

func (c *DefaultContext) GetOpId() string { _ = "STUB: not implemented"; return "" }

func (c *DefaultContext) GetInstanceId() int { _ = "STUB: not implemented"; return 0 }

func (c *DefaultContext) WithMeta(ruleId string, opId string) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithInstance(instanceId int) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithCancel() (api.StreamContext, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(api.StreamContext), *new(context.CancelFunc)
}
