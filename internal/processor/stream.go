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

package processor

import (
	"bytes"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
)

var log = conf.Log

type StreamProcessor struct {
	db             kv.KeyValue
	streamStatusDb kv.KeyValue
	tableStatusDb  kv.KeyValue
	tempDb         kv.KeyValue
}

type StreamDetail struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Format string `json:"format"`
}

// globalStreamProcessor is the singleton instance of StreamProcessor
var globalStreamProcessor *StreamProcessor

func NewStreamProcessor() *StreamProcessor { _ = "STUB: not implemented"; return nil }

// GetDataSource retrieves a stream/table definition by name from both persistent and temp stores.
// It first checks the persistent store, then falls back to the temp store if not found.
func (p *StreamProcessor) GetDataSource(name string) (*ast.StreamStmt, error) {
	_ = "STUB: not implemented"
	// Try persistent store first
	return nil, nil
}

// Try temp store

// GetStreamProcessorDataSource is a global function that uses the global StreamProcessor instance
// to retrieve stream/table definitions from both persistent and temp stores.
func GetStreamProcessorDataSource(name string) (*ast.StreamStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *StreamProcessor) ExecStmt(statement string) (result []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Table is also StreamStmt

func (p *StreamProcessor) RecoverLookupTable() (err error) { _ = "STUB: not implemented"; return nil }

func (p *StreamProcessor) execSave(stmt *ast.StreamStmt, statement string, replace bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *StreamProcessor) ExecReplaceStream(name string, statement string, st ast.StreamType) (info string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// compare version

func (p *StreamProcessor) ExecStreamSql(statement string) (info string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *StreamProcessor) execShow(st ast.StreamType) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *StreamProcessor) ShowStream(st ast.StreamType) (res []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *StreamProcessor) ShowStreamOrTableDetails(kind string, st ast.StreamType) (res []StreamDetail, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *StreamProcessor) ShowTable(kind string) (res []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *StreamProcessor) GetStream(name string, st ast.StreamType) (res string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *StreamProcessor) execDescribe(stmt ast.NameNode, st ast.StreamType) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func printOptions(opts *ast.Options, buff *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (p *StreamProcessor) DescStream(name string, st ast.StreamType) (r ast.Statement, err error) {
	_ = "STUB: not implemented"
	return *new(ast.Statement), nil
}

func (p *StreamProcessor) GetInferredSchema(name string, st ast.StreamType) (r ast.StreamFields, err error) {
	_ = "STUB: not implemented"
	return *new(ast.StreamFields), nil
}

// GetInferredJsonSchema return schema in json schema type
// TODO merge external schema and inferred dynamic schema
func (p *StreamProcessor) GetInferredJsonSchema(name string, st ast.StreamType) (r map[string]*ast.JsonStreamField, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *StreamProcessor) execExplain(stmt ast.NameNode, st ast.StreamType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *StreamProcessor) execDrop(stmt ast.NameNode, st ast.StreamType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *StreamProcessor) DropStream(name string, st ast.StreamType) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if the key exists without unmarshalling content
// This allows deleting corrupted streams (e.g., from v1.x migration)

// Delete from temp db

// Delete from main db

func printFieldType(ft ast.FieldType) (result string) { _ = "STUB: not implemented"; return "" }

// GetAll return all streams and tables defined to export.
func (p *StreamProcessor) GetAll() (result map[string]map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DescribeToJson takes the human redable text from execDescribe and converts it to json
// Intended use is for CLI when passing -json flag
func DescribeToJson(s string) string { _ = "STUB: not implemented"; return "" }
