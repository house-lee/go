package req

import (
    "io"
    "github.com/house-lee/SoarGO/v1/wkf"
)

type IRequest interface {
    RequestID() string
    IsAsync() bool
    Client() IClient
    OriginalRequest()(header, body []byte)
    SetRequestResult(header, body []byte, error)
    GetRequestResult()(header, body []byte, error)
    CurrentJob() wkf.IJob
}

type IClient interface {
    io.WriteCloser
    isAlive() bool
}