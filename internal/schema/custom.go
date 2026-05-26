package schema

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
)

type CustomType struct{}

func (c *CustomType) Scan(logger api.Logger, schemaDir string) (map[string]*modules.Files, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CustomType) Infer(logger api.Logger, schemaId string, _ string) (ast.StreamFields, error) {
	_ = "STUB: not implemented"
	return *new(ast.StreamFields), nil
}

var _ modules.SchemaTypeDef = &CustomType{}
