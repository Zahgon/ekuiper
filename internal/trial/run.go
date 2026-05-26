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
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
)

type RunDef struct {
	Id        string                    `json:"id"`
	Sql       string                    `json:"sql"`
	Mock      map[string]map[string]any `json:"mockSource"`
	SinkProps map[string]any            `json:"sinkProps"`

	endpoint string
}

func genTrialRuleID(def *RunDef) string { _ = "STUB: not implemented"; return "" }

func genTrialRule(rd *RunDef, sinkProps map[string]interface{}) *def.Rule {
	_ = "STUB: not implemented"
	return nil
}

// Let trial rule always send out error to show

func create(def *RunDef) (*topo.Topo, error) { _ = "STUB: not implemented"; return nil, nil }

// Add trial run prefix for rule id to avoid duplicate rule id with real rules in runtime or other trial rule

func trialRun(tp *topo.Topo, endpoint string) { _ = "STUB: not implemented"; return }

// If stop by EOF
