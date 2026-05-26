// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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

package props

import (
	"github.com/bwmarrin/snowflake"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var (
	SC     = &StaticConf{props: map[string]string{}}
	sfnode *snowflake.Node
)

type StaticConf struct {
	syncx.RWMutex
	props map[string]string
}

func InitProps() { _ = "STUB: not implemented"; return }

func (s *StaticConf) Get(propName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *StaticConf) Set(propName string, value string) { _ = "STUB: not implemented"; return }
