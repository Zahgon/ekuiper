// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

// EncryptNode encrypt raw bytes
// Immutable: false
// Input: RawTuple
// Output: RawTuple
type EncryptNode struct {
	*defaultSinkNode
	tool message.Encryptor
}

func NewEncryptOp(name string, rOpt *def.RuleOption, encryptMethod string, encProps map[string]any) (*EncryptNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *EncryptNode) Exec(ctx api.StreamContext, errCh chan<- error) {
	_ = "STUB: not implemented"
	return
}

func (o *EncryptNode) Worker(_ api.StreamContext, item any) []any {
	_ = "STUB: not implemented"
	return nil
}
