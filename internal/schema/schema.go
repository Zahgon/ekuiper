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

package schema

type Info struct {
	Type     string `json:"type" yaml:"type"`
	Name     string `json:"name" yaml:"name"`
	Content  string `json:"content,omitempty" yaml:"content,omitempty"`
	FilePath string `json:"file,omitempty" yaml:"filePath,omitempty"`
	SoPath   string `json:"soFile,omitempty" yaml:"soPath,omitempty"`
	Version  string `json:"version,omitempty" yaml:"version,omitempty"`
}

func (i *Info) InstallScript() string { _ = "STUB: not implemented"; return "" }

func (i *Info) Validate() error { _ = "STUB: not implemented"; return nil }
