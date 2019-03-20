package executor

import (
	"context"
	"io"
)

// Executor 运行时引擎接口.
type Executor interface {
	// Setup 设置环境.
	Setup(context.Context, *Spec) error

	// Create 创建状态
	Create(context.Context, *Spec, *Step) error

	// Start 开始步骤
	Start(context.Context, *Spec, *Step) error

	// Wait .
	Wait(context.Context, *Spec, *Step) (*State, error)

	// Tail tail logs
	Tail(context.Context, *Spec, *Step) (io.ReadCloser, error)

	// Destroy 销毁
	Destroy(context.Context, *Spec) error
}
