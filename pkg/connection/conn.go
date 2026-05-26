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

package connection

import (
	"sync"
	"sync/atomic"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type ConnWrapper struct {
	ID          string
	initialized bool
	conn        modules.Connection
	err         error
	l           syncx.RWMutex
	readCh      chan struct{}
	detachCh    chan struct{}
}

func (cw *ConnWrapper) setConn(conn modules.Connection, err error) {
	_ = "STUB: not implemented"
	return
}

// Wait will wait for connection connected or the caller interrupts (like rule exit)
func (cw *ConnWrapper) Wait(connectorCtx api.StreamContext) (modules.Connection, error) {
	_ = "STUB: not implemented"
	return *new(modules.Connection), nil
}

func (cw *ConnWrapper) IsInitialized() bool { _ = "STUB: not implemented"; return false }

func newConnWrapper(ctx api.StreamContext, meta *Meta) *ConnWrapper {
	_ = "STUB: not implemented"
	return nil
}

type Meta struct {
	ID    string         `json:"id"`
	Typ   string         `json:"typ"`
	Props map[string]any `json:"props"`
	// named means connection is created manually
	Named bool `json:"named"`

	refCount atomic.Int32 `json:"-"`
	ref      sync.Map     `json:"-"`
	cw       *ConnWrapper `json:"-"`
	// The first connection status
	// If connection is stateful, the status will update all the way
	// For stateless connection, the status needs to ping
	status    atomic.Value `json:"-"`
	lastError atomic.Value `json:"-"`
}

func (meta *Meta) NotifyStatus(status string, s string) { _ = "STUB: not implemented"; return }

func (meta *Meta) AddRef(refId string, sc api.StatusChangeHandler) {
	_ = "STUB: not implemented"
	return
}

func (meta *Meta) DeRef(refId string) { _ = "STUB: not implemented"; return }

func (meta *Meta) GetRefCount() int { _ = "STUB: not implemented"; return 0 }

func (meta *Meta) GetRefNames() (result []string) { _ = "STUB: not implemented"; return nil }

func (meta *Meta) GetStatus() (s string, e string) { _ = "STUB: not implemented"; return "", "" }

// if connected, cw, cw.conn should exist
