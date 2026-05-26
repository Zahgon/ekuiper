// Copyright 2022-2025 EMQ Technologies Co., Ltd.
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

import (
	"net"

	"github.com/sirupsen/logrus"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/internal/pkg/store"
	"github.com/lf-edge/ekuiper/v2/internal/processor"
	"github.com/lf-edge/ekuiper/v2/pkg/model"
)

var (
	logger                 = conf.Log
	startTimeStamp         int64
	version                = ""
	sysMetrics             *Metrics
	ruleProcessor          *processor.RuleProcessor
	streamProcessor        *processor.StreamProcessor
	rulesetProcessor       *processor.RulesetProcessor
	ruleMigrationProcessor *RuleMigrationProcessor
	stopSignal             chan struct{}
	cpuProfiler            = &ekuiperProfile{}
)

// newNetListener allows EdgeX Foundry, protected by OpenZiti to override and obtain a transport
// protected by OpenZiti's zero trust connectivity. See client_edgex.go where this function is
// set in an init() call
var newNetListener = newTcpListener

func newTcpListener(addr string, logger *logrus.Logger) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func stopEKuiper() { _ = "STUB: not implemented"; return }

// Create path if mount an empty dir. For edgeX, all the folders must be created priorly
func createPaths() { _ = "STUB: not implemented"; return }

// Create dir if not exist

// Create dir if not exist

func getStoreConfigByKuiperConfig(c *model.KuiperConf) (*store.StoreConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func canSetupCheckpointDB() (setup bool) { _ = "STUB: not implemented"; return false }

// checkpoint.db already exists, setup it.

// sqliteKV.db not exists, setup checkpoint.db.

// unexpected error happened

// sqliteKV.db exists/ checkpoint.db not exists, don't setup.

func StartUp(Version string) { _ = "STUB: not implemented"; return }

// Print inited modules

// register all extensions

// Bind the source, function, sink

// Start lookup tables

// Start rules

// Start rest service

// Start extend services

// Register conf managers

// Startup message

// Stop the services

// sleep 1 sec in order to let stop request got response

// wait rule checker exit

// kill all plugin process

// close extend services
