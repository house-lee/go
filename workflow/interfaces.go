package workflow

import "github.com/house-lee/SoarGO/dsa"


type IWorkflow interface {
    GetWorkflowID() string
    RegisterWorkStation(station IWorkStation)
    HandleRequest(reqID string, reqBody interface{}, isAsync bool) error
    LaunchWorkFlow(workflowID string) error
}

type ITaskMonitor interface {
    LaunchTaskMonitor(workflowID string, workStationID string, interval uint32) error
}

type IWorkStation interface {
    ITaskMonitor
    GetInputQueue() dsa.IQueue
    SetWorker(worker IWorker, num uint32) error
    GetStationID() string
    LaunchWorkStation(workStationID string) error
}

type IJob interface {
    ToString() string
}
type IWorker interface {
    Do(job IJob) (result IJob, err error)
}