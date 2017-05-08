package server

import (
	"github.com/house-lee/SoarGO/sys"
	"net"
	"net/http"
)

type HttpServer struct {
	server         *http.Server
	listener       net.Listener
	inheritManager sys.IInherit
}
