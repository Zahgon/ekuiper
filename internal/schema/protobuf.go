//go:build schema || !core

package schema

import (
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"            //nolint:staticcheck
	"github.com/jhump/protoreflect/desc/protoparse" //nolint:staticcheck
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

var protoParser *protoparse.Parser

func init() {
	etcDir, _ := conf.GetLoc("etc/schemas/protobuf/")
	dataDir, _ := conf.GetLoc("data/schemas/protobuf/")
	protoParser = &protoparse.Parser{ImportPaths: []string{etcDir, dataDir}}
}

type PbType struct{}

func (p *PbType) Scan(logger api.Logger, schemaDir string) (map[string]*modules.Files, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subdirectory: treat as a single schema ID containing multiple .proto files

// SchemaFile points to the directory itself

func (p *PbType) Infer(_ api.Logger, filePath string, messageId string) (ast.StreamFields, error) {
	_ = "STUB: not implemented"
	return *new(ast.StreamFields), nil
}

// collectProtoFiles returns a list of .proto file paths for the given path.
// If the path is a directory, it returns dir-relative paths (e.g. "multidir/msg_a.proto").
// If it is a single file, it returns the path as-is.
func collectProtoFiles(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func convertMessage(m *desc.MessageDescriptor) (ast.StreamFields, error) {
	_ = "STUB: not implemented"
	return *new(ast.StreamFields), nil
}

func convertField(f *desc.FieldDescriptor) (ast.StreamField, error) {
	_ = "STUB: not implemented"
	return *new(ast.StreamField), nil
}

func convertFieldType(tt dpb.FieldDescriptorProto_Type, f *desc.FieldDescriptor) (ast.FieldType, error) {
	_ = "STUB: not implemented"
	return *new(ast.FieldType), nil
}

var _ modules.SchemaTypeDef = &PbType{}
