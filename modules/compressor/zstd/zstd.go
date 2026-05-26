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

package zstd

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zstd"
)

type compressProps struct {
	WindowSize int `json:"windowSize"`
}

func NewZstdCompressor(props map[string]any) (*zstdCompressor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type zstdCompressor struct {
	writer *zstd.Encoder
	buffer bytes.Buffer
}

func (g *zstdCompressor) Compress(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewzstdDecompressor() (*zstdDecompressor, error) { _ = "STUB: not implemented"; return nil, nil }

type zstdDecompressor struct {
	decoder *zstd.Decoder
}

func (z *zstdDecompressor) Decompress(data []byte) ([]byte, error) {
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
