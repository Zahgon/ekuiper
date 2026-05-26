// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

package http

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
)

type RestSink struct {
	*ClientConf
	noHeaderTemplate   bool
	noFormdataTemplate bool
}

var bodyTypeFormat = map[string]string{
	"json": "json",
	"form": "urlencoded",
}

func (r *RestSink) Provision(ctx api.StreamContext, configs map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RestSink) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (r *RestSink) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RestSink) Collect(ctx api.StreamContext, item api.RawTuple) error {
	_ = "STUB: not implemented"
	return nil
}

// if auth is set, the auth is handled by the client connect

// do not record response body error as it is not an error in the sink action.

func GetSink() api.Sink { _ = "STUB: not implemented"; return *new(api.Sink) }

var _ api.BytesCollector = &RestSink{}
