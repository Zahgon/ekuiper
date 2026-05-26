// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

	"github.com/lf-edge/ekuiper/v2/internal/processor"
)

type ConfManager interface {
	Import(context.Context, map[string]string) map[string]string
	PartialImport(context.Context, map[string]string) map[string]string
	Export() map[string]string
	Status() map[string]string
	Reset()
}

var managers = map[string]ConfManager{}

func InitConfManagers() { _ = "STUB: not implemented"; return }

const ProcessErr = "process error"

type Configuration struct {
	Streams          map[string]string `json:"streams"`
	Tables           map[string]string `json:"tables"`
	Rules            map[string]string `json:"rules"`
	NativePlugins    map[string]string `json:"nativePlugins"`
	PortablePlugins  map[string]string `json:"portablePlugins"`
	SourceConfig     map[string]string `json:"sourceConfig"`
	SinkConfig       map[string]string `json:"sinkConfig"`
	ConnectionConfig map[string]string `json:"connectionConfig"`
	Service          map[string]string `json:"Service"`
	Schema           map[string]string `json:"Schema"`
	Uploads          map[string]string `json:"uploads"`
	Scripts          map[string]string `json:"scripts"`
}

func configurationExport() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func configurationExportHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func configurationReset() { _ = "STUB: not implemented"; return }

type ImportConfigurationStatus struct {
	ErrorMsg       string
	ConfigResponse Configuration
}

func configurationImport(ctx context.Context, data []byte, reboot bool) ImportConfigurationStatus {
	_ = "STUB: not implemented"
	return *new(ImportConfigurationStatus)
}

func configurationPartialImport(ctx context.Context, data []byte) ImportConfigurationStatus {
	_ = "STUB: not implemented"
	return *new(ImportConfigurationStatus)
}

type configurationInfo struct {
	Content  string `json:"content" yaml:"content"`
	FilePath string `json:"file" yaml:"filePath"`
}

func configurationImportHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func handleConfigurationImport(ctx context.Context, rsi *configurationInfo, partial bool, stop bool) (*ImportConfigurationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func configurationStatusExport() Configuration {
	_ = "STUB: not implemented"
	return *new(Configuration)
}

func configurationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func configurationStatusHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func importRuleSetPartial(all processor.Ruleset) processor.Ruleset {
	_ = "STUB: not implemented"
	return *new(processor.Ruleset)
}

// replace streams

// replace tables

func uploadsReset() { _ = "STUB: not implemented"; return }

func uploadsExport() map[string]string { _ = "STUB: not implemented"; return nil }

func uploadsStatusExport() map[string]string { _ = "STUB: not implemented"; return nil }

func uploadsImport(s map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }
