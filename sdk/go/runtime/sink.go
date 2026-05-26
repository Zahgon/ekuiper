// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

type sinkRuntime struct {
	s      api.Sink
	ch     connection.DataInChannel
	ackCh  connection.DataOutChannel
	ctx    api.StreamContext
	cancel context2.CancelFunc
	key    string
}

func setupSinkRuntime(con *Control, s api.Sink) (*sinkRuntime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sinkRuntime) run() { _ = "STUB: not implemented"; return }

// blocking read, must be interrupted when stopping
// Will be stopped by closing the socket in stop

// do nothing

// temporary remove golang plugin sink ack
//r := &ackResponse{}
//if err != nil {
//	r.Error = err.Error()
//}
//data, _ := json.Marshal(r)
//if err := s.ackCh.Send(data); err != nil {
//	s.ctx.GetLogger().Errorf("ack error: %s", err.Error())
//	_ = s.stop()
//	return
//}

type ackResponse struct {
	Error string `json:"error"`
}

func (s *sinkRuntime) stop() error { _ = "STUB: not implemented"; return nil }

func (s *sinkRuntime) isRunning() bool { _ = "STUB: not implemented"; return false }
