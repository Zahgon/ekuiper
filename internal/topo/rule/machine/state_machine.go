package machine

import (
	"github.com/lf-edge/ekuiper/contract/v2/api"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type ActionSignal int

const (
	ActionSignalStart ActionSignal = iota
	ActionSignalStop
	ActionSignalScheduledStart
	ActionSignalScheduledStop
)

type RunState int

const (
	Stopped RunState = iota
	Starting
	Running
	Stopping
	ScheduledStop
	StoppedByErr
)

var StateName = map[RunState]string{
	Stopped:       "stopped", // normal stop and schedule terminated are here
	Starting:      "starting",
	Running:       "running",
	Stopping:      "stopping",
	ScheduledStop: "stopped: waiting for next schedule.",
	StoppedByErr:  "stopped by error",
}

type StateMachine struct {
	syncx.RWMutex
	currentState RunState
	actionQ      []ActionSignal
	// Metric RunState
	lastStartTimestamp int64
	lastStopTimestamp  int64
	lastWill           string
	logger             api.Logger
}

func NewStateMachine(logger api.Logger) StateMachine {
	_ = "STUB: not implemented"
	return *new(StateMachine)
}

func (s *StateMachine) TriggerAction(action ActionSignal) bool {
	_ = "STUB: not implemented"
	return false
}

// do stop

// s.logger.Infof("ignore schedule start action, because current RunState is %s", StateName[ss])

func (s *StateMachine) Transit(newState RunState, lastWill string) (chainAction bool) {
	_ = "STUB: not implemented"
	return false
}

// do nothing

func (s *StateMachine) PopAction() ActionSignal {
	_ = "STUB: not implemented"
	return *new(ActionSignal)
}

func (s *StateMachine) LastWill() string { _ = "STUB: not implemented"; return "" }

func (s *StateMachine) CurrentState() RunState { _ = "STUB: not implemented"; return *new(RunState) }

func (s *StateMachine) LastStartTimestamp() int64 { _ = "STUB: not implemented"; return 0 }

func (s *StateMachine) CurrentStateName() string { _ = "STUB: not implemented"; return "" }

func (s *StateMachine) LastStopTimestamp() int64 { _ = "STUB: not implemented"; return 0 }
