package server

import (
	"github.com/house-lee/SoarGO/wkf"
	"net"
    "github.com/house-lee/SoarGO/coordination"
)

type IServer interface {
	Serve(listener net.Listener) error
	GracefulStop(maxTimeout int) error
    Stop() error
	RegisterWorkflow(routes []interface{}, workflow wkf.IWorkflow) error
    ID() string
    SetCoordinator(coordinator coordination.ICoordinator)
}

//IUDPServer hasn't been implemented yet
type IUDPServer interface {
    Serve(conn *net.UDPConn)
    GracefulStop(maxTimeout int) error
    Stop() error
    RegisterWorkflow(routes []interface{}, workflow wkf.IWorkflow) error
}