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

package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/lf-edge/ekuiper/v2/pkg/model"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type GlobalServerManager struct {
	syncx.RWMutex
	instanceID        int
	endpoint          map[string]string
	server            *http.Server
	router            *mux.Router
	routes            map[string]http.HandlerFunc
	upgrader          websocket.Upgrader
	websocketEndpoint map[string]*websocketEndpointContext
	sseEndpoint       map[string]*sseEndpointContext
}

var (
	manager     *GlobalServerManager
	managerLock syncx.RWMutex
)

func InitGlobalServerManager(ip string, port int, tlsConf *model.TlsConf) {
	_ = "STUB: not implemented"
	return
}

// Good practice to set timeouts to avoid Slowloris attacks.

// always allowed any origin

func ShutDown() { _ = "STUB: not implemented"; return }

func RegisterEndpoint(endpoint string, method string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func UnregisterEndpoint(endpoint, method string) { _ = "STUB: not implemented"; return }

const (
	TopicPrefix = "$$httppush/"
)

func (m *GlobalServerManager) RegisterEndpoint(endpoint string, method string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *GlobalServerManager) UnregisterEndpoint(endpoint, method string) {
	_ = "STUB: not implemented"
	return
}

func (m *GlobalServerManager) Shutdown() { _ = "STUB: not implemented"; return }

func handleError(w http.ResponseWriter, err error, prefix string) {
	_ = "STUB: not implemented"
	return
}

func buildKey(ep, method string) string { _ = "STUB: not implemented"; return "" }
