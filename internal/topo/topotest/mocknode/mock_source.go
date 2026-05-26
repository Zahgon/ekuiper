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

package mocknode

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/xsql"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type MockSource struct {
	data   []*xsql.Tuple
	offset int
	eof    api.EOFIngest
	syncx.RWMutex
}

const TIMELEAP = 200

func (m *MockSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockSource) SetEofIngest(eof api.EOFIngest) { _ = "STUB: not implemented"; return }

func (m *MockSource) Connect(_ api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockSource) Subscribe(ctx api.StreamContext, ingest api.TupleIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

// Mock timer, only send out the data once the mock time goes to the timestamp.
// Another mechanism must be imposed to move forward the mock time.

func (m *MockSource) GetOffset() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *MockSource) Rewind(offset interface{}) error { _ = "STUB: not implemented"; return nil }

func (m *MockSource) ResetOffset(input map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockSource) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

var (
	_ api.TupleSource = &MockSource{}
	_ api.Bounded     = &MockSource{}
)
