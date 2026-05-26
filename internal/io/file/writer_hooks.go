// Copyright 2023 EMQ Technologies Co., Ltd.
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

package file

type writerHooks interface {
	Header() []byte
	Line() []byte
	Footer() []byte
}

type jsonWriterHooks struct{}

func (j *jsonWriterHooks) Header() []byte { _ = "STUB: not implemented"; return nil }

func (j *jsonWriterHooks) Line() []byte { _ = "STUB: not implemented"; return nil }

func (j *jsonWriterHooks) Footer() []byte { _ = "STUB: not implemented"; return nil }

var jsonHooks = &jsonWriterHooks{}

type linesWriterHooks struct{}

func (l *linesWriterHooks) Header() []byte { _ = "STUB: not implemented"; return nil }

func (l *linesWriterHooks) Line() []byte { _ = "STUB: not implemented"; return nil }

func (l *linesWriterHooks) Footer() []byte { _ = "STUB: not implemented"; return nil }

var linesHooks = &linesWriterHooks{}

type csvWriterHooks struct {
	header []byte
}

func (c *csvWriterHooks) Header() []byte { _ = "STUB: not implemented"; return nil }

func (c *csvWriterHooks) Line() []byte { _ = "STUB: not implemented"; return nil }

func (c *csvWriterHooks) Footer() []byte { _ = "STUB: not implemented"; return nil }

func (c *csvWriterHooks) SetHeader(header string) { _ = "STUB: not implemented"; return }
