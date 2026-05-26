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

package cert

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

func GenTLSConfig(ctx api.StreamContext, props map[string]interface{}) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genTlsConfigurationOptions(props map[string]interface{}) (*model.TlsConfigurationOptions, *model.TlsKeys, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getTLSMinVersion(ctx api.StreamContext, userInput string) uint16 {
	_ = "STUB: not implemented"
	return 0
}

func getRenegotiationSupport(ctx api.StreamContext, userInput string) tls.RenegotiationSupport {
	_ = "STUB: not implemented"
	return *new(tls.RenegotiationSupport)
}

func isCertDefined(opts *model.TlsConfigurationOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func GenerateTLSForClient(ctx api.StreamContext, Opts *model.TlsConfigurationOptions, keys *model.TlsKeys) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildCert(ctx api.StreamContext, opts *model.TlsConfigurationOptions, keys *model.TlsKeys) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func certLoader(ctx api.StreamContext, certFilePath, keyFilePath string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func buildCA(ctx api.StreamContext, opts *model.TlsConfigurationOptions, tlsConfig *tls.Config, keys *model.TlsKeys) error {
	_ = "STUB: not implemented"
	return nil
}

func caLoader(ctx api.StreamContext, caFilePath string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
