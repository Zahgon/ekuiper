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

package bump

import (
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

const (
	currentVersion = 4
	bumpTable      = "eKuiperMeta_bump_version"
)

var GlobalBumpManager *BumpManager

type BumpManager struct {
	Version int
	store   kv.KeyValue
}

func InitBumpManager() error { _ = "STUB: not implemented"; return nil }

func loadVersionFromStorage() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// for the initialized case

func BumpToCurrentVersion(dataDir string) error { _ = "STUB: not implemented"; return nil }

func bumpFrom0To1(dir string) error { _ = "STUB: not implemented"; return nil }

func migrateDataIntoStorage(dataDir, confType string) error { _ = "STUB: not implemented"; return nil }

func storeGlobalVersion(ver int) error { _ = "STUB: not implemented"; return nil }
