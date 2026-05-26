// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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

const (
	sink   = `sink`
	source = `source`
)

type (
	author struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Company string `json:"company"`
		Website string `json:"website"`
	}
	fileLanguage struct {
		English string `json:"en_US"`
		Chinese string `json:"zh_CN"`
	}
	fileField struct {
		Name              string        `json:"name"`
		Default           interface{}   `json:"default"`
		Control           string        `json:"control"`
		ConnectionRelated bool          `json:"connection_related"`
		Optional          bool          `json:"optional"`
		Type              string        `json:"type"`
		Hint              *fileLanguage `json:"hint"`
		Label             *fileLanguage `json:"label"`
		Values            interface{}   `json:"values"`
	}
	fileAbout struct {
		Trial       bool          `json:"trial"`
		Installed   bool          `json:"installed"`
		Author      *author       `json:"author"`
		HelpUrl     *fileLanguage `json:"helpUrl"`
		Description *fileLanguage `json:"description"`
	}
	//fileNode struct {
	//	Category string        `json:"category"`
	//	Icon     string        `json:"iconPath"`
	//	Label    *fileLanguage `json:"label"`
	//}
	fileSink struct {
		About  *fileAbout   `json:"about"`
		Libs   []string     `json:"libs"`
		Fields []*fileField `json:"properties"`
		Node   interface{}  `json:"node"`
	}
	language struct {
		English string `json:"en"`
		Chinese string `json:"zh"`
	}
	about struct {
		Trial       bool      `json:"trial"`
		Installed   bool      `json:"installed"`
		Author      *author   `json:"author"`
		HelpUrl     *language `json:"helpUrl"`
		Description *language `json:"description"`
	}
	field struct {
		Exist             bool        `json:"exist"`
		Name              string      `json:"name"`
		Default           interface{} `json:"default"`
		Type              string      `json:"type"`
		Control           string      `json:"control"`
		ConnectionRelated bool        `json:"connection_related"`
		Optional          bool        `json:"optional"`
		Values            interface{} `json:"values"`
		Hint              *language   `json:"hint"`
		Label             *language   `json:"label"`
	}
	node struct {
		Category string    `json:"category"`
		Icon     string    `json:"iconPath"`
		Label    *language `json:"label"`
	}
	uiSink struct {
		About  *about      `json:"about"`
		Libs   []string    `json:"libs"`
		Fields []field     `json:"properties"`
		Node   interface{} `json:"node"`
		Type   string      `json:"type,omitempty"`
	}
)

func newLanguage(fi *fileLanguage) *language { _ = "STUB: not implemented"; return nil }

func newField(fis []*fileField) (uis []field, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAbout(fi *fileAbout) *about { _ = "STUB: not implemented"; return nil }

func newUiSink(fi *fileSink) (*uiSink, error) { _ = "STUB: not implemented"; return nil, nil }

var gSinkmetadata = make(map[string]*uiSink) // immutable

func ReadSinkMetaDir(checker InstallChecker) error { _ = "STUB: not implemented"; return nil }

func readSinkMetaDir(folder string, checker InstallChecker) error {
	_ = "STUB: not implemented"
	return nil
}

func UninstallSink(name string) { _ = "STUB: not implemented"; return }

func ReadSinkMetaFile(filePath string, installed bool) error { _ = "STUB: not implemented"; return nil }

func GetSinkMeta(pluginName, language string) (s *uiSink, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type pluginfo struct {
	Name  string `json:"name"`
	About *about `json:"about"`
	Type  string `json:"type,omitempty"`
}

func GetSinks() (sinks []*pluginfo) { _ = "STUB: not implemented"; return nil }
