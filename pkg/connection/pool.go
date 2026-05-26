// Copyright 2024 EMQ Technologies Co., Ltd.
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
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// Connection pool manages all connections in the system. There are two kinds of connections:
// 1. Named connection: Long running connection. Users can create it standalone through dedicated API without rules.
// The connection will run through all the eKuiper server lifecycle. When restarting, it will be loaded and run as server init.
// 2. Anonymous connection: It is a subsidiary of rules. The rule source/sink defines connection and the connection will
// be fetched when rules start. If no rule has accessed it, it will be closed and dropped.

type Manager struct {
	syncx.RWMutex
	// key is selId(explicitly specified or anonymous)
	connectionPool map[string]*Meta
}

var (
	globalConnectionManager *Manager
	mockErr                 = true
)

func init() {
	globalConnectionManager = &Manager{
		connectionPool: make(map[string]*Meta),
	}
}

func InitConnectionManager4Test() error { _ = "STUB: not implemented"; return nil }

func InitConnectionManager(ctx context.Context) { _ = "STUB: not implemented"; return }

const (
	DefaultInitialInterval = 100 * time.Millisecond
	DefaultMaxInterval     = 10 * time.Second
)

func PatrolConnectionStatusJob(ctx context.Context) { _ = "STUB: not implemented"; return }

func patrolConnectionStatus() { _ = "STUB: not implemented"; return }

// For now, we only patrol named connection

func NewExponentialBackOff() *backoff.ExponentialBackOff { _ = "STUB: not implemented"; return nil }

// FetchConnection is called by source/sink to get or create an anonymous connection instance in the pool
func FetchConnection(ctx api.StreamContext, refId, typ string, props map[string]interface{}, sc api.StatusChangeHandler) (*ConnWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReloadNamedConnection is called when server starts. It initializes all stored named connections
func ReloadNamedConnection() error { _ = "STUB: not implemented"; return nil }

// Connection API handlers

func CreateNamedConnection(ctx api.StreamContext, id, typ string, props map[string]any) (*ConnWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createNamedConnection(ctx api.StreamContext, id, typ string, props map[string]any) (*ConnWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAllConnectionsMeta(forceAll bool) []*Meta { _ = "STUB: not implemented"; return nil }

func GetConnectionDetail(_ api.StreamContext, id string) (*Meta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DropNameConnection(ctx api.StreamContext, selId string) error {
	_ = "STUB: not implemented"
	return nil
}

func dropNameConnection(ctx api.StreamContext, selId string) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateConnection(ctx api.StreamContext, id, typ string, props map[string]any) (*ConnWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isInternalConnection(id string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func DetachConnection(ctx api.StreamContext, conId string) error {
	_ = "STUB: not implemented"
	return nil
}

func getConnectionRef(id string) int { _ = "STUB: not implemented"; return 0 }

func storeConnectionMeta(plugin, id string, props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func dropConnectionStore(plugin, id string) error { _ = "STUB: not implemented"; return nil }

func attachConnection(conId string, refId string, sc api.StatusChangeHandler) (*ConnWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func detachConnection(ctx api.StreamContext, conId string) error {
	_ = "STUB: not implemented"
	return nil
}

func createConnection(connCtx api.StreamContext, meta *Meta) (modules.Connection, error) {
	_ = "STUB: not implemented"
	return *new(modules.Connection), nil
}

// Return the unique connection id and whether it is set explicitly
func extractSelID(props map[string]interface{}, anomId string) string {
	_ = "STUB: not implemented"
	return ""
}

func extractRefId(ctx api.StreamContext) string { _ = "STUB: not implemented"; return "" }
