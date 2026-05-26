package rule

import (
	"github.com/lf-edge/ekuiper/v2/internal/pkg/def"
	"github.com/lf-edge/ekuiper/v2/internal/topo"
	"github.com/lf-edge/ekuiper/v2/internal/topo/rule/machine"
)

const EOFMessage = "done"

func (s *State) doValidateAndRun(newRule *def.Rule) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Try plan with the new json. If err, revert to old rule

// validateRule only check plan is valid, topology shouldn't be changed before ruleState stop

// stop the old run

// start new rule

// Start the rule which runs async

// Discard the temp topo

// If validate error, the return tp is clean up and set to nil
func (s *State) validate() (tp *topo.Topo, err error) {
	_ = "STUB: not implemented"
	// Do validation
	return nil, nil
}

// clean topo if error happens

// DoStart runs internally
func (s *State) doStart() error {
	_ = "STUB: not implemented"
	// Start normally or start in schedule period Rule
	// doStart trigger the Rule run. If no trigger error, the Rule will run async and control the state by itself
	return nil
}

func (s *State) doStop(stateType machine.RunState, msg string) { _ = "STUB: not implemented"; return }

// This is called async
func (s *State) runTopo(tp *topo.Topo, ruleId string) { _ = "STUB: not implemented"; return }

// Only restart Rule for errors

// exit normally

// The run exit may be caused by user action or rule itself
// Only do clean up when it is exit automatically

func (s *State) cleanRule(tp *topo.Topo, hasError bool, lastWill string) {
	_ = "STUB: not implemented"
	return
}
