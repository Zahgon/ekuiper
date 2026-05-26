// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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

package schedule

import (
	"time"
)

const layout = "2006-01-02 15:04:05"

type DatetimeRange struct {
	Begin          string `json:"begin" yaml:"begin"`
	End            string `json:"end" yaml:"end"`
	BeginTimestamp int64  `json:"beginTimestamp" yaml:"beginTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp" yaml:"endTimestamp"`
}

func IsInScheduleRanges(now time.Time, timeRanges []DatetimeRange) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isInScheduleRange(now time.Time, start string, end string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isInScheduleRangeByTS(now time.Time, startTS int64, endTS int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isInTimeRange(now time.Time, start string, end string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func IsAfterTimeRanges(now time.Time, ranges []DatetimeRange) bool {
	_ = "STUB: not implemented"
	return false
}

func isAfterTime(now time.Time, compare time.Time) bool { _ = "STUB: not implemented"; return false }

func isAfterTimeByTS(now time.Time, end int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isAfterTimeRange(now time.Time, end string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsInRunningSchedule checks whether the rule should be running, eg:
// If the duration is 10min, and cron is "0 0 * * *", and the current time is 00:00:02
// And the rule should be started immediately instead of checking it on the next day.
func IsInRunningSchedule(cronExpr string, now time.Time, d time.Duration) (bool, time.Duration, error) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration), nil
}

func ValidateRanges(ranges []DatetimeRange) error { _ = "STUB: not implemented"; return nil }

func validateRange(r DatetimeRange) error { _ = "STUB: not implemented"; return nil }
