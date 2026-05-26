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
	"bufio"
	"io"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

func init() {
	modules.RegisterFileStreamReader("lines", func(ctx api.StreamContext) modules.FileStreamReader {
		return &LinesReader{}
	})
}

type LinesReader struct {
	scanner *bufio.Scanner
}

func (r *LinesReader) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *LinesReader) Bind(ctx api.StreamContext, fileStream io.Reader, maxSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// default to 1MB

func (r *LinesReader) Read(ctx api.StreamContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *LinesReader) IsBytesReader() bool { _ = "STUB: not implemented"; return false }

func (r *LinesReader) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }
