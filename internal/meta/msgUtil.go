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

package meta

import (
	"bytes"

	"gopkg.in/ini.v1"
)

var gUimsg map[string]*ini.File

func getMsg(language, section, key string) string { _ = "STUB: not implemented"; return "" }

func ReadUiMsgDir() error { _ = "STUB: not implemented"; return nil }

func ConstructJsonArray(jsonByteItems []fileContent) bytes.Buffer {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer)
}
