package server

import (
	wkf "github.com/house-lee/SoarGO/workflow"
	"net"
)

type IServer interface {
	Serve(listener net.Listener) error
	Stop() error
	RegisterWorkflow(routes []interface{}, workflow wkf.IWorkflow) error
}

type IUDPServer interface {
    Serve(conn *net.UDPConn)
    Stop() error
    RegisterWorkflow(routes []interface{}, workflow wkf.IWorkflow) error
}