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

//go:build ui || !core

package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	components["meta"] = metaComp{}
}

var metaEndpoints []restEndpoint

type metaComp struct{}

func (m metaComp) register() {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (m metaComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

// list sink plugin
func sinksMetaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Get sink metadata when creating rules
func newSinkMetaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list functions
func functionsMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// list operators
func operatorsMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// list source plugin
func sourcesMetaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list shareMeta
func connectionsMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Get source metadata when creating stream
func sourceMetaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Get source metadata when creating stream
func connectionMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func resourceHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Get source yaml
func sourceConfHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Get share yaml
func connectionConfHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Get sink yaml
func sinkConfHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Add  del confkey
func sourceConfKeyHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Add  del confkey
func sinkConfKeyHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Add  del confkey
func connectionConfKeyHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// get updatable resources
func resourcesHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func getLanguage(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func sinkConnectionHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func sourceConnectionHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func lookupConnectionHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
