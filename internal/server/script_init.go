// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

//go:build script || full

package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	components["script"] = scriptComp{}
}

type scriptComp struct{}

func (p scriptComp) register() { _ = "STUB: not implemented"; return }

func (p scriptComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

func (p scriptComp) exporter() ConfManager { _ = "STUB: not implemented"; return *new(ConfManager) }

func jsfuncsHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Problems decoding

func jsfuncHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Problems decoding
