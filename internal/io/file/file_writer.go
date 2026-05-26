// Copyright 2023-2025 EMQ Technologies Co., Ltd.
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

package file

import (
	"io"
	"os"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/io/file/writer"
)

type fileWriter struct {
	File       *os.File
	Writer     io.Writer
	Hook       writerHooks
	Start      time.Time
	Count      int
	Size       int64
	Compress   string
	fileBuffer *writer.BufioWrapWriter
	// Whether the file has written any data. It is only used to determine if new line is needed when writing data.
	Written bool
}

func (m *fileSink) createFileWriter(ctx api.StreamContext, fn string, ft FileType, headers string, compressAlgorithm string, encryption string) (_ *fileWriter, ge error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *fileSink) CreateWriter(_ api.StreamContext, currWriter io.Writer, compression string, encryption string) (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

func (fw *fileWriter) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

// Close the compressor and encryptor firstly
