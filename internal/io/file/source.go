// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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
	"encoding/gob"
	"io"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	_ "github.com/lf-edge/ekuiper/v2/internal/io/file/reader"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type SourceConfig struct {
	FileName         string            `json:"datasource"`
	FileType         string            `json:"fileType"`
	Path             string            `json:"path"`
	Interval         cast.DurationConf `json:"interval"`
	IsTable          bool              `json:"isTable"`
	SendInterval     cast.DurationConf `json:"sendInterval"`
	ActionAfterRead  int               `json:"actionAfterRead"`
	MoveTo           string            `json:"moveTo"`
	IgnoreStartLines int               `json:"ignoreStartLines"`
	IgnoreEndLines   int               `json:"ignoreEndLines"`
	// Only use for planning
	Decompression string `json:"decompression"`
	// state
	rewindMeta *FileDirSourceRewindMeta
}

// Source load data from file system.
// Depending on file types, it may read line by line like lines, csv.
// Otherwise, it reads the file as a whole and send to company reader node to read and split.
// The planner need to plan according to the file type.
type Source struct {
	file   string
	isDir  bool
	config *SourceConfig
	reader modules.FileStreamReader
	// attach to a reader
	decorator modules.FileStreamDecorator
	eof       api.EOFIngest
	// rewind support state
	rewindMeta *FileDirSourceRewindMeta
}

func (fs *Source) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO support later. If decompression is set, we need to read in the whole file

func (fs *Source) Connect(_ api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Pull file source may ingest bytes or tuple
// For stream source, it ingest one line
// For batch source, it ingest the whole file, thus it need a reader node to coordinate and read the content into lines/array
func (fs *Source) Pull(ctx api.StreamContext, _ time.Time, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

func (fs *Source) SetEofIngest(eof api.EOFIngest) { _ = "STUB: not implemented"; return }

func (fs *Source) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

type WithTime struct {
	name       string
	modifyTime time.Time
}

type WithTimeSlice []WithTime

func (f WithTimeSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (f WithTimeSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (f WithTimeSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (fs *Source) Load(ctx api.StreamContext, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

// may be just forget to put in the file

func (fs *Source) parseFile(ctx api.StreamContext, file string, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

// This is the buffer size, 1MB by default

// Read line or read all

// have error, do not need to do action after read

func ignoreLines(ctx api.StreamContext, reader io.Reader, decorator modules.FileStreamDecorator, ignoreStartLines int, ignoreEndLines int) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// This is a queue to store the lines that should be ignored

// Send EOF to decorator

// the last n line are left in the tempLines

// first round

func (fs *Source) Info() (i model.NodeInfo) {
	_ = "STUB: not implemented"
	// output batch raw, so need encrypt/decompress as a whole and then decode as a whole
	return *new(model.NodeInfo)
}

// decrypt/decompress in scan and output raw

// decrypt/decompress in scan and output decoded tuple

// TransformType must call after provision
func (fs *Source) TransformType() api.Source {
	_ = "STUB: not implemented"
	// If interval is not set, use watch source
	return *new(api.Source)
}

/// Rewind support

type FileDirSourceRewindMeta struct {
	LastModifyTime time.Time `json:"lastModifyTime"`
}

func init() {
	gob.Register(time.Time{})
	gob.Register(&FileDirSourceRewindMeta{})
}

func (fs *Source) GetOffset() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (fs *Source) Rewind(offset any) error { _ = "STUB: not implemented"; return nil }

func (fs *Source) ResetOffset(_ map[string]any) error { _ = "STUB: not implemented"; return nil }

func (fs *Source) checkFileRead(fileName string) (bool, time.Time, error) {
	_ = "STUB: not implemented"
	return false, *new(time.Time), nil
}

func (fs *Source) updateRewindMeta(_ string, modifyTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func GetSource() api.Source {
	_ = "STUB: not implemented"
	return *

	// ingest possibly []byte and tuple
	new(api.Source)
}

var (
	_ api.PullTupleSource = &Source{}
	// if interval is not set, it uses inotify
	_ api.Bounded    = &Source{}
	_ model.InfoNode = &Source{}
	_ api.Rewindable = &Source{}
)
