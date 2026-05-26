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

package runtime

import (
	context2 "context"

	"github.com/lf-edge/ekuiper/sdk/go/api"
	"github.com/lf-edge/ekuiper/sdk/go/connection"
)

// lifecycle controlled by plugin
// if stop by error, inform plugin

type sourceRuntime struct {
	s      api.Source
	ch     connection.DataOutChannel
	ctx    api.StreamContext
	cancel context2.CancelFunc
	key    string
}

func setupSourceRuntime(con *Control, s api.Source) (*sourceRuntime, error) {
	_ = "STUB: not implemented"
	// init context with args
	return nil, nil
}

// TODO check cmd error handling or using health check

// init config with args and call source config

// connect to mq server

func (s *sourceRuntime) run() { _ = "STUB: not implemented"; return }

func (s *sourceRuntime) stop() error { _ = "STUB: not implemented"; return nil }

func (s *sourceRuntime) isRunning() bool { _ = "STUB: not implemented"; return false }
