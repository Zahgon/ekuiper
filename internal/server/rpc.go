// Copyright 2021-2025 EMQ Technologies Co., Ltd.
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

//go:build rpc || !core

package server

import (
	"net/http"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/model"
)

const QueryRuleId = "internal-ekuiper_query_rule"

func init() {
	servers["rpc"] = &rpcComp{}
}

type rpcComp struct {
	s *http.Server
}

func (r *rpcComp) register() { _ = "STUB: not implemented"; return }

func (r *rpcComp) serve() {
	_ = "STUB: not implemented"
	// Start rpc service
	return
}

func (r *rpcComp) close() { _ = "STUB: not implemented"; return }

type Server int

func (t *Server) CreateQuery(sql string, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func stopQuery() { _ = "STUB: not implemented"; return }

func (t *Server) GetQueryResult(_ string, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) Stream(stream string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) CreateRule(rule *model.RPCArgDesc, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) GetStatusRule(name string, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) GetTopoRule(name string, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) StartRule(name string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) StopRule(name string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) RestartRule(name string, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) DescRule(name string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) ShowRules(_ int, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) DropRule(name string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) ValidateRule(rule *model.RPCArgDesc, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) Import(file string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) Export(file string, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) ImportConfiguration(arg *model.ImportDataDesc, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Server) GetStatusImport(_ int, reply *string) error { _ = "STUB: not implemented"; return nil }

func (t *Server) ExportConfiguration(arg *model.ExportDataDesc, reply *string) error {
	_ = "STUB: not implemented"
	return nil
}

// do not specify rules, export all

func marshalDesc(m interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func initQuery() { _ = "STUB: not implemented"; return }
