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

//go:build !core || (rpc && plugin)

package server

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/model"
)

func (t *Server) RegisterPlugin(arg *model.PluginDesc, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) ShowUdfs(_ int, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) DescUdf(arg string, reply *string) error { _ = "STUB: not implemented"; return nil }
