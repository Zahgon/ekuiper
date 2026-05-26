// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

package delimited

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

type Converter struct {
	Delimiter string   `json:"delimiter"`
	Cols      []string `json:"fields"`
	HasHeader bool     `json:"hasHeader"`
}

func NewConverter(props map[string]any) (message.Converter, error) {
	_ = "STUB: not implemented"
	return *new(message.Converter), nil
}

// Encode If no columns defined, the default order is sort by key
func (c *Converter) Encode(ctx api.StreamContext, d any) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode If the cols is not set, the default key name is col1, col2, col3...
// The return value is always a map
func (c *Converter) Decode(ctx api.StreamContext, b []byte) (ma any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
