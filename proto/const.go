package proto

type RunnerStatus int

const (
	RunnerInitialing RunnerStatus = iota
	RunnerPulling    RunnerStatus = 1
	RunnerStarting   RunnerStatus = 2
	RunnerStopping   RunnerStatus = 3
	RunnerIdle       RunnerStatus = 4
)

type TaskStatus int

const (
	TaskWaiting    TaskStatus = iota
	TaskProcessing TaskStatus = 1
	TaskSuccess    TaskStatus = 2
	TaskError      TaskStatus = 3
)
