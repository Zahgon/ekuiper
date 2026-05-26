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

//go:build (!core && !wasmedge) || (rpc && portable && plugin)

package server

import (
	"github.com/lf-edge/ekuiper/v2/internal/plugin"
)

func (t *Server) doRegister(pt plugin.PluginType, p plugin.Plugin) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) doDelete(pt plugin.PluginType, name string, stopRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) doDesc(pt plugin.PluginType, name string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Server) doShow(pt plugin.PluginType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
