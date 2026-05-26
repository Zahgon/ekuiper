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

package util

import (
	"os"
	"path"

	kconf "github.com/lf-edge/ekuiper/tools/kubernetes/conf"
)

type (
	command struct {
		Url         string      `json:"url"`
		Description string      `json:"description"`
		Method      string      `json:"method"`
		Data        interface{} `json:"data"`
		strLog      string
	}
	fileData struct {
		Commands []*command `json:"commands"`
	}
)

func (c *command) getLog() string { _ = "STUB: not implemented"; return "" }

func (c *command) call(host string) bool { _ = "STUB: not implemented"; return false }

type (
	historyFile struct {
		Name     string `json:"name"`
		LoadTime int64  `json:"loadTime"`
	}
	server struct {
		dirCommand     string
		fileHistory    string
		mapHistoryFile map[string]*historyFile
		logs           []string
	}
)

func (f *historyFile) setName(name string) { _ = "STUB: not implemented"; return }

func (f *historyFile) setLoadTime(loadTime int64) { _ = "STUB: not implemented"; return }

func (s *server) getLogs() []string { _ = "STUB: not implemented"; return nil }

func (s *server) printLogs() { _ = "STUB: not implemented"; return }

func (s *server) loadHistoryFile() bool { _ = "STUB: not implemented"; return false }

func (s *server) init() bool {
	s.mapHistoryFile = make(map[string]*historyFile)
	conf := kconf.GetConf()
	dirCommand := conf.GetCommandDir()
	s.dirCommand = dirCommand
	s.fileHistory = path.Join(path.Dir(dirCommand), ".history")
	if _, err := os.Stat(s.fileHistory); os.IsNotExist(err) {
		if _, err = os.Create(s.fileHistory); nil != err {
			kconf.Log.Info(err)
			return false
		}
		return true
	}
	return s.loadHistoryFile()
}

func (s *server) saveHistoryFile() bool { _ = "STUB: not implemented"; return false }

func (s *server) isUpdate(entry os.DirEntry) bool { _ = "STUB: not implemented"; return false }

func (s *server) processDir() bool { _ = "STUB: not implemented"; return false }

func (s *server) watchFolders() { _ = "STUB: not implemented"; return }

func Process() { _ = "STUB: not implemented"; return }

func joinHostPortInt(host string, port int) string { _ = "STUB: not implemented"; return "" }
