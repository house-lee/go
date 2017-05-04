package coordinator

import (
    wkf "github.com/house-lee/SoarGO/workflow"
)

type ICoordinator interface {
    SetServerID(serverID string)
    GenerateRequestID() (requestID string)
    Pulse() error
    IsServerAlive(serverID string) (bool, error)
    UpdateRequestProgress(requestID string, ws wkf.IWorkStation, job wkf.IJob)
}
