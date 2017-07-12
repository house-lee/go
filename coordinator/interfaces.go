package coordinator

import "github.com/house-lee/SoarGO/req"

type ICoordinator interface {
    Init(serverID string, config interface{}) error
    Pulse() error
    IsServerAlive(serverID string) (bool, error)
    GenerateRequestID() string
    SaveRequest(request req.IRequest) error
    UpdateRequestProgress(requestID, serverID, workstationID string, task string) error
    SaveRequestResult(requestID string, code int, output string) error //code == -1 indicates the request is still being processed
    GetRequestResult(requestID string) (code int, output string)//if code != -1, request result should be clear
    ClearRequest(requestID string) error
    GetRequestList(workstationID string) ([]req.IRequest, error)
    TakeOverRequest(request req.IRequest) error
}
