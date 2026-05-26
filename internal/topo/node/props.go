// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

package node

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

type SinkConf struct {
	Concurrency      int               `json:"concurrency"`
	Omitempty        bool              `json:"omitIfEmpty"`
	SendSingle       bool              `json:"sendSingle"`
	DataTemplate     string            `json:"dataTemplate"`
	Format           string            `json:"format"`
	SchemaId         string            `json:"schemaId"`
	Delimiter        string            `json:"delimiter"`
	BufferLength     int               `json:"bufferLength"`
	Fields           []string          `json:"fields"`
	ExcludeFields    []string          `json:"excludeFields"`
	DataField        string            `json:"dataField"`
	BatchSize        int               `json:"batchSize"`
	LingerInterval   cast.DurationConf `json:"lingerInterval"`
	Compression      string            `json:"compression"`
	CompressionProps map[string]any    `json:"compressionProps"`
	Encryption       string            `json:"encryption"`
	EncProps         map[string]any    `json:"encProps"`
	HasHeader        bool              `json:"hasHeader"`
	model.SinkConf
}

func ParseConf(logger api.Logger, props map[string]any) (*SinkConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
