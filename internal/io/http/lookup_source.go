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

package http

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type HttpLookupSource struct {
	*ClientConf
}

func (hls *HttpLookupSource) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (hls *HttpLookupSource) Close(ctx api.StreamContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (hls *HttpLookupSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (hls *HttpLookupSource) Lookup(ctx api.StreamContext, fields []string, keys []string, values []any) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pruneData(data []map[string]any, fields []string) []map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func findColumn(column string, fields []string) bool { _ = "STUB: not implemented"; return false }

func lookupJoin(dataMap []map[string]interface{}, keys []string, values []interface{}) []map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func GetLookUpSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var _ api.LookupSource = &HttpLookupSource{}

func doPull(ctx api.StreamContext, c *ClientConf, lastMD5 string) ([]map[string]any, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
