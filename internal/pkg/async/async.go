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

package async

import (
	"context"

	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const (
	TaskRegisterStatus = "register"
	TaskRunningStatus  = "running"
	TaskFinishStatus   = "finish"
	TaskErrorStatus    = "error"
	TaskCancelStatus   = "cancel"
)

var GlobalAsyncManager *AsyncManager

func InitManager() error { _ = "STUB: not implemented"; return nil }

type AsyncManager struct {
	syncx.RWMutex
	asyncDB    kvInterface
	taskCancel map[string]context.CancelFunc
}

func (m *AsyncManager) GetTaskIDList() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *AsyncManager) GetTask(taskID string) (*AsyncTaskStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *AsyncManager) RegisterTask(taskID string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (m *AsyncManager) CancelTask(taskID string) error { _ = "STUB: not implemented"; return nil }

func (m *AsyncManager) StartTask(taskID string) error { _ = "STUB: not implemented"; return nil }

func (m *AsyncManager) FinishTask(taskID, msg string) error { _ = "STUB: not implemented"; return nil }

func (m *AsyncManager) TaskFailed(taskID string, errMsg error) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *AsyncManager) isTaskExists(taskID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type AsyncTaskStatus struct {
	TaskID           string `json:"id"`
	Status           string `json:"status"`
	Message          string `json:"message"`
	CreatedTimestamp int64  `json:"createdTimestamp"`
	UpdatedTimestamp int64  `json:"updatedTimestamp"`
}

func (m *AsyncManager) getTaskStatus(taskID string) (*AsyncTaskStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *AsyncManager) storeTaskStatus(taskID string, taskStatus *AsyncTaskStatus) error {
	_ = "STUB: not implemented"
	return nil
}

type kvInterface interface {
	Set(key string, value string) error
	Get(key string) (string, bool, error)
	Keys() (keys []string, err error)
}

type asyncKV struct {
	m kv.KeyValue
}

func (k *asyncKV) Set(key string, value string) error { _ = "STUB: not implemented"; return nil }

func (k *asyncKV) Get(key string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (k *asyncKV) Keys() (keys []string, err error) { _ = "STUB: not implemented"; return nil, nil }
