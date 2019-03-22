package runner

import (
	"fmt"
	"strings"

	"github.com/laidingqing/stackbuild/core"
)

func buildEnviron(build *core.Build) map[string]string {
	env := map[string]string{
		"STACK_BUILD_BRANCH":         build.Target,
		"STACK_BUILD_TARGET_BRANCH":  build.Target,
		"STACK_BUILD_COMMIT":         build.After,
		"STACK_BUILD_COMMIT_SHA":     build.After,
		"STACK_BUILD_COMMIT_BEFORE":  build.Before,
		"STACK_BUILD_COMMIT_AFTER":   build.After,
		"STACK_BUILD_COMMIT_REF":     build.Ref,
		"STACK_BUILD_COMMIT_BRANCH":  build.Target,
		"STACK_BUILD_COMMIT_LINK":    build.Link,
		"STACK_BUILD_COMMIT_MESSAGE": build.Message,
		"STACK_BUILD_COMMIT_AUTHOR":  build.Author,
		"STACK_BUILD_BUILD_CREATED":  fmt.Sprint(build.Created),
		"STACK_BUILD_BUILD_STARTED":  fmt.Sprint(build.Started),
		"STACK_BUILD_BUILD_FINISHED": fmt.Sprint(build.Finished),
	}
	if strings.HasPrefix(build.Ref, "refs/tags/") {
		env["STACK_BUILD_TAG"] = strings.TrimPrefix(build.Ref, "refs/tags/")
	}
	return env
}
