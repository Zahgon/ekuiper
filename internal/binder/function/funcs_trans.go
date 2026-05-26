// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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

package function

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

type compressFunc struct {
	compressType string
	compressor   message.Compressor
}

func (c *compressFunc) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

// should never happen

func (c *compressFunc) Exec(ctx api.FunctionContext, args []any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (c *compressFunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }

type decompressFunc struct {
	compressType string
	decompressor message.Decompressor
}

func (d *decompressFunc) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

// should never happen

func (d *decompressFunc) Exec(ctx api.FunctionContext, args []any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (d *decompressFunc) IsAggregate() bool { _ = "STUB: not implemented"; return false }
