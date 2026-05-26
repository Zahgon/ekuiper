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

package timex

import (
	"os"
	"strings"
	"time"

	"github.com/benbjohnson/clock"
)

var (
	Clock     clock.Clock
	IsTesting bool
	Maxtime   = time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC)
)

func init() {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") {
			IsTesting = true
			break
		}
	}
	InitClock()
}

func InitClock() { _ = "STUB: not implemented"; return }

// GetTicker Time related. For Mock
func GetTicker(duration time.Duration) *clock.Ticker { _ = "STUB: not implemented"; return nil }

func After(duration time.Duration) <-chan time.Time { _ = "STUB: not implemented"; return nil }

func GetTimer(duration time.Duration) *clock.Timer { _ = "STUB: not implemented"; return nil }

func Sleep(duration time.Duration) { _ = "STUB: not implemented"; return }

func GetTimerByTime(t time.Time) *clock.Timer { _ = "STUB: not implemented"; return nil }

func GetNowInMilli() int64 { _ = "STUB: not implemented"; return 0 }

func GetNow() time.Time {
	_ = "STUB: not implemented"

	// Mock time, only use in test
	return *new(time.Time)
}

func Set(t int64) { _ = "STUB: not implemented"; return }

func SetNow(t time.Time) { _ = "STUB: not implemented"; return }

func Add(d time.Duration) { _ = "STUB: not implemented"; return }
