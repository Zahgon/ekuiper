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

package ast

const (
	TypeStream StreamType = iota
	TypeTable
)

var StreamTypeMap = map[StreamType]string{
	TypeStream: "stream",
	TypeTable:  "table",
}

const (
	StreamKindLookup = "lookup"
	StreamKindScan   = "scan"
)

type StreamType int

type StreamStmt struct {
	Name         StreamName
	StreamFields StreamFields
	Options      *Options
	StreamType   StreamType // default to TypeStream

	Statement
}

type StreamField struct {
	Name string
	FieldType
}

type JsonStreamField struct {
	Type       string                      `json:"type,omitempty"`
	Items      *JsonStreamField            `json:"items,omitempty"`
	Properties map[string]*JsonStreamField `json:"properties,omitempty"`
	HasIndex   bool                        `json:"hasIndex,omitempty"`
	Index      int                         `json:"index"`

	Selected bool `json:"selected,omitempty"`
}

func (u *StreamField) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON The json format follows json schema
func (sf *StreamFields) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (sf *StreamFields) UnmarshalFromMap(data map[string]*JsonStreamField) error {
	_ = "STUB: not implemented"
	return nil
}

func (sf *StreamFields) ToJsonSchema() map[string]*JsonStreamField {
	_ = "STUB: not implemented"
	return nil
}

func convertSchema(sfs StreamFields) map[string]*JsonStreamField {
	_ = "STUB: not implemented"
	return nil
}

func convertFieldType(sf FieldType) *JsonStreamField { _ = "STUB: not implemented"; return nil }

func fieldsTypeFromSchema(mjsf map[string]*JsonStreamField) (StreamFields, error) {
	_ = "STUB: not implemented"
	return *new(StreamFields), nil
}

func fieldTypeFromSchema(v *JsonStreamField) (FieldType, error) {
	_ = "STUB: not implemented"
	return *new(FieldType), nil
}

type StreamFields []StreamField

type FieldType interface {
	fieldType()
}

type BasicType struct {
	Type DataType
	FieldType
}

type ArrayType struct {
	Type DataType
	FieldType
}

type RecType struct {
	StreamFields StreamFields
	FieldType
}

// Options The stream AST tree
type Options struct {
	DATASOURCE        string `json:"datasource,omitempty"`
	KEY               string `json:"key,omitempty"`
	FORMAT            string `json:"format,omitempty"`
	CONF_KEY          string `json:"confKey,omitempty"`
	TYPE              string `json:"type,omitempty"`
	STRICT_VALIDATION bool   `json:"strictValidation,omitempty"`
	TIMESTAMP         string `json:"timestamp,omitempty"`
	TIMESTAMP_FORMAT  string `json:"timestampFormat,omitempty"`
	SHARED            bool   `json:"shared,omitempty"`
	SCHEMAID          string `json:"schemaid,omitempty"`
	VERSION           string `json:"version,omitempty"`
	EXTRA             string `json:"extra,omitempty"`
	Temp              bool   `json:"temp,omitempty"`
	// for scan table only
	RETAIN_SIZE int `json:"retainSize,omitempty"`
	// for table only, to distinguish lookup & scan
	KIND string `json:"kind,omitempty"`
	// for delimited format only
	DELIMITER string `json:"delimiter,omitempty"`

	RuleID       string                      `json:"-"`
	Schema       map[string]*JsonStreamField `json:"-"`
	IsWildCard   bool                        `json:"-"`
	IsSchemaLess bool                        `json:"-"`
	StreamName   string                      `json:"-"`
}

func (o Options) node() { _ = "STUB: not implemented"; return }

type ShowStreamsStatement struct {
	Statement
}

type DescribeStreamStatement struct {
	Name string

	Statement
}

type ExplainStreamStatement struct {
	Name string

	Statement
}

type DropStreamStatement struct {
	Name string

	Statement
}

func (dss *DescribeStreamStatement) GetName() string { _ = "STUB: not implemented"; return "" }

func (ess *ExplainStreamStatement) GetName() string { _ = "STUB: not implemented"; return "" }

func (dss *DropStreamStatement) GetName() string { _ = "STUB: not implemented"; return "" }

type ShowTablesStatement struct {
	Statement
}

type DescribeTableStatement struct {
	Name string

	Statement
}

type ExplainTableStatement struct {
	Name string

	Statement
}

type DropTableStatement struct {
	Name string

	Statement
}

func (dss *DescribeTableStatement) GetName() string { _ = "STUB: not implemented"; return "" }
func (ess *ExplainTableStatement) GetName() string  { _ = "STUB: not implemented"; return "" }
func (dss *DropTableStatement) GetName() string     { _ = "STUB: not implemented"; return "" }

func printFieldTypeForJson(ft FieldType) (result interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func doPrintFieldTypeForJson(ft FieldType) (result string, isLiteral bool) {
	_ = "STUB: not implemented"
	return "", false
}

func CheckSchemaIndex(schema map[string]*JsonStreamField) bool {
	_ = "STUB: not implemented"
	return false
}
