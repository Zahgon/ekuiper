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

// Runtime for symbol, to establish data connection

package runtime

import (
	"github.com/lf-edge/ekuiper/sdk/go/api"
	"github.com/lf-edge/ekuiper/sdk/go/connection"
)

type RuntimeInstance interface {
	run()
	stop() error
	isRunning() bool
}

func broadcast(ctx api.StreamContext, sock connection.DataOutChannel, data interface{}) {
	_ = "STUB: not implemented"
	// encode
	return
}

func parseContext(con *Control) (api.StreamContext, error) {
	_ = "STUB: not implemented"
	return *new(api.StreamContext), nil
}
