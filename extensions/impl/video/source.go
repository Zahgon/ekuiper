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

package video

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/cast"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

type Source struct {
	Url string `json:"url"`
	// Check https://www.ffmpeg.org/general.html#Video-Codecs, default to 'mjpeg'
	Codec     string            `json:"codec"`
	DebugResp bool              `json:"debugResp"`
	InputArgs map[string]any    `json:"inputArgs"`
	Interval  cast.DurationConf `json:"interval"`
	meta      map[string]any
}

func (s *Source) Provision(ctx api.StreamContext, props map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Source) Close(_ api.StreamContext) error {
	_ = "STUB: not implemented"
	// do nothing
	return nil
}

func (s *Source) Connect(_ api.StreamContext, sch api.StatusChangeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Source) Subscribe(ctx api.StreamContext, ingest api.BytesIngest, ingestError api.ErrorIngest) error {
	_ = "STUB: not implemented"
	return nil
}

// check if recoverable
// If process exit too fast (less than 2s) with specific error, return error

// backoff

func isFatalError(err error) bool { _ = "STUB: not implemented"; return false }

func (s *Source) runCurrent(ctx api.StreamContext, fps string, ingest api.BytesIngest) error {
	_ = "STUB: not implemented"
	return nil
}

// Goroutine 1: wait for process exit

// Goroutine 2: monitor context cancellation to unblock pipe scanners.
// We must kill the process immediately on cancellation, otherwise the scanner.Scan()
// loops below will block indefinitely waiting for data that will never come.

// We must read stderr even if DebugResp is false to prevent the ffmpeg pipe from
// blocking and to capture the last error message for diagnostics.

// Larger buffer for high res images

// Copy data because scanner bytes are reused

// Wait for process completion or cancellation.
// If context was cancelled, Goroutine 2 has already killed the process,
// which unblocked the scanners above.

func (s *Source) Info() model.NodeInfo { _ = "STUB: not implemented"; return *new(model.NodeInfo) }

func (s *Source) TransformType() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

func GetSource() api.Source { _ = "STUB: not implemented"; return *new(api.Source) }

var (
	_ api.BytesSource = &Source{}
	_ model.InfoNode  = &Source{}
)
