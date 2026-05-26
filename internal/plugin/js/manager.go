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

package js

import (
	"github.com/lf-edge/ekuiper/v2/internal/binder"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

var (
	manager *Manager
	_       binder.FuncFactory = manager
)

func GetManager() *Manager { _ = "STUB: not implemented"; return nil }

type Manager struct {
	db             kv.KeyValue
	importStatusDb kv.KeyValue
}

type Script struct {
	Id     string `json:"id"`
	Desc   string `json:"description"`
	Script string `json:"script"`
	IsAgg  bool   `json:"isAgg"`
}

// InitManager initialize the manager, only called once by the server
func InitManager() error { _ = "STUB: not implemented"; return nil }

func (m *Manager) UpsertByJson(k string, v string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) Create(script *Script) error { _ = "STUB: not implemented"; return nil }

func validate(script *Script) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) GetScript(id string) (*Script, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Manager) List() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Manager) Update(script *Script) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) Delete(id string) error { _ = "STUB: not implemented"; return nil }
