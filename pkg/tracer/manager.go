// Copyright 2024-2024 EMQ Technologies Co., Ltd.
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

//go:build trace || !core

package tracer

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type SpanExporter struct {
	remoteSpanExport *otlptrace.Exporter
	spanStorage      LocalSpanStorage
}

func NewSpanExporter(remoteCollector bool, remoteEndpoint string) (*SpanExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *SpanExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *SpanExporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *SpanExporter) GetTraceById(traceID string) (*LocalSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *SpanExporter) GetTraceByRuleID(ruleID string, limit int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LocalSpanStorage interface {
	SaveSpan(span sdktrace.ReadOnlySpan) error
	GetTraceById(traceID string) (*LocalSpan, error)
	GetTraceByRuleID(ruleID string, limit int64) ([]string, error)
}

type LocalSpanMemoryStorage struct {
	syncx.RWMutex
	queue *Queue
	// traceid -> spanid -> span
	m map[string]map[string]*LocalSpan
	// rule -> traceID, traceIDs will have duplicates, need to dedup when return
	ruleTraces map[string][]string
}

func newLocalSpanMemoryStorage(capacity int) *LocalSpanMemoryStorage {
	_ = "STUB: not implemented"
	return nil
}

func (l *LocalSpanMemoryStorage) SaveSpan(span sdktrace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LocalSpanMemoryStorage) saveSpan(localSpan *LocalSpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LocalSpanMemoryStorage) GetTraceById(traceID string) (*LocalSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LocalSpanMemoryStorage) GetTraceByRuleID(ruleID string, limit int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findRootSpan(allSpans map[string]*LocalSpan) *LocalSpan { _ = "STUB: not implemented"; return nil }

func buildSpanLink(cur *LocalSpan, OtherSpans map[string]*LocalSpan) {
	_ = "STUB: not implemented"
	// should only build once?
	return
}

// Queue is traceID FIFO queue with sized capacity
type Queue struct {
	m        map[string]struct{}
	items    []string
	capacity int
}

func NewQueue(capacity int) *Queue { _ = "STUB: not implemented"; return nil }

func (q *Queue) Enqueue(item *LocalSpan) string { _ = "STUB: not implemented"; return "" }

func (q *Queue) Dequeue() string { _ = "STUB: not implemented"; return "" }

func (q *Queue) Len() int { _ = "STUB: not implemented"; return 0 }

type sqlSpanStorage struct{}

func newSqlspanStorage() *sqlSpanStorage { _ = "STUB: not implemented"; return nil }

func (s *sqlSpanStorage) SaveSpan(span sdktrace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlSpanStorage) GetTraceById(traceID string) (*LocalSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlSpanStorage) GetTraceByRuleID(ruleID string, limit int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlSpanStorage) saveLocalSpan(span *LocalSpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlSpanStorage) loadTraceByRuleID(ruleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlSpanStorage) loadTraceByTraceID(traceID string) (*LocalSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gcSqliteSpan() error { _ = "STUB: not implemented"; return nil }
