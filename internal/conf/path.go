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

package conf

func init() {
	PathConfig.LoadFileType = "relative"
	PathConfig.Dirs = AbsoluteMapping
}

type PathConfigure struct {
	LoadFileType string
	Dirs         map[string]string
}

const (
	etcDir        = "etc"
	dataDir       = "data"
	logDir        = "log"
	pluginsDir    = "plugins"
	metricsDir    = "metrics"
	KuiperBaseKey = "KuiperBaseKey"
)

var (
	PathConfig      PathConfigure
	AbsoluteMapping = map[string]string{
		etcDir:     "/etc/kuiper",
		dataDir:    "/var/lib/kuiper/data",
		logDir:     "/var/log/kuiper",
		pluginsDir: "/var/lib/kuiper/plugins",
	}
)

func GetConfLoc() (s string, err error) { _ = "STUB: not implemented"; return "", nil }

func GetLogLoc() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetMetricsLoc() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetDataLoc() (s string, err error) { _ = "STUB: not implemented"; return "", nil }

func GetPluginsLoc() (s string, err error) { _ = "STUB: not implemented"; return "", nil }

func absolutePath(loc string) (dir string, err error) { _ = "STUB: not implemented"; return "", nil }

// GetLoc subdir must be a relative path
func GetLoc(subdir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func relativePath(subdir string) (dir string, err error) { _ = "STUB: not implemented"; return "", nil }

// Log.Printf("Trying to load file from %s", confDir)

// Log.Printf("Trying to load file from %s", confDir)

func InitMetricsFolder() error { _ = "STUB: not implemented"; return nil }
