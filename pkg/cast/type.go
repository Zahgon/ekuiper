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

package cast

import (
	"time"
)

type DurationConf time.Duration

func (dp *DurationConf) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (d DurationConf) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (dp *DurationConf) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DurationConf) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func ConvertDuration(s any) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// from json

type TypedNil struct{}

var TNil = (*TypedNil)(nil)
