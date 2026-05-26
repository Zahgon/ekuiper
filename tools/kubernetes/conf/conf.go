// Copyright 2021 EMQ Technologies Co., Ltd.
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

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type (
	config struct {
		Port         int    `yaml:"port"`
		Timeout      int    `yaml:"timeout"`
		IntervalTime int    `yaml:"intervalTime"`
		Ip           string `yaml:"ip"`
		ConsoleLog   bool   `yaml:"consoleLog"`
		FileLog      bool   `yaml:"fileLog"`
		LogPath      string `yaml:"logPath"`
		CommandDir   string `yaml:"commandDir"`
	}
)

var gConf config

func GetConf() *config { _ = "STUB: not implemented"; return nil }

func (c *config) GetIntervalTime() int { _ = "STUB: not implemented"; return 0 }

func (c *config) GetIp() string { _ = "STUB: not implemented"; return "" }

func (c *config) GetPort() int { _ = "STUB: not implemented"; return 0 }

func (c *config) GetLogPath() string { _ = "STUB: not implemented"; return "" }

func (c *config) GetCommandDir() string { _ = "STUB: not implemented"; return "" }

func processPath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *config) initConfig() bool { _ = "STUB: not implemented"; return false }

var (
	Log     *logrus.Logger
	gClient http.Client
)

func (c *config) initTimeout() { _ = "STUB: not implemented"; return }

func (c *config) initLog() bool { _ = "STUB: not implemented"; return false }

func (c *config) Init() bool { _ = "STUB: not implemented"; return false }

func fetchContents(request *http.Request) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
	if respon.StatusCode < 200 || respon.StatusCode > 299 {
		return data, fmt.Errorf("http return code: %d and error message %s.", respon.StatusCode, string(data))
	}
*/

func Get(inUrl string) (data []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func Post(inHead, inBody string) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Put(inHead, inBody string) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Delete(inUrl string) (data []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func LoadFileUnmarshal(path string, ret interface{}) error { _ = "STUB: not implemented"; return nil }

func SaveFileMarshal(path string, content interface{}) error { _ = "STUB: not implemented"; return nil }
