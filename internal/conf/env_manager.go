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

package conf

func init() {
	globalEnvManager = &EnvManager{}
	SetupEnv()
}

func SetupEnv() { _ = "STUB: not implemented"; return }

func GetEnv() map[string]string { _ = "STUB: not implemented"; return nil }

var globalEnvManager *EnvManager

type EnvManager struct {
	env             map[string]string
	connectionProps map[string]map[string]map[string]interface{}
}

func (e *EnvManager) Setup() { _ = "STUB: not implemented"; return }

func (e *EnvManager) GetEnv() map[string]string { _ = "STUB: not implemented"; return nil }

func (e *EnvManager) loadEnv() { _ = "STUB: not implemented"; return }
