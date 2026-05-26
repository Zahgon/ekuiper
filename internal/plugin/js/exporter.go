// Copyright 2024 EMQ Technologies Co., Ltd.
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

package js

import (
	"context"
)

// Exporter is used to export and import the JavaScript functions
// The functions are stored in the key-value store only. So import and export are just to read and write the key-value store

// Import the JavaScript functions from the map. This is usually called after reset to override all settings
func (m *Manager) Import(ctx context.Context, scripts map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) PartialImport(ctx context.Context, scripts map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) Export() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) Status() map[string]string { _ = "STUB: not implemented"; return nil }

func (m *Manager) Reset() { _ = "STUB: not implemented"; return }
