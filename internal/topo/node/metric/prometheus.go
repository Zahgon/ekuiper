// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package metric

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var (
	prometheuseMetrics *PrometheusMetrics
	mutex              syncx.RWMutex
)

func GetPrometheusMetrics() *PrometheusMetrics { _ = "STUB: not implemented"; return nil }

type MetricGroup struct {
	TotalRecordsIn         *prometheus.CounterVec
	TotalRecordsOut        *prometheus.CounterVec
	TotalMessagesProcessed *prometheus.CounterVec
	TotalExceptions        *prometheus.CounterVec
	ProcessLatencyHist     *prometheus.HistogramVec
	ProcessLatency         *prometheus.GaugeVec
	BufferLength           *prometheus.GaugeVec
	ConnectionStatus       *prometheus.GaugeVec
}

type PrometheusMetrics struct {
	vecs []*MetricGroup
}

func newPrometheusMetrics() *PrometheusMetrics { _ = "STUB: not implemented"; return nil }

// prometheus initialization

// 10us ~ 5s

func (m *PrometheusMetrics) GetMetricsGroup(opType string) *MetricGroup {
	_ = "STUB: not implemented"
	return nil
}
