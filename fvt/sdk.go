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

package fvt

import (
	"net/http"
	"net/url"
	"time"

	"github.com/lf-edge/ekuiper/v2/internal/server"
)

const ContentTypeJson = "application/json"

type SDK struct {
	baseUrl    *url.URL
	httpClient *http.Client
}

func NewSdk(baseUrl string) (*SDK, error) { _ = "STUB: not implemented"; return nil, nil }

func (sdk *SDK) Get(command string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) Post(command string, body string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) Import(content string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) PostWithParam(command string, param string, body string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) Req(command string, method string, body string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) Delete(command string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) CreateStream(streamJson string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) CreateTable(tableJson string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) DeleteStream(name string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) DeleteTables(name string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) GetStreamSchema(name string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) GetRuleSchema(name string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) CreateRule(ruleJson string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) RestartRule(ruleId string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) StartRule(ruleId string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) StopRule(ruleId string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) UpdateRule(name, ruleJson string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) DeleteRule(name string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) GetRuleStatus(name string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetResponseText(resp *http.Response) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetResponseResultMap(resp *http.Response) (result map[string]any, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetResponseResultTextAndMap(resp *http.Response) (body []byte, result map[string]any, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (sdk *SDK) CreateConf(confpath string, conf map[string]any) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) BatchRequest(reqs []*server.EachRequest) ([]*server.EachResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) ResetRuleTags(name string, tags []string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) AddRuleTags(name string, tags []string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) RemoveRuleTags(name string, keys []string) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sdk *SDK) GetRulesByTags(tags []string) (list []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TryAssert(count int, interval time.Duration, tryFunc func() bool) bool {
	_ = "STUB: not implemented"
	return false
}
