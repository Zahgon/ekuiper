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

//go:build portable || !core

package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lf-edge/ekuiper/v2/internal/plugin/portable"
)

var portableManager *portable.Manager

func init() {
	components["portable"] = portableComp{}
}

type portableComp struct{}

func (p portableComp) register() { _ = "STUB: not implemented"; return }

func (p portableComp) rest(r *mux.Router) { _ = "STUB: not implemented"; return }

func (p portableComp) exporter() ConfManager { _ = "STUB: not implemented"; return *new(ConfManager) }

func portablesHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Problems decoding

func portableStatusHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func portableHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Problems decoding

type portableExporter struct{}

func (e portableExporter) Import(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e portableExporter) PartialImport(ctx context.Context, plugins map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (e portableExporter) Export() map[string]string { _ = "STUB: not implemented"; return nil }

func (e portableExporter) Status() map[string]string { _ = "STUB: not implemented"; return nil }

func (e portableExporter) Reset() { _ = "STUB: not implemented"; return }
