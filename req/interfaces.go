package req

import "github.com/house-lee/SoarGO/factory"

type IRequest interface {
    factory.ITask
    IsAsync() bool
    OriginalReq() []byte
    CurrentServerID() (serverID string)
    CurrentWorkStation() factory.IWorkStation
    CurrentTask() factory.ITask
    CurrentResult() (code int, output string)
    SendResponse(code int, message string)
}

type IError interface {
    Code() int
    Message() string
}
