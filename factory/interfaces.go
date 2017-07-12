package factory

import (
    "github.com/house-lee/SoarGO/coordinator"
    "github.com/house-lee/SoarGO/dsa"
    "github.com/house-lee/SoarGO/req"
)



type IWorkflow interface {
    Init(id string) error
    HandleRequest(request req.IRequest) error
    RegisterWorkStation(ws IWorkStation)
}

type TaskHandler func(input ITask) (output ITask, err req.IError)

type IWorkStation interface {
	Init(id string, c coordinator.ICoordinator, workerNum uint, handler TaskHandler) error
    TaskQueue() dsa.IQueue
    ResetWorkerNum(num uint)
    Start() error
    Restart() error
    LaunchTaskMonitor(interval uint32) error
}

type ITask interface {
	ToString() string
}
