package wkf

import (
    "github.com/house-lee/SoarGO/dsa"
    "github.com/house-lee/SoarGO/v1/req"
)


type IWorkflow interface {
    GetWorkflowID() string
    RegisterWorkStation(station IWorkStation)
    HandleRequest(request req.IRequest) error
    LaunchWorkFlow(workflowID string) error
}

type ITaskMonitor interface {
    LaunchTaskMonitor(serverID string, workflowID string, workStationID string, interval uint32) error
}

type IWorkStation interface {
    ITaskMonitor
    //GetInputQueue() dsa.IQueue //replace with NewTask maybe?
    //SetWorker(worker IWorker, num uint32) error
    GetStationID() string
    LaunchWorkStation(workStationID string) error
    SetServerIdentification(serverID string)
}

type IJob interface {
    Serialize() ([]byte, error)
    Unserialize(obj interface{}) error
}

type TaskHandler func(job IJob)(result IJob, err error)