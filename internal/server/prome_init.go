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

package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	p := &promeComp{}
	servers["prometheus"] = p
	components["prometheus"] = p
}

type promeComp struct {
	s *http.Server
}

func (p *promeComp) register() {
	_ = "STUB: not implemented"
	// Do nothing
	return
}

func (p *promeComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

func (p *promeComp) serve() { _ = "STUB: not implemented"; return }

// Start prometheus service

func (p *promeComp) close() { _ = "STUB: not implemented"; return }
