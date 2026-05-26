// Copyright 2021-2023 EMQ Technologies Co., Ltd.
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

package store

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/store/definition"
)

type StoreConf struct {
	Type         string
	ExtStateType string
	RedisConfig  definition.RedisConfig
	SqliteConfig definition.SqliteConfig
	FdbConfig    definition.FdbConfig
	PebbleConfig definition.PebbleConfig
}

func SetupDefault(dataDir string) error { _ = "STUB: not implemented"; return nil }

func SetupWithConfig(sc *StoreConf, setupCheckpointDB bool) error {
	_ = "STUB: not implemented"
	return nil
}

func Setup(config definition.Config, setupCheckpointDB bool) error {
	_ = "STUB: not implemented"
	return nil
}

// write sth to ensure checkpoint.db created
