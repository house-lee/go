package svr

import (
	"github.com/house-lee/SoarGO/v1/sys"
	"net"
	"net/http"
)

type HttpServer struct {
	server         *http.Server
	listener       net.Listener
	inheritManager sys.IInherit
}
