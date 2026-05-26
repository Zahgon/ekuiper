// Copyright 2023-2025 EMQ Technologies Co., Ltd.
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
	"time"

	"github.com/lf-edge/ekuiper/v2/internal/conf/logger"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

const (
	ConfFileName  = "kuiper.yaml"
	DebugLogLevel = "debug"
	InfoLogLevel  = "info"
	WarnLogLevel  = "warn"
	ErrorLogLevel = "error"
	FatalLogLevel = "fatal"
	PanicLogLevel = "panic"
)

var (
	Config    *model.KuiperConf
	IsTesting bool
	TestId    string
)

func InitConf() { _ = "STUB: not implemented"; return }

// 5 minutes

// Init when env is set OR enable is true

func SetLogLevel(level string, debug bool) { _ = "STUB: not implemented"; return }

func SetConsoleAndFileLog(consoleLog, fileLog bool) error { _ = "STUB: not implemented"; return nil }

// gc outdated log files by logrus itself

func SetLogFormat(disableTimestamp bool) { _ = "STUB: not implemented"; return }

func ValidateRuleOption(option *def.RuleOption) error { _ = "STUB: not implemented"; return nil }

func init() {
	logger.Log.Debugf("conf init")
	IsTesting = logger.IsTesting
}

func gcOutdatedLog(filePath string, maxDuration time.Duration) { _ = "STUB: not implemented"; return }

func isLogOutdated(name string, now time.Time, maxDuration time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}
