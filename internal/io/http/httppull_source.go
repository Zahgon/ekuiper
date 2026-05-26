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

package http

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type HttpPullSource struct {
	*ClientConf
	lastMD5 string
	psc     *pullSourceConfig
}

func (hps *HttpPullSource) GetOffset() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (hps *HttpPullSource) Rewind(offset any) error { _ = "STUB: not implemented"; return nil }

func (hps *HttpPullSource) ResetOffset(input map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (hps *HttpPullSource) Pull(ctx api.StreamContext, trigger time.Time, ingest api.TupleIngest, ingestError api.ErrorIngest) {
	_ = "STUB: not implemented"
	return
}

func (hps *HttpPullSource) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (hps *HttpPullSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

type pullSourceConfig struct {
	Path   string         `json:"datasource"`
	States map[string]any `json:"states"`
}

func (hps *HttpPullSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (hps *HttpPullSource) doPull(ctx api.StreamContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hps *HttpPullSource) doPullInternal(ctx api.StreamContext, c *ClientConf, lastMD5 string) ([]map[string]any, string, error) {
	_ = "STUB: not implemented"
	// if auth is set, the auth is handled by the client connect
	return nil, "", nil
}

func (hps *HttpPullSource) updateState(results []map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var _ api.PullTupleSource = &HttpPullSource{}
