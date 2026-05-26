// Copyright 2025 EMQ Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/utahta/go-cronowriter"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

const OpenMetricsEOF = "# EOF\n"

func InitMetricsDumpJob(ctx context.Context) { _ = "STUB: not implemented"; return }

func GetMetricsZipFile(startTime time.Time, endTime time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func IsMetricsDumpEnabled() bool { _ = "STUB: not implemented"; return false }

func StartMetricsManager() error { _ = "STUB: not implemented"; return nil }

func StopMetricsManager() { _ = "STUB: not implemented"; return }

var metricsManager = &MetricsDumpManager{}

type MetricsDumpManager struct {
	syncx.Mutex
	enabeld          bool
	writer           *cronowriter.CronoWriter
	metricsPath      string
	retainedDuration time.Duration
	regex            *regexp.Regexp
	cancel           context.CancelFunc
	wg               *sync.WaitGroup
	dryRun           bool
}

func (m *MetricsDumpManager) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m *MetricsDumpManager) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (m *MetricsDumpManager) Stop() { _ = "STUB: not implemented"; return }

func (m *MetricsDumpManager) Start() error { _ = "STUB: not implemented"; return nil }

func (m *MetricsDumpManager) init(parCtx context.Context) error {
	if err := conf.InitMetricsFolder(); err != nil {
		return fmt.Errorf("init metrics folder err:%v", err)
	}
	ctx, cancel := context.WithCancel(parCtx)
	m.cancel = cancel
	m.wg = &sync.WaitGroup{}
	m.enabeld = true
	metricsPath, err := conf.GetMetricsLoc()
	if err != nil {
		return err
	}
	m.metricsPath = metricsPath
	w := cronowriter.MustNew(fmt.Sprintf("%s/metrics.", m.metricsPath) + `%Y%m%d-%H` + `.log`)
	m.writer = w
	m.retainedDuration = conf.Config.Basic.MetricsDumpConfig.RetainedDuration
	m.regex = regexp.MustCompile(`^metrics\.(\d{4})(\d{2})(\d{2})-(\d{2})\.log$`)
	m.wg.Add(2)
	go m.gcOldMetricsJob(ctx)
	go m.dumpMetricsJob(ctx)
	conf.Log.Infof("metrics dump enabled, folder:%v, retension:%v", m.metricsPath, m.retainedDuration.String())
	return nil
}

func (m *MetricsDumpManager) gcOldMetricsJob(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (m *MetricsDumpManager) gcOldMetrics() error { _ = "STUB: not implemented"; return nil }

func (m *MetricsDumpManager) needGCFile(filename string, gcTime time.Time) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *MetricsDumpManager) dumpMetricsJob(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *MetricsDumpManager) dumpMetrics() error { _ = "STUB: not implemented"; return nil }

func (m *MetricsDumpManager) dumpMetricsFile(startTime time.Time, endTime time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *MetricsDumpManager) dumpMetricsFileIntoZip(filenames []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *MetricsDumpManager) writeOpenMetricsIntoFile(filenames []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *MetricsDumpManager) extractFileTime(fileName string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func isFileIncludeMetricsTime(fileTime, metricsTime time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func stringToPtr(a string) *string { _ = "STUB: not implemented"; return nil }
