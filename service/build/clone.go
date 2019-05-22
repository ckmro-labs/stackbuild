package build

import (
	"github.com/laidingqing/stackbuild/core"
)

const cloneStepName = "clone"

// create clone build step..
func createClone(src *core.Stage) *core.BuildStep {

	return &core.BuildStep{
		Name: cloneStepName,
	}
}
