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

package service

import (
	"github.com/jhump/protoreflect/desc"    //nolint:staticcheck
	"github.com/jhump/protoreflect/dynamic" //nolint:staticcheck
	"google.golang.org/protobuf/reflect/protoreflect"
)

type httpConnMeta struct {
	Method string
	Uri    string // The Uri is a relative path which must start with /
	Body   []byte
}

type httpMapping interface {
	ConvertHttpMapping(method string, params []interface{}) (*httpConnMeta, error)
}

const (
	httpAPI      = "google.api.http"
	wildcardBody = "*"
	emptyBody    = ""
)

type httpOptions struct {
	Method      string
	UriTemplate *uriTempalte // must not nil
	BodyField   string
}

type uriTempalte struct {
	Template string
	Fields   []*field
}

type field struct {
	name   string
	prefix string
}

func (d *wrappedSchemalessDescriptor) ConvertHttpMapping(_ string, params []interface{}) (*httpConnMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *wrappedProtoDescriptor) parseHttpOptions() error { _ = "STUB: not implemented"; return nil }

// Find http option and exit loop at once. If not found, http option is nil

func (d *wrappedProtoDescriptor) ConvertHttpMapping(method string, params []interface{}) (*httpConnMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove all params to be used in the params, the left params are for BODY

// If options are not set, use the default setting

func getMessageFieldWithDots(message *dynamic.Message, name string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getUriOpt(d protoreflect.FieldDescriptor, v protoreflect.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ho *httpOptions) convertUri(md *desc.MethodDescriptor, uriOpt string, bodyOpt string) error {
	_ = "STUB: not implemented"
	return nil
	// the value represents if the key is still available (not used) so that they can be removed from *
}

func (u *uriTempalte) updateUriParams(md *desc.MessageDescriptor, prefix string, fmap map[string]bool, paramAdded bool) bool {
	_ = "STUB: not implemented"
	return false
}

// The first level field which are not consumed or the second level field
