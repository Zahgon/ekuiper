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

package gzip

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/gzip"
)

func NewGzipCompressor() (*gzipCompressor, error) { _ = "STUB: not implemented"; return nil, nil }

type gzipCompressor struct {
	writer *gzip.Writer
	buffer bytes.Buffer
}

func (g *gzipCompressor) Compress(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGzipDecompressor() (*gzipDecompressor, error) { _ = "STUB: not implemented"; return nil, nil }

type gzipDecompressor struct {
	reader *gzip.Reader
}

func (z *gzipDecompressor) Decompress(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReader(r io.Reader) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func NewWriter(w io.Writer) (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}
