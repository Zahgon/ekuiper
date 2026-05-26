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

package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lf-edge/ekuiper/contract/v2/api"

	kctx "github.com/lf-edge/ekuiper/v2/internal/topo/context"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

const (
	ContentType     = "Content-Type"
	ContentTypeJSON = "application/json"
)

var (
	uploadDir       string
	uploadsDb       kv.KeyValue
	uploadsStatusDb kv.KeyValue
)

type statementDescriptor struct {
	Sql string `json:"sql,omitempty"`
}

func decodeStatementDescriptor(reader io.ReadCloser) (statementDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(statementDescriptor), nil
}

// Problems decoding

// Handle applies the specified error and error concept to the HTTP response writer
func handleError(w http.ResponseWriter, err error, prefix string, logger api.Logger) {
	_ = "STUB: not implemented"
	return
}

func packageInternalErrorCode(err error, msg string) string { _ = "STUB: not implemented"; return "" }

func jsonResponse(i interface{}, w http.ResponseWriter, logger api.Logger) {
	_ = "STUB: not implemented"
	return
}

// Problems encoding

func jsonByteResponse(buffer bytes.Buffer, w http.ResponseWriter, logger api.Logger) {
	_ = "STUB: not implemented"
	return
}

// Problems encoding

func traceMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

var router *mux.Router

func createRestServer(ip string, port int, needToken bool) *http.Server {
	_ = "STUB: not implemented"
	return nil
}

// r.HandleFunc("/connection/websocket", connectionHandler).Methods(http.MethodGet, http.MethodPost, http.MethodDelete)

// dump metrics

// Register extended routes

// Good practice to set timeouts to avoid Slowloris attacks.

type fileContent struct {
	Name     string `json:"name" yaml:"name"`
	Content  string `json:"content,omitempty" yaml:"content,omitempty"`
	FilePath string `json:"file,omitempty" yaml:"filePath,omitempty"`
}

func (f *fileContent) InstallScript() string { _ = "STUB: not implemented"; return "" }

func (f *fileContent) Validate() error { _ = "STUB: not implemented"; return nil }

func upload(file *fileContent) error { _ = "STUB: not implemented"; return nil }

func saveUploadFile(filename string, src io.Reader) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Create parent dir for the file if needed, recursively within root

// Split path and create each level

// Clean up the file if copy fails
// Close the file explicit before remove, otherwise it will fail on windows

func getFile(file *fileContent) error { _ = "STUB: not implemented"; return nil }

func explainRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// fetch the rule which will be explained

// resp := planner.BuildExplainResultFromLp(lp, 0)

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"

	// Upload or overwrite a file
	return
}

// Maximum upload of 1 GB files

// Get handler for filename, size and headers

// Get the list of files in the upload directory

func fileDeleteHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type information struct {
	Version       string `json:"version"`
	Os            string `json:"os"`
	Arch          string `json:"arch"`
	UpTimeSeconds int64  `json:"upTimeSeconds"`
	CpuUsage      string `json:"cpuUsage,omitempty"`
	MemoryUsed    string `json:"memoryUsed,omitempty"`
	MemoryTotal   string `json:"memoryTotal"`
}

func stopHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// The handler for root
func rootHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func pingHandler(w http.ResponseWriter, _ *http.Request) { _ = "STUB: not implemented"; return }

func sourceDetailsManageHandler(w http.ResponseWriter, r *http.Request, st ast.StreamType) {
	_ = "STUB: not implemented"
	return
}

func sourcesManageHandler(w http.ResponseWriter, r *http.Request, st ast.StreamType) {
	_ = "STUB: not implemented"
	return
}

func checkStreamBeforeDrop(name string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func sourceManageHandler(w http.ResponseWriter, r *http.Request, st ast.StreamType) {
	_ = "STUB: not implemented"
	return
}

// list or create streams
func streamDetailsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// list or create streams
func streamsHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// describe or delete a stream
func streamHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list or create streams
func tableDetailsHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// list or create tables
func tablesHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func tableHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func streamSchemaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func tableSchemaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func sourceSchemaHandler(w http.ResponseWriter, r *http.Request, st ast.StreamType) {
	_ = "STUB: not implemented"
	return
}

// list or create rules
func rulesHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// describe or delete a rule
func ruleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// delete rule will wait until rule close

func getAllRuleStatusHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// get status of a rule
func getStatusV2RulHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// get status of a rule
func getStatusRuleHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// start a rule
func startRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// stop a rule
func stopRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// restart a rule
func restartRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type EnableRuleTraceRequest struct {
	Strategy string `json:"strategy"`
}

func enableRuleTraceHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func disableRuleTraceHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func setIsRuleTraceEnabledHandler(name string, isEnabled bool, stra kctx.TraceStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

// get topo of a rule
func getTopoRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func ruleSchemaHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate a rule
func validateRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type rulesetInfo struct {
	Content  string `json:"content"`
	FilePath string `json:"file"`
}

func importHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func exportHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func testRuleHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func testRuleStartHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func testRuleStopHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func rulesTopCpuUsageHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func batchRequestHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

type EachRequest struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Body   string `json:"body"`
}

type EachResponse struct {
	Code     int    `json:"code"`
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}
