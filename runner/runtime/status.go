package runtime

import "github.com/laidingqing/stackbuild/runner/executor"

type status struct {
	step  *executor.Step
	state *executor.State
}

// isSerial returns true if the steps are to be executed
func isSerial(spec *executor.Spec) bool {
	for _, step := range spec.Steps {
		if len(step.DependsOn) != 0 {
			return false
		}
	}
	return true
}

// nextStep ..
func nextStep(spec *executor.Spec, states map[string]*status) *executor.Step {
LOOP:
	for _, step := range spec.Steps {
		state := states[step.Metadata.Name]
		if state.state != nil {
			continue
		}
		if len(step.DependsOn) == 0 {
			return step
		}
		for _, name := range step.DependsOn {
			state, ok := states[name]
			if !ok {
				continue
			}
			if state.step.Detach {
				continue
			}
			if state.step.RunPolicy == executor.RunNever {
				continue
			}
			if state.state == nil || state.state.Exited == false {
				break LOOP
			}
		}
		return step
	}
	return nil
}
