// Copyright 2021-2022 EMQ Technologies Co., Ltd.
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

package common

import (
	"testing"

	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

var (
	Keys   = []int64{1000, 1500, 2000, 3000}
	Values = []string{"bar1", "bar15", "bar2", "bar3"}
)

func TestTsSet(ks kv.Tskv, t *testing.T) { _ = "STUB: not implemented"; return }

func TestTsLast(ks kv.Tskv, t *testing.T) { _ = "STUB: not implemented"; return }

func TestTsGet(ks kv.Tskv, t *testing.T) { _ = "STUB: not implemented"; return }

func TestTsDelete(ks kv.Tskv, t *testing.T) { _ = "STUB: not implemented"; return }

func TestTsDeleteBefore(ks kv.Tskv, t *testing.T) { _ = "STUB: not implemented"; return }

func load(ks kv.Tskv, t *testing.T) { _ = "STUB: not implemented"; return }
