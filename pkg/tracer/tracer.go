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

//go:build trace || !core

package tracer

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const (
	TraceCfgKey = "$$tracer_cfg"
)

var globalTracerManager *GlobalTracerManager

func init() {
	globalTracerManager = &GlobalTracerManager{}
}

type GlobalTracerManager struct {
	syncx.RWMutex
	Init                 bool
	ServiceName          string
	EnableRemoteEndpoint bool
	RemoteEndpoint       string
	SpanExporter         *SpanExporter
}

func (g *GlobalTracerManager) InitIfNot() { _ = "STUB: not implemented"; return }

func (g *GlobalTracerManager) SetTracer(enableRemote bool, serviceName, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlobalTracerManager) GetTraceById(traceID string) (root *LocalSpan, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GlobalTracerManager) GetTraceByRuleID(ruleID string, limit int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetTracer() trace.Tracer { _ = "STUB: not implemented"; return *new(trace.Tracer) }

func GetSpanByTraceID(traceID string) (root *LocalSpan, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetTracer(config *TracerConfig) error { _ = "STUB: not implemented"; return nil }

func InitTracer() error { _ = "STUB: not implemented"; return nil }

func saveTracerConfig(config *TracerConfig) error { _ = "STUB: not implemented"; return nil }

func loadTracerConfig() (*TracerConfig, error) { _ = "STUB: not implemented"; return nil, nil }

func GetTraceIDListByRuleID(ruleID string, limit int64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
