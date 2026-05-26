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

//go:build edgex || full

package edgex

import (
	"io"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type VaultSecret struct {
	scheme          string
	host            string
	port            int16
	secretName      string
	vaultToken      string
	renewalFactor   float64
	client          *http.Client
	authContext     map[string]interface{}
	logger          *logrus.Logger
	renewalCallback func()
	started         bool
	wg              sync.WaitGroup
}

const (
	DEFAULT_EDGEX_SERVICE_NAME = "rules-engine"
	DEFAULT_CREDENTIAL_FILE    = "/tmp/edgex/secrets/rules-engine/secrets-token.json"
	DEFAULT_VAULT_HOST         = "vault"
	DEFAULT_VAULT_PORT         = 8200
)

var vaultSecret *VaultSecret

func SecretProvider(logger *logrus.Logger) *VaultSecret { _ = "STUB: not implemented"; return nil }

func (v *VaultSecret) Start() { _ = "STUB: not implemented"; return }

func (v *VaultSecret) autoRenew() { _ = "STUB: not implemented"; return }

func (v *VaultSecret) renewToken() error { _ = "STUB: not implemented"; return nil }

func (v *VaultSecret) exchangeVaultToken() error { _ = "STUB: not implemented"; return nil }

// using the result, setup a callback schedule to keep the token fresh in case it's ever needed

func (v *VaultSecret) readToken(content []byte) error { _ = "STUB: not implemented"; return nil }

func (v *VaultSecret) callVault(req *http.Request) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (v *VaultSecret) Jwt() string { _ = "STUB: not implemented"; return "" }
