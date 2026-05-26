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

package tspoint

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type WriteOptions struct {
	PrecisionStr string `json:"precision"`

	Tags        map[string]string `json:"tags"`
	TsFieldName string            `json:"tsFieldName"`
	Fields      []string          `json:"fields"`
}

func (o *WriteOptions) Validate() error { _ = "STUB: not implemented"; return nil }

// no error

func (o *WriteOptions) ValidateTagTemplates(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

type RawPoint struct {
	Fields map[string]any
	Tags   map[string]string
	Tt     time.Time
	Ts     int64
}

func SinkTransform(ctx api.StreamContext, data any, options *WriteOptions) ([]*RawPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO possible problem here that the ts filed is transformed out

// Method to convert map to influxdb point, including the sink transforms + map to point
func singleMapToPoint(ctx api.StreamContext, dd map[string]any, options *WriteOptions) (*RawPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Internal method to transform map to influxdb point
func mapToPoint(ctx api.StreamContext, mm map[string]any, options *WriteOptions, tt time.Time, ts int64) (*RawPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertAll has no error

// Internal method to get time from map with tsFieldName
func getTime(data map[string]any, tsFieldName string, precisionStr string) (time.Time, int64, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), 0, nil
}

func getTS(data map[string]any, tsFieldName string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
