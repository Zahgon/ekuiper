// Copyright 2024 EMQ Technologies Co., Ltd.
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

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
	"github.com/lf-edge/ekuiper/v2/internal/schema"
	"github.com/lf-edge/ekuiper/v2/internal/service"
)

type MetaConfiguration struct {
	SourceConfig     map[string]map[string]any `json:"sourceConfig,omitempty" yaml:"sourceConfig,omitempty"`
	SinkConfig       map[string]map[string]any `json:"sinkConfig,omitempty" yaml:"sinkConfig,omitempty"`
	ConnectionConfig map[string]map[string]any `json:"connectionConfig,omitempty" yaml:"connectionConfig,omitempty"`
	// plugins
	NativePlugins   map[string]*plugin.IOPlugin `json:"nativePlugins,omitempty" yaml:"nativePlugins,omitempty"`
	PortablePlugins map[string]*plugin.IOPlugin `json:"portablePlugins,omitempty" yaml:"portablePlugins,omitempty"`
	// others
	Service map[string]*service.ServiceCreationRequest `json:"service,omitempty" yaml:"service,omitempty"`
	Schema  map[string]*schema.Info                    `json:"schema,omitempty" yaml:"schema,omitempty"`
	Uploads map[string]*fileContent                    `json:"uploads,omitempty" yaml:"uploads,omitempty"`
	// rules related
	Streams map[string]*DatasourceExport `json:"streams" yaml:"streams"`
	Tables  map[string]*DatasourceExport `json:"tables,omitempty" yaml:"tables,omitempty"`
	Rules   map[string]*def.Rule         `json:"rules" yaml:"rules"`
}

type DatasourceExport struct {
	SQL string `json:"sql" yaml:"sql"`
}

func GenMetaConfiguration() (*MetaConfiguration, error) { _ = "STUB: not implemented"; return nil, nil }

func addConfiguration(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func addPlugins(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func addSchema(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func addUploads(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func addService(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func yamlConfigurationExportHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type importConfiguration struct {
	configurationInfo

	// TODO: support these later
	Partial bool `json:"partial" yaml:"partial"`
	Reboot  bool `json:"reboot" yaml:"reboot"`
}

func yamlConfImportHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func importFromByte(content []byte) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

const (
	mockErrStart int = iota
	mockImportFromByteErr
	mockSourcesErr
	mockSinksErr
	mockConnectionsErr
	mockStreamsErr
	mockTablesErr
	mockRulesErr
	mockUploadErr
	mockServiceErr
	mockSchemaErr
	mockPortablePluginErr
	mockNativePluginErr
	mockErrEnd
)

func mockImportErr(err error, errSwitch int) error { _ = "STUB: not implemented"; return nil }

func importYamlConf(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

func importSchema(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func importService(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func importUploads(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func importByManager(want map[string]string, m ConfManager, typ string) error {
	_ = "STUB: not implemented"
	return nil
}

func importPortablePlugins(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func importNativePlugins(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

func importRules(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

func importConfigurations(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

func importDataSource(m *MetaConfiguration) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

//nolint:staticcheck

func writeConf(key string, value map[string]any) error { _ = "STUB: not implemented"; return nil }

func splitConfKey(key string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func replaceConfigurations(key string, props map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
