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

package trial

import (
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// TrialManager Manager Initialized in the binder
var TrialManager = &Manager{
	runs: make(map[string]Run),
}

// Manager In memory manager for all trial rules
type Manager struct {
	syncx.RWMutex
	// ruleId -> *Topo
	runs map[string]Run
}

type Run struct {
	def  *RunDef
	topo *topo.Topo
}

func (m *Manager) CreateRule(ruleDef string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If the rule exists, stop it first

func (m *Manager) StopRule(ruleId string) { _ = "STUB: not implemented"; return }

func (m *Manager) StartRule(ruleId string) error { _ = "STUB: not implemented"; return nil }
