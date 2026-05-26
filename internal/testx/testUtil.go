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

package testx

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// errstring returns the string representation of an error.
func Errstring(err error) string { _ = "STUB: not implemented"; return "" }

func InitEnv(id string) { _ = "STUB: not implemented"; return }

func InitBroker(id string) (string, func(), error) {
	_ = "STUB: not implemented"
	// Create the new MQTT Server.
	return "", nil, nil
}

// Allow all connections.

// Create a TCP listener on a standard port.

// wait server close

func TestHttp(client *http.Client, url string, method string) error {
	_ = "STUB: not implemented"
	return nil
}

var body = []byte(`{
        "title": "Post title",
        "body": "Post description",
        "userId": 1
    }`)

func CreateWebsocketClient(ip string, port int, endpoint string) (*websocket.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MockTuple struct {
	Map      map[string]any
	Template map[string]string
}

func (m MockTuple) DynamicProps(template string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (m MockTuple) AllProps() map[string]string { _ = "STUB: not implemented"; return nil }

func (m MockTuple) Value(key, table string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (m MockTuple) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

type MockRawTuple struct {
	Content  []byte
	Template map[string]string
}

func (m *MockRawTuple) DynamicProps(template string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (m *MockRawTuple) AllProps() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *MockRawTuple) Raw() []byte { _ = "STUB: not implemented"; return nil }

func (m *MockRawTuple) Replace(newContent []byte) { _ = "STUB: not implemented"; return }
