package coordinator

import (
    wkf "github.com/house-lee/SoarGO/workflow"
)

type ICoordinator interface {
    NewRequest() (requestID string)
    UpdateRequestProgress(requestID string, workflow wkf.IWorkflow, workStation wkf.IWorkStation, job wkf.IJob, msg string) error
    GetRequestProgress(requestID string) (workStationID string)
    SaveRequestOutput(requestID string, output string, append bool) error
    GetRequestOutput(requestID string) (output string, err error)
    GetTaskList(workStationID string)//think about the return value
}
