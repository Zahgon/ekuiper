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

package def

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/schedule"
	"github.com/lf-edge/ekuiper/v2/pkg/cast"
)

type RuleOption struct {
	Debug                     bool                     `json:"debug" yaml:"debug"`
	LogFilename               string                   `json:"logFilename,omitempty" yaml:"logFilename,omitempty"`
	IsEventTime               bool                     `json:"isEventTime" yaml:"isEventTime"`
	LateTol                   cast.DurationConf        `json:"lateTolerance,omitempty" yaml:"lateTolerance,omitempty"`
	Concurrency               int                      `json:"concurrency" yaml:"concurrency"`
	BufferLength              int                      `json:"bufferLength" yaml:"bufferLength"`
	SendMetaToSink            bool                     `json:"sendMetaToSink" yaml:"sendMetaToSink"`
	SendNil                   bool                     `json:"sendNilField" yaml:"sendNilField"`
	SendError                 bool                     `json:"sendError" yaml:"sendError"`
	Qos                       Qos                      `json:"qos,omitempty" yaml:"qos,omitempty"`
	CheckpointInterval        cast.DurationConf        `json:"checkpointInterval,omitempty" yaml:"checkpointInterval,omitempty"`
	RestartStrategy           *RestartStrategy         `json:"restartStrategy,omitempty" yaml:"restartStrategy,omitempty"`
	Cron                      string                   `json:"cron,omitempty" yaml:"cron,omitempty"`
	Duration                  string                   `json:"duration,omitempty" yaml:"duration,omitempty"`
	CronDatetimeRange         []schedule.DatetimeRange `json:"cronDatetimeRange,omitempty" yaml:"cronDatetimeRange,omitempty"`
	PlanOptimizeStrategy      *PlanOptimizeStrategy    `json:"planOptimizeStrategy,omitempty" yaml:"planOptimizeStrategy,omitempty"`
	NotifySub                 bool                     `json:"notifySub,omitempty" yaml:"notifySub,omitempty"`
	DisableBufferFullDiscard  bool                     `json:"disableBufferFullDiscard,omitempty" yaml:"disableBufferFullDiscard,omitempty"`
	EnableSaveStateBeforeStop bool                     `json:"enableSaveStateBeforeStop,omitempty" yaml:"enableSaveStateBeforeStop,omitempty"`
	ForceExitTimeout          cast.DurationConf        `json:"forceExitTimeout,omitempty" yaml:"forceExitTimeout,omitempty"`
	Experiment                *ExpOpts                 `json:"experiment,omitempty" yaml:"experiment,omitempty"`
}

type ExpOpts struct {
	UseSliceTuple bool `json:"useSliceTuple" yaml:"useSliceTuple"`
}

type PlanOptimizeStrategy struct {
	EnableIncrementalWindow bool             `json:"enableIncrementalWindow" yaml:"enableIncrementalWindow"`
	EnableAliasPushdown     bool             `json:"enableAliasPushdown,omitempty" yaml:"enableAliasPushdown,omitempty"`
	DisableAliasRefCal      bool             `json:"disableAliasRefCal,omitempty" yaml:"disableAliasRefCal,omitempty"`
	OptimizeControl         *OptimizeControl `json:"optimizeControl,omitempty" yaml:"optimizeControl,omitempty"`
	WindowOption            *WindowOption    `json:"windowOption,omitempty" yaml:"windowOption,omitempty"`
}

type WindowOption struct {
	EnableSendSlidingWindowTwice bool   `json:"enableSendSlidingWindowTwice,omitempty" yaml:"enableSendSlidingWindowTwice,omitempty"`
	WindowVersion                string `json:"windowVersion,omitempty" yaml:"windowVersion,omitempty"`
}

func (p *PlanOptimizeStrategy) GetWindowVersion() string { _ = "STUB: not implemented"; return "" }

func (p *PlanOptimizeStrategy) IsAliasRefCalEnable() bool { _ = "STUB: not implemented"; return false }

func (p *PlanOptimizeStrategy) IsOptimizeEnabled(name string) bool {
	_ = "STUB: not implemented"
	return false
}

type OptimizeControl struct {
	DisableOptimizeRules []string `json:"disableOptimizeRules" yaml:"disableOptimizeRules"`
}

func (oc *OptimizeControl) IsOptimizeEnabled(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PlanOptimizeStrategy) IsSlidingWindowSendTwiceEnable() bool {
	_ = "STUB: not implemented"
	return false
}

type RestartStrategy struct {
	Attempts int `json:"attempts,omitempty" yaml:"attempts,omitempty"`
}

type PrintableTopo struct {
	Sources []string                 `json:"sources" yaml:"sources"`
	Edges   map[string][]interface{} `json:"edges" yaml:"edges"`
}

type GraphNode struct {
	Type     string                 `json:"type" yaml:"type"`
	NodeType string                 `json:"nodeType" yaml:"nodeType"`
	Props    map[string]interface{} `json:"props" yaml:"props"`
	// UI is a placeholder for ui properties
	UI map[string]interface{} `json:"ui" yaml:"ui"`
}

// SourceMeta is the metadata of a source node. It describes what existed stream/table to refer to.
// It is part of the Props in the GraphNode and it is optional
type SourceMeta struct {
	SourceName string `json:"sourceName"` // the name of the stream or table
	SourceType string `json:"sourceType"` // stream or table
}

type RuleGraph struct {
	Nodes map[string]*GraphNode `json:"nodes" yaml:"nodes"`
	Topo  *PrintableTopo        `json:"topo" yaml:"topo"`
}

// Rule the definition of the business logic
// Sql and Graph are mutually exclusive, at least one of them should be set
type Rule struct {
	Triggered bool                     `json:"triggered" yaml:"triggered"`
	Id        string                   `json:"id,omitempty" yaml:"id,omitempty"`
	Name      string                   `json:"name,omitempty" yaml:"name,omitempty"` // The display name of a rule
	Version   string                   `json:"version,omitempty" yaml:"version,omitempty"`
	Temp      bool                     `json:"temp,omitempty" yaml:"temp,omitempty"`
	Sql       string                   `json:"sql,omitempty" yaml:"sql,omitempty"`
	Graph     *RuleGraph               `json:"graph,omitempty" yaml:"graph,omitempty"`
	Actions   []map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`
	Options   *RuleOption              `json:"options,omitempty" yaml:"options,omitempty"`
	Tags      []string                 `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func (r *Rule) IsTagsMatch(tags []string) bool { _ = "STUB: not implemented"; return false }

func (r *Rule) IsDurationRule() bool { _ = "STUB: not implemented"; return false }

func (r *Rule) IsScheduleRule() bool { _ = "STUB: not implemented"; return false }

func (r *Rule) GetNextScheduleStartTime() int64 { _ = "STUB: not implemented"; return 0 }

func GetDefaultRule(name, sql string) *Rule { _ = "STUB: not implemented"; return nil }

const (
	AtMostOnce Qos = iota
	AtLeastOnce
	ExactlyOnce
)

type Qos int
