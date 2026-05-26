// Copyright 2025 EMQ Technologies Co., Ltd.
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

package simulator

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type SimulatorLookupSource struct {
	cfg *sLookupConfig
}

func (s *SimulatorLookupSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimulatorLookupSource) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimulatorLookupSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimulatorLookupSource) Lookup(ctx api.StreamContext, lookupFields []string, cmpKeys []string, cmpValues []any) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type sLookupConfig struct {
	Data []map[string]any `json:"data"`
}
