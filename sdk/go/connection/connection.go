// Copyright 2022 EMQ Technologies Co., Ltd.
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

package connection

import (
	"time"

	"go.nanomsg.org/mangos/v3"
	// introduce ipc
	_ "go.nanomsg.org/mangos/v3/transport/ipc"

	"github.com/lf-edge/ekuiper/sdk/go/api"
)

// Options Initialized in plugin.go Start according to the config
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

type ReplyFunc func([]byte) []byte

type ControlChannel interface {
	// reply with string message
	Run(ReplyFunc) error
	Closable
}

type DataInChannel interface {
	Recv() ([]byte, error)
	Closable
}

type DataOutChannel interface {
	Send([]byte) error
	Closable
}

type DataInOutChannel interface {
	Run(ReplyFunc) error
	Closable
}

type NanomsgRepChannel struct {
	sock mangos.Socket
}

// Run until process end
func (r *NanomsgRepChannel) Run(f ReplyFunc) error { _ = "STUB: not implemented"; return nil }

// After timeout or protocol state error, REQ socket needs to send before recv.
// Re-send handshake to reset protocol state before trying to recv again.

// Successfully received message

func (r *NanomsgRepChannel) Close() error { _ = "STUB: not implemented"; return nil }

func CreateControlChannel(pluginName string) (ControlChannel, error) {
	_ = "STUB: not implemented"
	return *new(ControlChannel), nil
}

func CreateSourceChannel(ctx api.StreamContext) (DataOutChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataOutChannel), nil
}

func CreateFuncChannel(symbolName string) (DataInOutChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataInOutChannel), nil
}

// Add recv timeout to prevent indefinite blocking during idle periods

func CreateSinkChannel(ctx api.StreamContext) (DataInChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataInChannel), nil
}

func CreateSinkAckChannel(ctx api.StreamContext) (DataOutChannel, error) {
	_ = "STUB: not implemented"
	return *new(DataOutChannel), nil
}

func setSockOptions(sock mangos.Socket, sockOptions map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func listenWithRetry(sock mangos.Socket, url string) error { _ = "STUB: not implemented"; return nil }
