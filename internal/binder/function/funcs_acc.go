// Copyright 2025 EMQ Technologies Co., Ltd.
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

package function

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

func registerGlobalAggFunc() { _ = "STUB: not implemented"; return }

func handleAccFunc(ctx api.FunctionContext, args []interface{}, accFunc accFunc) (*accStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleOnCondAccFunc(ctx api.FunctionContext, args []interface{}, validData bool, partitionKey string, status *accStatus, accFunc accFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func extractAccArgs(ctx api.FunctionContext, args []interface{}, accFunc accFunc) (validData bool, partitionKey string, status *accStatus, err error) {
	_ = "STUB: not implemented"
	return false, "", nil, nil
}

// accFuncWithCond execute acc function with onBegin and onReset condition with following 4 steps:
// 1. Check HasBegin at the beginning, if it's false, it means any result won't be calculated, thus we need to always reset the value
// 2. Check onBegin condition to set the HasBegin
// 3. Check HasBegin to determine whether calculate the acc function
// 4. Check onReset to set the HasBegin
func accFuncWithCond(ctx api.FunctionContext, value interface{}, onBegin, onReset bool, validData bool, partitionKey string, status *accStatus, accFunc accFunc) {
	_ = "STUB: not implemented"
	return
}

type accStatus struct {
	Err      error
	Value    interface{}
	HasBegin bool
}

type accFunc interface {
	accFuncExec(ctx api.FunctionContext, value interface{}, validData bool, partitionKey string, status *accStatus, skipStatusSave bool)
	accReset(status *accStatus)
}

type accCountFunc struct{}

func (a accCountFunc) accFuncExec(ctx api.FunctionContext, value interface{}, validData bool, partitionKey string, status *accStatus, skipStatusSave bool) {
	_ = "STUB: not implemented"
	return
}

func (a accCountFunc) accReset(status *accStatus) { _ = "STUB: not implemented"; return }

type accSumFunc struct{}

func (a accSumFunc) accFuncExec(ctx api.FunctionContext, value interface{}, validData bool, partitionKey string, status *accStatus, skipStatusSave bool) {
	_ = "STUB: not implemented"
	return
}

// only for unit test

func (a accSumFunc) accReset(status *accStatus) { _ = "STUB: not implemented"; return }

type accMinFunc struct{}

func (a accMinFunc) accFuncExec(ctx api.FunctionContext, value interface{}, validData bool, partitionKey string, status *accStatus, skipStatusSave bool) {
	_ = "STUB: not implemented"
	return
}

func (a accMinFunc) accReset(status *accStatus) { _ = "STUB: not implemented"; return }

type accMaxFunc struct{}

func (a accMaxFunc) accFuncExec(ctx api.FunctionContext, value interface{}, validData bool, partitionKey string, status *accStatus, skipStatusSave bool) {
	_ = "STUB: not implemented"
	return
}

func (a accMaxFunc) accReset(status *accStatus) { _ = "STUB: not implemented"; return }

type accAvgFunc struct{}

func (a accAvgFunc) accFuncExec(ctx api.FunctionContext, value interface{}, validData bool, partitionKey string, status *accStatus, skipStatusSave bool) {
	_ = "STUB: not implemented"
	return
}

func (a accAvgFunc) accReset(status *accStatus) { _ = "STUB: not implemented"; return }

type accAvgStatus struct {
	sum   float64
	count int64
	avg   float64
}
