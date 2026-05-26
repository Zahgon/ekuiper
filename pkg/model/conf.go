// Copyright 2025 EMQ Technologies Co., Ltd.
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

package model

import (
	"time"

	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
)

func (k *KuiperConf) PprofSameAsRest() bool { _ = "STUB: not implemented"; return false }

type KuiperConf struct {
	Hack struct {
		Cold bool `yaml:"cold"`
	}
	Basic struct {
		LogLevel                string                `yaml:"logLevel"`
		Debug                   bool                  `yaml:"debug"`
		ConsoleLog              bool                  `yaml:"consoleLog"`
		FileLog                 bool                  `yaml:"fileLog"`
		LogDisableTimestamp     bool                  `yaml:"logDisableTimestamp"`
		Syslog                  *SyslogConf           `yaml:"syslog"`
		RotateTime              int                   `yaml:"rotateTime"`
		MaxAge                  int                   `yaml:"maxAge"`
		RotateSize              int64                 `yaml:"rotateSize"`
		RotateCount             int                   `yaml:"rotateCount"`
		TimeZone                string                `yaml:"timezone"`
		Ip                      string                `yaml:"ip"`
		Port                    int                   `yaml:"port"`
		RestIp                  string                `yaml:"restIp"`
		RestPort                int                   `yaml:"restPort"`
		RestTls                 *TlsConf              `yaml:"restTls"`
		Prometheus              bool                  `yaml:"prometheus"`
		PrometheusPort          int                   `yaml:"prometheusPort"`
		Pprof                   bool                  `yaml:"pprof"`
		PprofIp                 string                `yaml:"pprofIp"`
		PprofPort               int                   `yaml:"pprofPort"`
		PluginHosts             string                `yaml:"pluginHosts"`
		Authentication          bool                  `yaml:"authentication"`
		IgnoreCase              bool                  `yaml:"ignoreCase"`
		SQLConf                 *SQLConf              `yaml:"sql"`
		RulePatrolInterval      cast.DurationConf     `yaml:"rulePatrolInterval"`
		EnableOpenZiti          bool                  `yaml:"enableOpenZiti"`
		AesKey                  string                `yaml:"aesKey"`
		GracefulShutdownTimeout cast.DurationConf     `yaml:"gracefulShutdownTimeout"`
		ResourceProfileConfig   ResourceProfileConfig `yaml:"ResourceProfileConfig"`
		MetricsDumpConfig       MetricsDumpConfig     `yaml:"metricsDumpConfig"`
		EnableRestAuditLog      bool                  `yaml:"enableRestAuditLog"`
		EnablePrivateNet        bool                  `yaml:"enablePrivateNet"`
		AllowExternalFileAccess bool                  `yaml:"allowExternalFileAccess"`
	}
	Rule   def.RuleOption
	Sink   *SinkConf
	Source *SourceConf
	Store  struct {
		Type         string `yaml:"type"`
		ExtStateType string `yaml:"extStateType"`
		Redis        struct {
			Host               string            `yaml:"host"`
			Port               int               `yaml:"port"`
			Password           string            `yaml:"password"`
			Timeout            cast.DurationConf `yaml:"timeout"`
			ConnectionSelector string            `yaml:"connectionSelector"`
		}
		Sqlite struct {
			Name string `yaml:"name"`
		}
		Fdb struct {
			Path string `yaml:"path"`
		}
		Pebble struct {
			Path string `yaml:"path"`
			Name string `yaml:"name"`
		}
	}
	Portable struct {
		PythonBin   string            `yaml:"pythonBin"`
		InitTimeout cast.DurationConf `yaml:"initTimeout"`
		SendTimeout time.Duration     `yaml:"sendTimeout"`
		RecvTimeout time.Duration     `yaml:"recvTimeout"`
	}
	Connection struct {
		BackoffMaxElapsedDuration cast.DurationConf `yaml:"backoffMaxElapsedDuration"`
	}
	OpenTelemetry OpenTelemetry `yaml:"openTelemetry"`
	AesKey        []byte
	Security      *SecurityConf
}

type TlsConf struct {
	Certfile string `yaml:"certfile"`
	Keyfile  string `yaml:"keyfile"`
}

type SinkConf struct {
	MemoryCacheThreshold int               `json:"memoryCacheThreshold" yaml:"memoryCacheThreshold"`
	MaxDiskCache         int               `json:"maxDiskCache" yaml:"maxDiskCache"`
	BufferPageSize       int               `json:"bufferPageSize" yaml:"bufferPageSize"`
	EnableCache          bool              `json:"enableCache" yaml:"enableCache"`
	ResendInterval       cast.DurationConf `json:"resendInterval" yaml:"resendInterval"`
	CleanCacheAtStop     bool              `json:"cleanCacheAtStop" yaml:"cleanCacheAtStop"`
	ResendAlterQueue     bool              `json:"resendAlterQueue" yaml:"resendAlterQueue"`
	ResendPriority       int               `json:"resendPriority" yaml:"resendPriority"`
	ResendIndicatorField string            `json:"resendIndicatorField" yaml:"resendIndicatorField"`
	ResendDestination    string            `json:"resendDestination" yaml:"resendDestination"`
}

// Validate the configuration and reset to the default value for invalid values.
func (sc *SinkConf) Validate(logger api.Logger) error { _ = "STUB: not implemented"; return nil }

type SourceConf struct {
	HttpServerIp   string   `json:"httpServerIp" yaml:"httpServerIp"`
	HttpServerPort int      `json:"httpServerPort" yaml:"httpServerPort"`
	HttpServerTls  *TlsConf `json:"httpServerTls" yaml:"httpServerTls"`
}

func (sc *SourceConf) Validate(logger api.Logger) error { _ = "STUB: not implemented"; return nil }

type SQLConf struct {
	MaxConnections int `yaml:"maxConnections"`
}

type SyslogConf struct {
	Enable  bool   `yaml:"enable"`
	Network string `yaml:"network"`
	Address string `yaml:"address"`
	Tag     string `yaml:"tag"`
	Level   string `yaml:"level"`
}

func (s *SyslogConf) Validate() error { _ = "STUB: not implemented"; return nil }

// valid, do nothing

type MetricsDumpConfig struct {
	Enable           bool          `yaml:"enable"`
	RetainedDuration time.Duration `yaml:"retainedDuration"`
}

type ResourceProfileConfig struct {
	Enable   bool          `yaml:"enable"`
	Interval time.Duration `yaml:"interval"`
}

type OpenTelemetry struct {
	ServiceName           string `yaml:"serviceName"`
	EnableRemoteCollector bool   `yaml:"enableRemoteCollector"`
	RemoteEndpoint        string `yaml:"remoteEndpoint"`
	LocalTraceCapacity    int    `yaml:"localTraceCapacity"`
	EnableLocalStorage    bool   `yaml:"enableLocalStorage"`
}
