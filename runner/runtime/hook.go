package runtime

// Hook 提供一组运行时钩子
type Hook struct {
	// Before 执行前调用
	Before func(*State) error

	//BeforeEach ..each step is executed
	BeforeEach func(*State) error

	// After 执行后调用.
	After func(*State) error

	//AfterEach ..each step is executed
	AfterEach func(*State) error

	// GotLine 当生产一行容器日志时调用
	GotLine func(*State, *Line) error

	//GotLogs when the logs are completed
	GotLogs func(*State, []*Line) error
}
