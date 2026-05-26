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

package image

import (
	"context"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type c struct {
	Path        string `json:"path"`
	ImageFormat string `json:"imageFormat"`
	MaxAge      int    `json:"maxAge"`
	MaxCount    int    `json:"maxCount"`
}

type imageSink struct {
	c      *c
	cancel context.CancelFunc
}

func (m *imageSink) Provision(_ api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *imageSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *imageSink) delFile(logger api.Logger) error { _ = "STUB: not implemented"; return nil }

func (m *imageSink) getSuffix() string { _ = "STUB: not implemented"; return "" }

func (m *imageSink) saveFile(b []byte, fpath string) error { _ = "STUB: not implemented"; return nil }

func (m *imageSink) saveFiles(images map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *imageSink) Collect(ctx api.StreamContext, item api.MessageTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *imageSink) CollectList(ctx api.StreamContext, items api.MessageTupleList) error {
	_ = "STUB: not implemented"
	// TODO handle partial errors
	return nil
}

func (m *imageSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var _ api.TupleCollector = &imageSink{}
