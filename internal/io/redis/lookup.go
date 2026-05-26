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

package redis

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/redis/go-redis/v9"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/util"
)

type conf struct {
	// host:port address.
	Addr     string `json:"addr,omitempty"`
	Username string `json:"username,omitempty"`
	// Optional password. Must match the password specified in the
	Password string `json:"password,omitempty"`
	DataType string `json:"dataType,omitempty"`
	DB       string `json:"datasource,omitempty"`
}

type lookupSource struct {
	c   *conf
	db  int
	cli *redis.Client
}

func (s *lookupSource) Ping(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// use default DB

func (s *lookupSource) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *lookupSource) Connect(ctx api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// use default DB

func (s *lookupSource) Lookup(ctx api.StreamContext, _ []string, keys []string, values []any) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *lookupSource) Validate(props map[string]any) error { _ = "STUB: not implemented"; return nil }

func (s *lookupSource) Open(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (s *lookupSource) Close(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func GetLookupSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var (
	_ api.LookupSource  = &lookupSource{}
	_ util.PingableConn = &lookupSource{}
)
