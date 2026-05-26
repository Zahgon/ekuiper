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

package context

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const (
	LoggerKey        = "$$logger"
	RuleStartKey     = "$$ruleStart"
	RuleWaitGroupKey = "$$ruleWaitGroup"
	TraceStrategyKey = "$$TraceStrategyKey"
)

const (
	AlwaysTraceStrategy = iota
	HeadTraceStrategy
)

type TraceStrategy int

type TraceStrategyWrapper struct {
	syncx.RWMutex
	Strategy TraceStrategy
}

type DefaultContext struct {
	ruleId         string
	opId           string
	instanceId     int
	runId          int
	ctx            context.Context
	err            error
	isTraceEnabled *atomic.Bool
	strategy       *TraceStrategyWrapper
	// Only initialized after withMeta set
	store    api.Store
	state    *sync.Map
	snapshot map[string]interface{}
	mu       syncx.Mutex
	// cache
	tpReg sync.Map
	jpReg sync.Map
}

func RuleBackground(ruleName string) *DefaultContext { _ = "STUB: not implemented"; return nil }

func Background() *DefaultContext { _ = "STUB: not implemented"; return nil }

func WithContext(ctx context.Context) *DefaultContext { _ = "STUB: not implemented"; return nil }

func WithValue(parent *DefaultContext, key, val interface{}) *DefaultContext {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultContext) PropagateTracer(par *DefaultContext) { _ = "STUB: not implemented"; return }

// Deadline Implement context interface
func (c *DefaultContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *DefaultContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) Err() error { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *DefaultContext) GetLogger() api.Logger { _ = "STUB: not implemented"; return *new(api.Logger) }

func (c *DefaultContext) GetRuleId() string { _ = "STUB: not implemented"; return "" }

func (c *DefaultContext) GetOpId() string { _ = "STUB: not implemented"; return "" }

func (c *DefaultContext) GetInstanceId() int { _ = "STUB: not implemented"; return 0 }

func (c *DefaultContext) GetRunId() int { _ = "STUB: not implemented"; return 0 }

func (c *DefaultContext) GetRootPath() string { _ = "STUB: not implemented"; return "" }

func (c *DefaultContext) SetError(err error) {
	_ = "STUB: not implemented"

	// ParseTemplate parse template string against data
	// The templates are built only once and cached in the context by its raw string as the key
	// If the prop string is not a template, a nil template is cached to indicate it has been parsed, and it will return the original string
	return
}

func (c *DefaultContext) ParseTemplate(prop string, data interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// not parsed before

// check if it is a template

func (c *DefaultContext) ParseJsonPath(prop string, data interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *DefaultContext) WithMeta(ruleId string, opId string, store api.Store) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithInstance(instanceId int) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithRuleId(ruleId string) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithOpId(opId string) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithRun(runId int) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *DefaultContext) WithCancel() (api.StreamContext, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(api.StreamContext), *new(context.CancelFunc)
}

func (c *DefaultContext) IncrCounter(key string, amount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultContext) GetCounter(key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *DefaultContext) GetAllState() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultContext) PutState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultContext) GetState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *DefaultContext) DeleteState(key string) error { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) Snapshot() error { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) SaveState(checkpointId int64) error { _ = "STUB: not implemented"; return nil }

func (c *DefaultContext) EnableTracer(enabled bool) { _ = "STUB: not implemented"; return }

func (c *DefaultContext) IsTraceEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *DefaultContext) GetStrategy() TraceStrategy {
	_ = "STUB: not implemented"
	return *new(TraceStrategy)
}

func (c *DefaultContext) SetStrategy(s TraceStrategy) { _ = "STUB: not implemented"; return }

func StringToStrategy(s string) TraceStrategy {
	_ = "STUB: not implemented"
	return *new(TraceStrategy)
}
