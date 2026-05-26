// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

package io

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/binder"
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
)

var ( // init once and read only
	sourceFactories      []binder.SourceFactory
	sourceFactoriesNames []string
	sinkFactories        []binder.SinkFactory
	sinkFactoriesNames   []string
)

func init() {
	f := binder.FactoryEntry{
		Name:    "built-in",
		Factory: GetManager(),
	}
	applyFactory(f)
}

func Initialize(factories []binder.FactoryEntry) error { _ = "STUB: not implemented"; return nil }

func applyFactory(f binder.FactoryEntry) { _ = "STUB: not implemented"; return }

func Source(name string) (api.Source, error) {
	_ = "STUB: not implemented"
	return *new(api.Source), nil
}

func GetSourcePlugin(name string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func Sink(name string) (api.Sink, error) { _ = "STUB: not implemented"; return *new(api.Sink), nil }

func GetSinkPlugin(name string) (plugin.EXTENSION_TYPE, string, string) {
	_ = "STUB: not implemented"
	return *new(plugin.EXTENSION_TYPE), "", ""
}

func LookupSource(name string) (api.Source, error) {
	_ = "STUB: not implemented"
	return *new(api.Source), nil
}
