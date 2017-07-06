package coordinator

type ICoordinator interface {
    SetServerID(serverID string)
    UpdateRequestProgress(reqID string, workStationID string, job []byte)
    SaveFailedRequest(reqID string, errCode int, errMsg string)
}
