// Copyright 2023-2025 EMQ Technologies Co., Ltd.
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

package http

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/sirupsen/logrus"

	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/message"
)

// ClientConf is the configuration for http client
// It is shared by httppull source and rest sink to configure their http client
type ClientConf struct {
	config       *RawConf
	client       *http.Client
	decompressor message.Decompressor // decompressor used to payload decompression when specifies compressAlgorithm

	// auth related, all handled inside client sync
	// In Conn, try to auth
	// In Send, do refreshing
	accessConf  *AccessTokenConf
	refreshConf *RefreshTokenConf

	tokenLastUpdateAt   time.Time
	parsedHeaders       map[string]string
	parsedBody          string
	parsedRefreshHeader map[string]string
	parsedRefreshBody   string
}

type AccessTokenConf struct {
	Url            string            `json:"url"`
	Body           string            `json:"body"`
	Expire         string            `json:"expire"`
	Headers        map[string]string `json:"headers"`
	ExpireInSecond int
}

type RefreshTokenConf struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type RawConf struct {
	Url           string            `json:"url"`
	Method        string            `json:"method"`
	Body          string            `json:"body"`
	BodyType      string            `json:"bodyType"`
	Format        string            `json:"format"`
	Headers       map[string]string `json:"headers"`
	FormData      map[string]string `json:"formData"`
	FileFieldName string            `json:"fileFieldName"`
	Timeout       cast.DurationConf `json:"timeout"`
	Incremental   bool              `json:"incremental"`

	OAuth      map[string]map[string]interface{} `json:"oauth"`
	SendSingle bool                              `json:"sendSingle"`
	// Could be code or body
	ResponseType string `json:"responseType"`
	Compression  string `json:"compression"` // Compression specifies the algorithms used to payload compression

	DebugResp bool `json:"debugResp"`
}

const (
	DefaultTimeout = 5000 * time.Millisecond
)

type bodyResp struct {
	Code int `json:"code"`
}

var bodyTypeMap = map[string]string{"none": "", "text": "text/plain", "json": "application/json", "html": "text/html", "xml": "application/xml", "javascript": "application/javascript", "form": "", "binary": "application/octet-stream", "formdata": "multipart/form-data"}

// newTransport allows EdgeX Foundry, protected by OpenZiti to override and obtain a transport
// protected by OpenZiti's zero trust connectivity. See client_edgex.go where this function is
// set in an init() call
var newTransport = getTransport

func getTransport(tlscfg *tls.Config, logger *logrus.Logger) *http.Transport {
	_ = "STUB: not implemented"
	return nil
}

func (cc *ClientConf) InitConf(ctx api.StreamContext, device string, props map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Set default body type if not set

// correct

// validate oAuth. In order to adapt to manager, the validation is closed to allow empty value

// validate access token

// validate refresh token, it is optional

// that means payload need compression and decompression, so we need initialize compressor and decompressor

func (cc *ClientConf) Conn(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConf) Send(ctx api.StreamContext, bodyType string, method string, u string, headers map[string]string, formData map[string]string, formFieldName string, v any) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check token refresh after send

// initialize the oAuth access token
func (cc *ClientConf) auth(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func parseHeaders(ctx api.StreamContext, oHeaders map[string]string, data map[string]interface{}) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *ClientConf) refresh(ctx api.StreamContext) error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConf) updateToken(ctx api.StreamContext, tk map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	BODY_ERR = "response body error"
	CODE_ERR = "response code error"
)

// responseBodyDecompress used to decompress the specified response body bytes, decompression algorithm indicated
// by response header 'Content-Encoding' value.
func (cc *ClientConf) responseBodyDecompress(ctx api.StreamContext, resp *http.Response, body []byte) ([]byte, error) {
	_ = "STUB: not implemented"

	// we need check response header key Content-Encoding is exist, if not that means remote server probably not support
	// configured compression algorithm and we should throw error.
	return nil, nil
}

// parse the response status. For rest sink, it will not return the body by default if not need to debug
func (cc *ClientConf) parseResponse(ctx api.StreamContext, resp *http.Response, lastMD5 string, returnBody bool, skipDecompression bool) ([]map[string]interface{}, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// For rest sink who only need to know if the request is successful

//{"code":0,"message":"success","data":null}

func (cc *ClientConf) parseHeaders(ctx api.StreamContext, data map[string]interface{}) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decode(data []byte) ([]map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMD5Hash(text []byte) string { _ = "STUB: not implemented"; return "" }
