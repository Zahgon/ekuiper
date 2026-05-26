// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

package nng

import (
	"sync/atomic"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.nanomsg.org/mangos/v3"
	_ "go.nanomsg.org/mangos/v3/transport/ipc"
	_ "go.nanomsg.org/mangos/v3/transport/tcp"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type SockConf struct {
	Url      string `json:"url"`
	Protocol string `json:"protocol"`
}

type Sock struct {
	mangos.Socket
	url string
	id  string
	// a ont-time signal to inform the sock is connected at the first time. later connection status are handled below
	ready chan struct{}
	// connection status through lifecycle
	scHandler api.StatusChangeHandler
	connected atomic.Bool
	status    atomic.Value
}

func (s *Sock) SetStatusChangeHandler(ctx api.StreamContext, handler api.StatusChangeHandler) {
	_ = "STUB: not implemented"
	return
}

func (s *Sock) Status(_ api.StreamContext) modules.ConnectionStatus {
	_ = "STUB: not implemented"
	return *new(modules.ConnectionStatus)
}

func (s *Sock) GetId(_ api.StreamContext) string { _ = "STUB: not implemented"; return "" }

func (s *Sock) Provision(ctx api.StreamContext, conId string, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// options consider to export

func (s *Sock) Dial(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	// sock.SetOption(mangos.OptionWriteQLen, 100)
	// sock.SetOption(mangos.OptionReadQLen, 100)
	// sock.SetOption(mangos.OptionBestEffort, false)
	return nil
}

// will not report error and keep connecting

// make it block until first connected

var nngTimeout = 5 * time.Second

func CreateConnection(_ api.StreamContext) modules.Connection {
	_ = "STUB: not implemented"
	return *new(modules.Connection)
}

func ValidateConf(props map[string]any) (*SockConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse the URL

func (s *Sock) Ping(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *Sock) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *Sock) Send(ctx api.StreamContext, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

var _ modules.StatefulDialer = &Sock{}
