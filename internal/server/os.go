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

package server

const (
	EtcOsRelease    string = "/etc/os-release"
	UsrLibOsRelease string = "/usr/lib/os-release"
)

// Read and return os-release, trying EtcOsRelease, followed by UsrLibOsRelease.
// err will contain an error message if neither file exists or failed to parse
func Read() (osrelease map[string]string, err error) { _ = "STUB: not implemented"; return nil, nil }

// Similar to Read(), but takes the name of a file to load instead
func ReadFile(filename string) (osrelease map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadString is similar to Read(), but takes a string to load instead
func ReadString(content string) (osrelease map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFile(filename string) (lines []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseString(content string) (lines []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseLine(line string) (key string, value string, err error) {
	_ = "STUB: not implemented"

	// skip empty lines
	return "", "", nil
}

// skip comments

// try to split string at the first '='

// trim white space from key and value

// Handle double quotes

// expand anything else that could be escaped
