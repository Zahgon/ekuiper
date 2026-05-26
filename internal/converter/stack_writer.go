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

package converter

import (
	"bytes"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

type StackWriter struct {
	// The internal writer. When flushing, create a new one.
	converter message.Converter
	buffer    *bytes.Buffer
}

func NewStackWriter(_ api.StreamContext, converter message.Converter) (message.ConvertWriter, error) {
	_ = "STUB: not implemented"
	return *new(message.ConvertWriter), nil
}

func (w *StackWriter) New(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (w *StackWriter) Write(ctx api.StreamContext, d any) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *StackWriter) Flush(ctx api.StreamContext) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
