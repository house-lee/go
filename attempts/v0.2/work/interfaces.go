package work

import (
	"github.com/house-lee/SoarGO/coordinator"
	"github.com/house-lee/SoarGO/dsa"
)

type IRequest interface {
	ID() string
	Respond(resp interface{})
	Job() []byte
	UpdateJob(job []byte)
}

type IError interface {
	Code() int
	Message() string
}

type TaskHandler func(input []byte) (output []byte, err IError)

type IWorkStation interface {
	ID() string
	TaskQueue() dsa.IQueue
	SetCoordinator(c coordinator.ICoordinator)
	SetTaskHandler(handler TaskHandler)
	SetWorkerNum(num int)
	Restart()
	Start()
	LaunchTaskMonitor(serverID string, interval uint32) error
}
