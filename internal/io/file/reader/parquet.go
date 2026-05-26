// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

//go:build parquet || full

package reader

import (
	"io"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/parquet-go/parquet-go"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

func init() {
	modules.RegisterFileStreamReader("parquet", func(ctx api.StreamContext) modules.FileStreamReader {
		return &ParquetReader{}
	})
}

type ParquetReader struct {
	pf         *parquet.File
	groups     []parquet.RowGroup
	curGroup   int
	rowsReader parquet.Rows
}

func (pr *ParquetReader) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *ParquetReader) Bind(ctx api.StreamContext, fr io.Reader, _ int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pr *ParquetReader) Read(_ api.StreamContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (pr *ParquetReader) IsBytesReader() bool { _ = "STUB: not implemented"; return false }

func (pr *ParquetReader) Close(_ api.StreamContext) error { _ = "STUB: not implemented"; return nil }

var _ modules.FileStreamReader = &ParquetReader{}
