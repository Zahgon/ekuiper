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

package delimited

import (
	"bytes"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

type CsvWriter struct {
	// The internal writer. When flushing, create a new one.
	converter *Converter
	buffer    *bytes.Buffer
	header    string
}

func NewCsvWriter(_ api.StreamContext, props map[string]any) (message.ConvertWriter, error) {
	_ = "STUB: not implemented"
	return *new(message.ConvertWriter), nil
}

// Header are now creating by batch writer

func (w *CsvWriter) New(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (w *CsvWriter) Write(ctx api.StreamContext, d any) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *CsvWriter) Flush(ctx api.StreamContext) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
