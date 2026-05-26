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
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/prometheus/client_golang/prometheus"
)

func getStatManager(ctx api.StreamContext, dsm *DefaultStatManager) (StatManager, error) {
	_ = "STUB: not implemented"
	return *new(StatManager), nil
}

// assign prometheus

type PrometheusStatManager struct {
	*DefaultStatManager
	// prometheus metrics
	pTotalMessagesProcessed prometheus.Counter
	pTotalRecordsIn         prometheus.Counter
	pTotalRecordsOut        prometheus.Counter
	pTotalExceptions        prometheus.Counter
	pProcessLatency         prometheus.Gauge
	pProcessLatencyHist     prometheus.Observer
	pBufferLength           prometheus.Gauge
	pConnectionStatus       prometheus.Gauge
}

func (sm *PrometheusStatManager) IncTotalRecordsIn() { _ = "STUB: not implemented"; return }

func (sm *PrometheusStatManager) IncTotalMessagesProcessed(n int64) {
	_ = "STUB: not implemented"
	return
}

func (sm *PrometheusStatManager) IncTotalRecordsOut() { _ = "STUB: not implemented"; return }

func (sm *PrometheusStatManager) IncTotalExceptions(err string) { _ = "STUB: not implemented"; return }

func (sm *PrometheusStatManager) ProcessTimeEnd() { _ = "STUB: not implemented"; return }

func (sm *PrometheusStatManager) SetBufferLength(l int64) { _ = "STUB: not implemented"; return }

func (sm *PrometheusStatManager) Clean(ruleId string) { _ = "STUB: not implemented"; return }

func (sm *PrometheusStatManager) SetConnectionState(state string, message string) {
	_ = "STUB: not implemented"
	return
}
