package wkf

import (
    "github.com/house-lee/SoarGO/dsa"
    "github.com/house-lee/SoarGO/req"
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
    GetInputQueue() dsa.IQueue
    SetWorker(worker IWorker, num uint32) error
    GetStationID() string
    LaunchWorkStation(workStationID string) error
    SetServerIdentification(serverID string)
}

type IJob interface {
}
type IWorker interface {
    Do(job IJob) (result IJob, err error)
}
