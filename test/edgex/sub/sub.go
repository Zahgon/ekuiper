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

package main

import (
	"os"
)

func subEventsFromMQTT(host string) { _ = "STUB: not implemented"; return }

// log.Infof("The connection to edgex messagebus is established successfully.")

func main() {
	if len(os.Args) == 3 {
		if v := os.Args[1]; v == "mqtt" {
			subEventsFromMQTT(os.Args[2])
		}
		if v := os.Args[1]; v == "redis" {
			panic("edgex v4 does not support redis")
		}
	}
}
