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

package mock

import (
	"context"
	"time"

	filename "github.com/keepeye/logrus-filename"
	"github.com/sirupsen/logrus"

	"github.com/lf-edge/ekuiper/sdk/go/api"
)

type mockContext struct {
	Ctx    context.Context
	RuleId string
	OpId   string
}

// Implement context interface
func (c *mockContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *mockContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *mockContext) Err() error { _ = "STUB: not implemented"; return nil }

func (c *mockContext) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// Stream metas
func (c *mockContext) GetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *mockContext) GetLogger() api.Logger { _ = "STUB: not implemented"; return *new(api.Logger) }

func (c *mockContext) GetRuleId() string { _ = "STUB: not implemented"; return "" }

func (c *mockContext) GetOpId() string { _ = "STUB: not implemented"; return "" }

func (c *mockContext) GetInstanceId() int { _ = "STUB: not implemented"; return 0 }

func (c *mockContext) GetRootPath() string {
	_ = "STUB: not implemented"
	// loc, _ := conf.GetLoc("")
	return ""
}

func (c *mockContext) SetError(err error) { _ = "STUB: not implemented"; return }

func (c *mockContext) WithMeta(ruleId string, opId string) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *mockContext) WithInstance(_ int) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

func (c *mockContext) WithCancel() (api.StreamContext, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(api.StreamContext), *new(context.CancelFunc)
}

func (c *mockContext) IncrCounter(key string, amount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *mockContext) GetCounter(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *mockContext) PutState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *mockContext) GetState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *mockContext) DeleteState(key string) error { _ = "STUB: not implemented"; return nil }

func (c *mockContext) Snapshot() error { _ = "STUB: not implemented"; return nil }

func (c *mockContext) SaveState(checkpointId int64) error { _ = "STUB: not implemented"; return nil }

func newMockContext(ruleId string, opId string) api.StreamContext {
	_ = "STUB: not implemented"
	return *new(api.StreamContext)
}

type mockFuncContext struct {
	api.StreamContext
	funcId int
}

func (fc *mockFuncContext) GetFuncId() int { _ = "STUB: not implemented"; return 0 }

func newMockFuncContext(ctx api.StreamContext, id int) api.FunctionContext {
	_ = "STUB: not implemented"
	return *new(api.FunctionContext)
}

var Logger *logrus.Logger

func init() {
	l := logrus.New()
	filenameHook := filename.NewHook()
	filenameHook.Field = "file"
	l.AddHook(filenameHook)
	l.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   true,
		FullTimestamp:   true,
	})
	l.WithField("type", "main")
	Logger = l
}
