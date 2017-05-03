package coordinator

import (
    . "github.com/house-lee/SoarGO/workflow"
)

type ICoordinator interface {
    NewRequest() (requestID string)
    UpdateRequestProgress(requestID string, workflow IWorkflow, workStation IWorkStation, job IJob, msg string) error
    GetRequestProgress(requestID string) (workStationID string)
    SaveRequestOutput(requestID string, output string, append bool) error
    GetRequestOutput(requestID string) (output string, err error)
    GetTaskList(workStationID string)//think about the return value
}
