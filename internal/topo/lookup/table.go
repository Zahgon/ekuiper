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

package lookup

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// Table is a lookup table runtime instance. It will run once the table is created.
// It will only stop once the table is dropped.

type info struct {
	ls    api.Source
	count int32
}

var (
	instances = make(map[string]*info)
	lock      = &syncx.Mutex{}
)

// Attach called by lookup nodes. Add a count to the info
func Attach(name string) (api.Source, error) {
	_ = "STUB: not implemented"
	return *new(api.Source), nil
}

// Detach called by lookup nodes when it is closed
func Detach(name string) error { _ = "STUB: not implemented"; return nil }

// CreateInstance called when create a lookup table
func CreateInstance(name string, sourceType string, options *ast.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the lookup source according to the source options

// TODO lookup table connection status support

// do nothing

// DropInstance called when drop a lookup table
func DropInstance(name string) error { _ = "STUB: not implemented"; return nil }
