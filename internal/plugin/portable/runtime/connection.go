// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.nanomsg.org/mangos/v3"

	// introduce ipc
	_ "go.nanomsg.org/mangos/v3/transport/ipc"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// TODO to design timeout strategy

// sockOptions Initialized in config
var (
	dialOptions = map[string]interface{}{
		mangos.OptionDialAsynch:       false,
		mangos.OptionMaxReconnectTime: 5 * time.Second,
		mangos.OptionReconnectTime:    100 * time.Millisecond,
	}
)

type Closable interface {
	Close() error
}

type ControlChannel interface {
	Handshake() error
	SendCmd(arg []byte) error
	Closable
}

// NanomsgReqChannel shared by symbols
type NanomsgReqChannel struct {
	syncx.Mutex
	sock mangos.Socket
}

func (r *NanomsgReqChannel) Close() error { _ = "STUB: not implemented"; return nil }

func (r *NanomsgReqChannel) SendCmd(arg []byte) error { _ = "STUB: not implemented"; return nil }

// resend if protocol state wrong, because of plugin restart or other problems

// Handshake should only be called once
func (r *NanomsgReqChannel) Handshake() error { _ = "STUB: not implemented"; return nil }

type DataInChannel interface {
	Recv() ([]byte, error)
	Closable
}

type DataOutChannel interface {
	Send([]byte) error
	Closable
}

type DataReqChannel interface {
	Req([]byte) ([]byte, error)
	Closable
}

type NanomsgReqRepChannel struct {
	syncx.Mutex
	sock mangos.Socket
}

func (r *NanomsgReqRepChannel) Close() error { _ = "STUB: not implemented"; return nil }

func (r *NanomsgReqRepChannel) Req(arg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resend if protocol state wrong, because of plugin restart or other problems

func CreateSourceChannel(ctx api.StreamContext) (DataInChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataInChannel), nil
}

func CreateFunctionChannel(symbolName string) (DataReqChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataReqChannel), nil
}

// Function must send out data quickly and wait for the response with some buffer

func CreateSinkChannel(ctx api.StreamContext) (DataOutChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataOutChannel), nil
}

func CreateSinkAckChannel(ctx api.StreamContext) (DataInChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataInChannel), nil
}

func CreateControlChannel(pluginName string) (ControlChannel, error) {
	_ = "STUB: not implemented"
	return *new(ControlChannel), nil
}

// NO time out now for control channel
// because the plugin instance liveness can be detected
// thus, if the plugin exit, the control channel will be closed

func setSockOptions(sock mangos.Socket, sockOptions map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func listenWithRetry(sock mangos.Socket, url string) error { _ = "STUB: not implemented"; return nil }
