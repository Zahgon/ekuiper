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

package tracenode

import (
	"context"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"go.opentelemetry.io/otel/trace"

	topoContext "github.com/lf-edge/ekuiper/v2/internal/topo/context"
	"github.com/lf-edge/ekuiper/v2/internal/xsql"
)

const (
	DataKey = "data"
	RuleKey = "rule"
)

func RecordRowOrCollection(input interface{}, span trace.Span) { _ = "STUB: not implemented"; return }

func TraceInput(ctx api.StreamContext, d any, opName string, opts ...trace.SpanStartOption) (bool, api.StreamContext, trace.Span) {
	_ = "STUB: not implemented"
	return false, *new(api.StreamContext), *new(trace.Span)
}

func StartTraceBackground(ctx api.StreamContext, opName string, opts ...trace.SpanStartOption) (bool, api.StreamContext, trace.Span) {
	_ = "STUB: not implemented"
	return false, *new(api.StreamContext), *new(trace.Span)
}

func StartTraceByID(ctx api.StreamContext, parentId string, opts ...trace.SpanStartOption) (bool, api.StreamContext, trace.Span) {
	_ = "STUB: not implemented"
	return false, *new(api.StreamContext), *new(trace.Span)
}

func ToStringRow(r xsql.Row) string { _ = "STUB: not implemented"; return "" }

func ToStringCollection(r api.MessageTupleList) string { _ = "STUB: not implemented"; return "" }

// TODO all tuple list must be treated the same in the future. Let ToMaps work anywhere

func BuildTraceParentId(traceID [16]byte, spanID [8]byte) string {
	_ = "STUB: not implemented"
	return ""
}

func checkCtxByStrategy(ctx, tracerCtx api.StreamContext) bool {
	_ = "STUB: not implemented"
	return false
}

func ExtractStrategy(ctx api.StreamContext) topoContext.TraceStrategy {
	_ = "STUB: not implemented"
	return *new(topoContext.TraceStrategy)
}

func hasTraceContext(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
