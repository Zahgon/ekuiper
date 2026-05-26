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

package protobuf

import (
	"github.com/jhump/protoreflect/desc"            //nolint:staticcheck
	"github.com/jhump/protoreflect/desc/protoparse" //nolint:staticcheck
	"github.com/lf-edge/ekuiper/contract/v2/api"

	kconf "github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

type Converter struct {
	descriptor *desc.MessageDescriptor
	fc         *FieldConverter
}

var protoParser *protoparse.Parser

func init() {
	etcDir, _ := kconf.GetLoc("etc/schemas/protobuf/")
	dataDir, _ := kconf.GetLoc("data/schemas/protobuf/")
	protoParser = &protoparse.Parser{ImportPaths: []string{etcDir, dataDir}}
}

func NewConverter(schemaFile string, soFile string, messageName string) (message.Converter, error) {
	_ = "STUB: not implemented"
	return *new(message.Converter), nil
}

// collectProtoFiles returns a list of .proto file paths for the given path.
// If the path is a directory, it returns full paths for directory entries.
// If it is a single file, it returns the path as-is.
func collectProtoFiles(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Converter) Encode(ctx api.StreamContext, d any) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Converter) Decode(ctx api.StreamContext, b []byte) (m any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
