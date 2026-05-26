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

package httpx

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
)

var BodyTypeMap = map[string]string{"none": "", "text": "text/plain", "json": "application/json", "html": "text/html", "xml": "application/xml", "javascript": "application/javascript", "form": "application/x-www-form-urlencoded;param=value"}

// Send v must be a []byte or map
func Send(logger api.Logger, client *http.Client, bodyType string, method string, u string, headers map[string]string, v any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SendWithFormData(logger api.Logger, client *http.Client, bodyType string, method string, u string, headers map[string]string, formData map[string]string, formFieldName string, v any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IsValidUrl(uri string) bool { _ = "STUB: not implemented"; return false }

// ReadFile Need to close the return reader
func ReadFile(uri string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// deal with windows path

// When external file access is not allowed, restrict to data/uploads dir

// Check if path is under uploads directory

// Use OpenRoot for sandboxed file access

// Get the data

func GetSSRFDialContext(timeout time.Duration) func(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

func DownloadFile(folder string, name string, uri string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Write the body to file

func IsHttpUrl(str string) error { _ = "STUB: not implemented"; return nil }
