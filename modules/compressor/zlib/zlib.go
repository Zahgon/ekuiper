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

package zlib

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zlib"
)

func NewZlibCompressor() (*zlibCompressor, error) { _ = "STUB: not implemented"; return nil, nil }

type zlibCompressor struct {
	writer *zlib.Writer
	buffer bytes.Buffer
}

func (z *zlibCompressor) Compress(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewZlibDecompressor() (*zlibDecompressor, error) { _ = "STUB: not implemented"; return nil, nil }

type zlibDecompressor struct {
	reader io.ReadCloser
}

func (z *zlibDecompressor) Decompress(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
