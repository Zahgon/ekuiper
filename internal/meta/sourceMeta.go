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

import (
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type (
	fileSource struct {
		About      *fileAbout              `json:"about"`
		Libs       []string                `json:"libs"`
		DataSource interface{}             `json:"dataSource,omitempty"`
		ConfKeys   map[string][]*fileField `json:"properties"`
		Node       interface{}             `json:"node"`
	}
	uiSource struct {
		About      *about             `json:"about"`
		Libs       []string           `json:"libs"`
		DataSource interface{}        `json:"dataSource,omitempty"`
		ConfKeys   map[string][]field `json:"properties"`
		Node       interface{}        `json:"node"`
		Type       string             `json:"type,omitempty"`
		isScan     bool
		isLookup   bool
	}
)

func newUiSource(fi *fileSource, isScan bool, isLookup bool) (*uiSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	gSourcemetaLock = syncx.RWMutex{}
	gSourcemetadata = make(map[string]*uiSource)
)

func UninstallSource(name string) { _ = "STUB: not implemented"; return }

func ReadSourceMetaFile(filePath string, isScan bool, isLookup bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO currently, only show installed source in ui

func ReadSourceMetaDir(scanChecker InstallChecker, lookupChecker InstallChecker) error {
	_ = "STUB: not implemented"
	// load etc/sources meta data
	return nil
}

func GetSourceMeta(sourceName, language string) (ptrSourceProperty *uiSource, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetSourcesPlugins(kind string) (sources []*pluginfo) { _ = "STUB: not implemented"; return nil }
