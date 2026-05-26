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
	context2 "context"
	"sync"

	"github.com/lf-edge/ekuiper/sdk/go/api"
	"github.com/lf-edge/ekuiper/sdk/go/connection"
)

type funcRuntime struct {
	s      api.Function
	ch     connection.DataInOutChannel
	ctx    context2.Context
	cancel context2.CancelFunc
	key    string
}

func setupFuncRuntime(con *Control, s api.Function) (*funcRuntime, error) {
	_ = "STUB: not implemented"
	// connect to mq server
	return nil, nil
}

// TODO how to stop? Nearly never end because each function only have one instance
func (s *funcRuntime) run() { _ = "STUB: not implemented"; return }

// TODO multiple error
func (s *funcRuntime) stop() error { _ = "STUB: not implemented"; return nil }

func (s *funcRuntime) isRunning() bool { _ = "STUB: not implemented"; return false }

func encodeReply(state bool, arg interface{}) []byte { _ = "STUB: not implemented"; return nil }

func parseFuncContextArgs(args []interface{}) ([]interface{}, api.FunctionContext, error) {
	_ = "STUB: not implemented"
	return nil, *new(api.FunctionContext), nil
}

var exeFuncCtxMap = &sync.Map{}
