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

package file

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type sinkConf struct {
	RollingInterval    cast.DurationConf `json:"rollingInterval"`
	RollingCount       int               `json:"rollingCount"`
	RollingNamePattern string            `json:"rollingNamePattern"` // where to add the timestamp to the file name
	RollingHook        string            `json:"rollingHook"`
	RollingHookProps   map[string]any    `json:"rollingHookProps"`
	RollingSize        int64             `json:"rollingSize"`
	CheckInterval      cast.DurationConf `json:"checkInterval"`
	Path               string            `json:"path"` // support dynamic property, when rolling, make sure the path is updated
	FileType           FileType          `json:"fileType"`
	HasHeader          bool              `json:"hasHeader"`
	Delimiter          string            `json:"delimiter"`
	Format             string            `json:"format"` // only use for validation; transformation is done in sink_node
	Compression        string            `json:"compression"`
	Encryption         string            `json:"encryption"`
	Fields             []string          `json:"fields"` // only use for extracting header for csv; transformation is done in sink_node
}

type fileSink struct {
	c *sinkConf

	mux      syncx.Mutex
	fws      map[string]*fileWriter
	rollHook modules.RollHook
	headers  string
}

func (m *fileSink) Provision(ctx api.StreamContext, props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fileSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the files have opened longer than the rolling interval, if so close it and create a new one

// this will never panic

// TODO how to deal with this error

func (m *fileSink) Collect(ctx api.StreamContext, tuple api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *fileSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (m *fileSink) roll(ctx api.StreamContext, k string, v *fileWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// The file will be created when the next item comes

// GetFws returns the file writer for the given file name, if the file writer does not exist, it will create one
// The item is used to get the csv header if needed
func (m *fileSink) GetFws(ctx api.StreamContext, fn string, item []byte) (*fileWriter, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var (
	_ api.BytesCollector = &fileSink{}
	_ model.StreamWriter = &fileSink{}
)
