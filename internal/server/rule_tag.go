// Copyright 2025 EMQ Technologies Co., Ltd.
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

package server

import (
	"net/http"
)

type RuleTagRequest struct {
	Tags []string `json:"tags,omitempty"`
}

type RuleTagResponse struct {
	Rules []string `json:"rules,omitempty"`
}

func resetRuleTags(ruleJson string, newTags []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func updateRuleTags(ruleJson string, tags []string, addOrRemove bool) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func addNewTagsIntoExistTags(newTags []string, existTags []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func removeTagsFromExistTags(rTags []string, existTags []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func ruleTagHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func rulesTagsHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type BulkOperationResponse struct {
	RuleID  string `json:"ruleId"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func rulesBulkStartHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func rulesBulkStopHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func findRules(tags *RuleTagRequest) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
