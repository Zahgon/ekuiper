// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package schema

import (
	"context"

	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/modules"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

// Initialize in the server startup
var (
	registry       *Registry
	schemaDb       kv.KeyValue
	schemaStatusDb kv.KeyValue
)

// Registry is a global registry for schemas
// It stores the schema ids and the ref to its file content in memory
// The schema definition is stored in the file system and will only be loaded once used
type Registry struct {
	syncx.RWMutex
	// The map of schema files for all types
	schemas map[string]map[string]*modules.Files
}

// Registry provide the method to add, update, get and parse and delete schemas

// InitRegistry initialize the registry, only called once by the server
func InitRegistry() error { _ = "STUB: not implemented"; return nil }

// TODO shall we allow to delete etc schema?

// Read from etcDir firstly, then read from dataDir
// Compare version and leave the newer version

// merge schemas

func GetAllForType(schemaType string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Register(info *Info) error { _ = "STUB: not implemented"; return nil }

func CreateOrUpdateSchema(info *Info) error { _ = "STUB: not implemented"; return nil }

// compare version

// make sure info.Type does not escape from root

// If file path is a .zip, it must have the name.type file and a folder of the same name to hold the supporting files. Other files will all be ignored.
// Otherwise, save the file in the upper folder

// Check if it's the exact file we want

// Skip files that don't match our criteria

// clean up old ffs

func GetSchema(schemaType string, name string) (*Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSchemaFile return main schema file if schema id is defined; otherwise return the original schema id (possibly the file path)
func GetSchemaFile(schemaType string, name string) (*modules.Files, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteSchema(schemaType string, name string) error { _ = "STUB: not implemented"; return nil }

func doDelete(name string, schemaFile *modules.Files) error {
	_ = "STUB: not implemented"
	// If the schema is a folder, delete the folder otherwise delete the single file
	return nil
}

const BOOT_INSTALL = "$boot_install"

func GetAllSchema() map[string]string { _ = "STUB: not implemented"; return nil }

func GetAllSchemaStatus() map[string]string { _ = "STUB: not implemented"; return nil }

func UninstallAllSchema() { _ = "STUB: not implemented"; return }

func hasInstallFlag() bool { _ = "STUB: not implemented"; return false }

func clearInstallFlag() { _ = "STUB: not implemented"; return }

func ImportSchema(ctx context.Context, schema map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// set the flag to install the plugins when eKuiper reboot

// SchemaPartialImport compare the schema to be installed and the one in database
// if not exist in database, install;
// if existed, ignore
func SchemaPartialImport(ctx context.Context, schemas map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func schemaRegisterForImport(k, v string) error { _ = "STUB: not implemented"; return nil }

func schemaInstallWhenReboot() { _ = "STUB: not implemented"; return }

func storeSchemaInstallScript(info *Info) { _ = "STUB: not implemented"; return }

func removeSchemaInstallScript(schemaType string, name string) { _ = "STUB: not implemented"; return }

func GetSchemaInstallScript(key string) (string, string) { _ = "STUB: not implemented"; return "", "" }
