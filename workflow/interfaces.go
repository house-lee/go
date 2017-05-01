package workflow

import "github.com/house-lee/SoarGO/dsa"


type IWorkflow interface {
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
    LaunchWorkStation(workStationID string) error
}

type IWorker interface {
    Do(data interface{}) (result interface{}, err error)
}