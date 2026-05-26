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

package plugin

type PluginType int

const (
	SOURCE PluginType = iota
	SINK
	FUNCTION
	PORTABLE
)

var PluginTypes = []string{"sources", "sinks", "functions", "portable"}

var PluginTypeMap = map[string]PluginType{
	"sources":   SOURCE,
	"sinks":     SINK,
	"functions": FUNCTION,
	"portable":  PORTABLE,
}

type Plugin interface {
	GetName() string
	GetFile() string
	GetShellParas() []string
	GetSymbols() []string
	SetName(n string)
	GetInstallScripts() []byte
}

// IOPlugin Unify model. Flat all properties for each kind.
type IOPlugin struct {
	Name       string   `json:"name" yaml:"name"`
	File       string   `json:"file" yaml:"file"`
	ShellParas []string `json:"shellParas,omitempty" yaml:"shellParas,omitempty"`
}

func (p *IOPlugin) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *IOPlugin) GetFile() string { _ = "STUB: not implemented"; return "" }

func (p *IOPlugin) GetShellParas() []string { _ = "STUB: not implemented"; return nil }

func (p *IOPlugin) GetSymbols() []string { _ = "STUB: not implemented"; return nil }

func (p *IOPlugin) SetName(n string) { _ = "STUB: not implemented"; return }

func (p *IOPlugin) GetInstallScripts() []byte { _ = "STUB: not implemented"; return nil }

func NewPluginByType(t PluginType) Plugin { _ = "STUB: not implemented"; return *new(Plugin) }

type FuncPlugin struct {
	IOPlugin
	// Optional, if not specified, a default element with the same name of the file will be registered
	Functions []string `json:"functions"`
}

func (fp *FuncPlugin) GetSymbols() []string { _ = "STUB: not implemented"; return nil }

type EXTENSION_TYPE int

const (
	NONE_EXTENSION EXTENSION_TYPE = iota
	INTERNAL
	NATIVE_EXTENSION
	PORTABLE_EXTENSION
	SERVICE_EXTENSION
	JS_EXTENSION
)

var ExtensionTypes = []string{
	"none", "internal", "native", "portable", "service", "wasm", "js",
}
