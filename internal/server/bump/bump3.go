// Copyright 2024 EMQ Technologies Co., Ltd.
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

package bump

import (
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/store"
)

func bumpFrom2TO3() error { _ = "STUB: not implemented"; return nil }

func rewriteSQLSinkConfiguration() error { _ = "STUB: not implemented"; return nil }

func rewriteSQLSourceConfiguration() error { _ = "STUB: not implemented"; return nil }

func rewriteCfg(cfg *OriginSqlSourceCfg) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func extractIndexField(icfg *OriginInternalSqlQueryCfg) *store.IndexField {
	_ = "STUB: not implemented"
	return nil
}

func extractIndexField2(icfg *OriginTemplateSqlQueryCfg) *store.IndexField {
	_ = "STUB: not implemented"
	return nil
}

// OriginSqlSourceCfg tends to rewrite index Field into index Fields
type OriginSqlSourceCfg struct {
	DBUrl               string                     `json:"dburl"`
	Interval            cast.DurationConf          `json:"interval"`
	InternalSqlQueryCfg *OriginInternalSqlQueryCfg `json:"internalSqlQueryCfg,omitempty"`
	TemplateSqlQueryCfg *OriginTemplateSqlQueryCfg `json:"templateSqlQueryCfg,omitempty"`
}

type OriginInternalSqlQueryCfg struct {
	Table                    string              `json:"table"`
	Limit                    int                 `json:"limit"`
	IndexFieldName           string              `json:"indexField"`
	IndexFieldValue          interface{}         `json:"indexValue"`
	IndexFieldDataType       string              `json:"indexFieldType"`
	IndexFieldDateTimeFormat string              `json:"dateTimeFormat"`
	IndexFields              []*store.IndexField `json:"indexFields"`
}

type OriginTemplateSqlQueryCfg struct {
	TemplateSQL              string              `json:"templateSql"`
	IndexFieldName           string              `json:"indexField"`
	IndexFieldValue          interface{}         `json:"indexValue"`
	IndexFieldDataType       string              `json:"indexFieldType"`
	IndexFieldDateTimeFormat string              `json:"dateTimeFormat"`
	IndexFields              []*store.IndexField `json:"indexFields"`
}
