package coordination

import (
    "github.com/house-lee/SoarGO/wkf"
)

type ICoordinator interface {
    SetServerID(serverID string)
    GenerateRequestID() (requestID string)
    Pulse() error
    IsServerAlive(serverID string) (bool, error)
    UpdateRequestProgress(requestID string, ws wkf.IWorkStation, job wkf.IJob)//use iRequest
    GetRequestProgress(requestID string) (serverID string, workStationID string)
    SaveRequestProgress(requestID string, header []byte, body []byte) error
    GetRequestResult(requestID string) (header []byte, body []byte, error)
    GetTaskList(workStationID string) []ITask
    TakeOverTask(task ITask) (job wkf.IJob, error)
}

type ITask interface {
    RequestID() string
    ServerID() string
    WorkStationID() string
}
