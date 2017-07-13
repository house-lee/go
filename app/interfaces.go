package app

import (
    "github.com/house-lee/SoarGO/factory"
    "time"
)

type IRoute interface {}

type IServer interface {
    Init(serverID string, config interface{})
    RegisterWorkflow(routingRules IRoute, wf factory.IWorkflow)
    FindWorkflow(route IRoute) (factory.IWorkflow, error)
    Serve()
    Stop() error
    GracefulStop(maxWaitTime time.Duration) error
}

type IApp interface {
    RegisterServer(s IServer)
    //if maxRealThreads == 0, maxRealThreads = Num(Worker)
    //elif maxRealThreads == -1, maxRealThreads = Num(CPU)
    Run(maxRealThreads int32) error
    Shutdown() error
    GracefulShutdown(maxWaitTime time.Duration) error
    Restart(newBinaryPath string, maxWaitTime time.Duration) error
}