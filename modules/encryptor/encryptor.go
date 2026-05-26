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

package encryptor

import (
	"io"

	"github.com/lf-edge/ekuiper/v2/pkg/message"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

// GetEncryptor currently, decryptor and encryptor are the same instance
func GetEncryptor(name string, encryptProps map[string]any, conf *model.KuiperConf) (message.Encryptor, error) {
	_ = "STUB: not implemented"
	return *new(message.Encryptor), nil
}

func GetDecryptorWithKey(name string, key []byte, encryptProps map[string]any) (message.Decryptor, error) {
	_ = "STUB: not implemented"
	return *new(message.Decryptor), nil
}

func getAESKey(conf *model.KuiperConf) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func GetEncryptWriter(name string, output io.Writer, conf *model.KuiperConf) (io.Writer, error) {
	_ = "STUB: not implemented"
	// TODO support encryption props later
	return *new(io.Writer), nil
}
