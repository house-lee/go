package coordinator

import (
    . "github.com/house-lee/SoarGO/workflow"
)

type ICoordinator interface {
    NewRequest() (requestID string)
    UpdateRequestProgress(requestID string, workflow IWorkflow, workStation IWorkStation, job IJob, msg string) error
}
