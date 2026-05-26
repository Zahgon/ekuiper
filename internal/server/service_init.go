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

//go:build service || !core

package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lf-edge/ekuiper/v2/internal/service"
)

var serviceManager *service.Manager

func init() {
	components["service"] = serviceComp{}
}

type serviceComp struct{}

func (s serviceComp) register() { _ = "STUB: not implemented"; return }

func (s serviceComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

func (s serviceComp) exporter() ConfManager { _ = "STUB: not implemented"; return *new(ConfManager) }

func servicesHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Problems decoding

func serviceHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Problems decoding

func serviceFunctionsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func serviceFunctionHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type serviceExporter struct{}

func (e serviceExporter) Import(ctx context.Context, services map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e serviceExporter) PartialImport(ctx context.Context, services map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e serviceExporter) Export() map[string]string { _ = "STUB: not implemented"; return nil }

func (e serviceExporter) Status() map[string]string { _ = "STUB: not implemented"; return nil }

func (e serviceExporter) Reset() { _ = "STUB: not implemented"; return }
