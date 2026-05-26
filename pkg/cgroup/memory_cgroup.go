// Copyright 2023 EMQ Technologies Co., Ltd.
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

package cgroup

const (
	cGroupMemLimitPath = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	cGroupMemUsagePath = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
)

// MemTotalCGroup returns the total amount of RAM on this system in container environment.
func MemTotalCGroup() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// MemUsedCGroup returns the total used amount of RAM on this system in container environment.
func MemUsedCGroup() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// refer to https://github.com/containerd/cgroups/blob/318312a373405e5e91134d8063d04d59768a1bff/utils.go#L243
func readUint(path string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// refer to https://github.com/containerd/cgroups/blob/318312a373405e5e91134d8063d04d59768a1bff/utils.go#L251
func parseUint(s string, base, bitSize int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// 1. Handle negative values greater than MinInt64 (and)
// 2. Handle negative values lesser than MinInt64
