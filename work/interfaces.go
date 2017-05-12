package work

import (
	"github.com/house-lee/SoarGO/coordinator"
	"github.com/house-lee/SoarGO/dsa"
	"io"
)

type IRequest interface {
	ID() string
	ResponseWriter() io.Writer
	Job() []byte
	UpdateJob(interface{})
}

type IError interface {
	Code() int
	Message() string
}

type TaskHandler func(input interface{}) (output interface{}, err IError)

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
