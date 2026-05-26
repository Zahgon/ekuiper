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

//go:build edgex || full

package edgex

import (
	"github.com/sirupsen/logrus"
)

const (
	OPENZITI_LOG_FORMAT         = "openziti: %s"
	OPENZITI_DEFAULT_LOG_FORMAT = "default openziti: %s"
)

func adaptLogging(log *logrus.Logger) {
	_ = "STUB: not implemented"
	// with EdgeX enabled as of 2024, it includes OpenZiti support. OpenZiti uses the default
	// logrus logger. This quiets duplicative logging
	// Check if the hook is already added
	return
}

type LogrusAdaptor struct {
	lc *logrus.Logger
}

func (f *LogrusAdaptor) Format(entry *logrus.Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	// Implement your custom formatting logic here
	return nil, nil
}

func (f *LogrusAdaptor) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (f *LogrusAdaptor) Fire(e *logrus.Entry) error { _ = "STUB: not implemented"; return nil }
