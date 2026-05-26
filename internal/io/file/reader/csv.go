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

package reader

import (
	"encoding/csv"
	"io"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

func init() {
	modules.RegisterFileStreamReader("csv", func(ctx api.StreamContext) modules.FileStreamReader {
		return &CsvReader{}
	})
}

type csvConf struct {
	HasHeader bool     `json:"hasHeader"`
	Columns   []string `json:"columns"`
	Delimiter string   `json:"delimiter"`
}

type CsvReader struct {
	csvR   *csv.Reader
	config *csvConf

	cols []string
}

func (r *CsvReader) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *CsvReader) Bind(ctx api.StreamContext, fileStream io.Reader, _ int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *CsvReader) Read(ctx api.StreamContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *CsvReader) IsBytesReader() bool { _ = "STUB: not implemented"; return false }

func (r *CsvReader) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }
