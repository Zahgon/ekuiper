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

package service

import (
	"net/http"
	"net/url"
	"time"

	// TODO: replace with `google.golang.org/protobuf/proto` pkg.
	//nolint:staticcheck
	//nolint:staticcheck
	//nolint:staticcheck
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"google.golang.org/grpc"
)

type exeIns func(desc descriptor, opt *interfaceOpt, i *interfaceInfo) (executor, error)

var executors = map[protocol]exeIns{
	GRPC: newGrpcExecutor,
	REST: newHttpExecutor,
}

func newHttpExecutor(desc descriptor, opt *interfaceOpt, i *interfaceInfo) (executor, error) {
	_ = "STUB: not implemented"
	return *new(executor), nil
}

func newGrpcExecutor(desc descriptor, opt *interfaceOpt, _ *interfaceInfo) (executor, error) {
	_ = "STUB: not implemented"
	return *new(executor), nil
}

// NewExecutor
// Each interface definition maps to one executor instance. It is supposed to have only one thread running.
func NewExecutor(i *interfaceInfo) (executor, error) {
	_ = "STUB: not implemented"
	// No validation here, suppose the validation has been done in json parsing
	return *new(executor), nil
}

type executor interface {
	InvokeFunction(ctx api.FunctionContext, name string, params []interface{}) (interface{}, error)
}

type interfaceOpt struct {
	addr    *url.URL
	timeout time.Duration
}

type grpcExecutor struct {
	descriptor protoDescriptor
	*interfaceOpt

	conn *grpc.ClientConn
}

func (d *grpcExecutor) InvokeFunction(_ api.FunctionContext, name string, params []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

// connect successfully, do nothing

// TODO reconnect if fail and error handling

// connect successfully, do nothing

type httpExecutor struct {
	descriptor multiplexDescriptor
	*interfaceOpt
	restOpt *restOption

	conn *http.Client
}

var testIndex int

func (h *httpExecutor) InvokeFunction(ctx api.FunctionContext, name string, params []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpExecutor) invokeFunction(ctx api.FunctionContext, name string, params []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
