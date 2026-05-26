// Copyright 2024-2025 EMQ Technologies Co., Ltd.
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

package aes

import (
	"crypto/cipher"
)

type GcmEncrypter struct {
	gcm           cipher.AEAD
	constantNonce []byte
	aad           []byte
	tagSize       int
}

func NewGcmEncrypter(key []byte, cc *c) (*GcmEncrypter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *GcmEncrypter) Encrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *GcmEncrypter) Decrypt(secret []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
