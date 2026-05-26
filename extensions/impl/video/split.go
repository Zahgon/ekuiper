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

// splitJPEGs is a helper to find the start and end of JPEG images in a byte stream
func splitJPEGs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Find Start of Image (SOI): FF D8

// If we aren't at the start, skip until we find FF D8
// (This handles garbage data at the start)

// If no start found, skip everything except possibly the last byte (in case it is FF of start)
// Wait, if it might be FF D8 across chunks... strictly we should be careful.
// But for simplicity/robustness similar to user provided code:
// If we can't find FF D8, we skip everything.

// Found it at index, consume up to index

// Find End of Image (EOI): FF D9
// We look for the ending marker
// We start search from 2 because 0,1 are SOI

// Found it! The index is relative to data[2:], so actual index of FF is index+2.
// EOI is 2 bytes (FF D9).
// So total length is (index + 2) + 2 = index + 4

// Request more data
