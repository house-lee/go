package server

import (
	wkf "github.com/house-lee/SoarGO/workflow"
	"net"
)

type IServer interface {
	Serve(listener net.Listener) error
	GracefulStop(maxTimeout int) error
    Stop() error
	RegisterWorkflow(routes []interface{}, workflow wkf.IWorkflow) error
}

type IUDPServer interface {
    Serve(conn *net.UDPConn)
    GracefulStop(maxTimeout int) error
    Stop() error
    RegisterWorkflow(routes []interface{}, workflow wkf.IWorkflow) error
}