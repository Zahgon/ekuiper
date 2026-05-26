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
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	components["schema"] = schemaComp{}
}

type schemaComp struct{}

func (sc schemaComp) register() { _ = "STUB: not implemented"; return }

func (sc schemaComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

func (sc schemaComp) exporter() ConfManager { _ = "STUB: not implemented"; return *new(ConfManager) }

func schemasHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func schemaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type schemaExporter struct{}

func (e schemaExporter) Import(ctx context.Context, s map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e schemaExporter) PartialImport(ctx context.Context, s map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e schemaExporter) Export() map[string]string { _ = "STUB: not implemented"; return nil }

func (e schemaExporter) Status() map[string]string { _ = "STUB: not implemented"; return nil }

func (e schemaExporter) Reset() { _ = "STUB: not implemented"; return }

func (e schemaExporter) InstallScript(s string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}
