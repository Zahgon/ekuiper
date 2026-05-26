// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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

// INTECH Process Automation Ltd.
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

package conf

var LoadConfigCache map[string]map[string]interface{}

func init() {
	LoadConfigCache = make(map[string]map[string]interface{})
}

func clearLoadConfigCache() { _ = "STUB: not implemented"; return }

const Separator = "__"

func LoadConfig(c interface{}) error { _ = "STUB: not implemented"; return nil }

func LoadConfigByName(name string, c interface{}) error { _ = "STUB: not implemented"; return nil }

func LoadConfigFromPath(p string, c interface{}) error { _ = "STUB: not implemented"; return nil }

// Make all keys to lowercase to match environment variables then revert it back by checking json defs

// checking json keys

func CorrectsConfigKeysByJson(configs map[string]interface{}, jsonFilePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func getPrefix(p string) string { _ = "STUB: not implemented"; return "" }

func process(configMap map[string]interface{}, env map[string]string, prefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func handle(conf map[string]interface{}, keysLeft []string, val string) {
	_ = "STUB: not implemented"
	return
}

func trimPrefix(key string, prefix string) string { _ = "STUB: not implemented"; return "" }

func nameToKeys(key string) []string { _ = "STUB: not implemented"; return nil }

func getConfigKey(key string) string { _ = "STUB: not implemented"; return "" }

func getValueType(val string) interface{} { _ = "STUB: not implemented"; return nil }

func normalize(m map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func applyKeys(m map[string]interface{}, list []string) { _ = "STUB: not implemented"; return }

func applyKey(m map[string]interface{}, key string) { _ = "STUB: not implemented"; return }

func extractKeysFromJsonIfExists(yamlPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadJsonForYaml(filePath string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func jsonPathForFile(yamlPath string) string { _ = "STUB: not implemented"; return "" }

func extractNamesFromProperties(jsonMap map[string]interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractNamesFromElement(jsonMap map[string]interface{}) []string {
	_ = "STUB: not implemented"
	return nil
}

// If not a list/map, or an empty list/map, then it's a single element

func Printable(m map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}
