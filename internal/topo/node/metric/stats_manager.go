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
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const (
	RecordsInTotal                    = "records_in_total"
	RecordsOutTotal                   = "records_out_total"
	MessagesProcessedTotal            = "messages_processed_total"
	ProcessLatencyUs                  = "process_latency_us"
	ProcessLatencyUsHist              = "process_latency_us_hist"
	LastInvocation                    = "last_invocation"
	BufferLength                      = "buffer_length"
	ExceptionsTotal                   = "exceptions_total"
	LastException                     = "last_exception"
	LastExceptionTime                 = "last_exception_time"
	ConnectionStatus                  = "connection_status"
	ConnectionLastConnectedTime       = "connection_last_connected_time"
	ConnectionLastDisconnectedTime    = "connection_last_disconnected_time"
	ConnectionLastDisconnectedMessage = "connection_last_disconnected_message"
	ConnectionLastTryTime             = "connection_last_try_time"
)

var MetricNames = []string{RecordsInTotal, RecordsOutTotal, MessagesProcessedTotal, ProcessLatencyUs, BufferLength, LastInvocation, ExceptionsTotal, LastException, LastExceptionTime, ConnectionStatus, ConnectionLastConnectedTime, ConnectionLastDisconnectedTime, ConnectionLastDisconnectedMessage, ConnectionLastTryTime}

type StatManager interface {
	IncTotalRecordsIn()
	IncTotalRecordsOut()
	IncTotalMessagesProcessed(n int64)
	IncTotalExceptions(err string)
	ProcessTimeStart()
	ProcessTimeEnd()
	SetBufferLength(l int64)
	SetProcessTimeStart(t time.Time)
	// 0 is connecting, 1 is connected, -1 is disconnected
	SetConnectionState(state string, message string)
	GetMetrics() []any
	// Clean remove all metrics history
	Clean(ruleId string)
}

// DefaultStatManager The statManager is not thread safe. Make sure it is used in only one instance
type DefaultStatManager struct {
	// metrics
	totalRecordsIn         int64
	totalRecordsOut        int64
	totalMessagesProcessed int64

	processLatency    int64
	lastInvocation    time.Time
	bufferLength      int64
	totalExceptions   int64
	lastException     string
	lastExceptionTime time.Time

	connectionState *ConnectionStatManager
	// configs
	opType           string //"source", "op", "sink"
	prefix           string
	processTimeStart time.Time
	opId             string
	instanceId       int
	syncx.RWMutex
}

func NewStatManager(ctx api.StreamContext, opType string) StatManager {
	_ = "STUB: not implemented"
	return *new(StatManager)
}

func (sm *DefaultStatManager) SetConnectionState(status string, message string) {
	_ = "STUB: not implemented"
	return
}

func (sm *DefaultStatManager) IncTotalRecordsIn() { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) IncTotalMessagesProcessed(n int64) { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) IncTotalRecordsOut() { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) IncTotalExceptions(err string) { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) incTotalExceptions(err string) { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) ProcessTimeStart() { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) ProcessTimeEnd() { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) SetBufferLength(l int64) { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) SetProcessTimeStart(t time.Time) { _ = "STUB: not implemented"; return }

func (sm *DefaultStatManager) GetMetrics() []any { _ = "STUB: not implemented"; return nil }

func (sm *DefaultStatManager) Clean(_ string) {
	_ = "STUB: not implemented"
	// do nothing
	return
}

type ConnectionStatManager struct {
	connStatus         int
	lastConnectedTime  time.Time
	lastTryTime        time.Time
	lastDisconnect     string
	lastDisconnectTime time.Time
}

func (csm *ConnectionStatManager) SetConnectionState(state string, message string) {
	_ = "STUB: not implemented"
	return
}

func setMemConnState(csm *ConnectionStatManager, state string, message string) {
	_ = "STUB: not implemented"
	return
}
