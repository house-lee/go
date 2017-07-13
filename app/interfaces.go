package app

import (
    "github.com/house-lee/SoarGO/factory"
    "time"
)

type IRoute interface {

}

type IServer interface {
    Init(serverID string, config interface{})
    RegisterWorkflow(routingRules IRoute, wf factory.IWorkflow)
    Serve()
    Stop() error
    GracefulStop(maxWaitTime time.Duration) error
}

type IApp interface {
    RegisterServer(s IServer)
    Run() error
    Shutdown() error
    GracefulShutdown(maxWaitTime time.Duration) error
    Restart(newBinaryPath string, maxWaitTime time.Duration) error
}