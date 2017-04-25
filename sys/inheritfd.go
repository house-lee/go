package sys

import (
	"net"
	"os"
	"sync"
	"errors"
)

type inheritFds struct {
	fdAttr        iFdAttr
	tcpListeners  []int
	tcpLock       sync.RWMutex
	udpConns      []int
	udpLock       sync.RWMutex
	unixListeners []int
	unixLock      sync.RWMutex
	fileFds       []int
	fileLock      sync.RWMutex
}

var fds *inheritFds
var initInheritFd sync.Once

func InheritFd() *inheritFds {
	initInheritFd.Do(func() {
		fds = &inheritFds{
			fdAttr:        FdAttr(),
			tcpListeners:  make([]int, 0, 10),
			udpConns:      make([]int, 0, 10),
			unixListeners: make([]int, 0, 10),
			fileFds:       make([]int, 0, 10),
		}
	})
	return fds
}

func (iFds *inheritFds)RegisterInheritFd(v interface{}) error {
	switch v := v.(type) {
	case *net.TCPListener:
		fd, e := getFd(v)
		if e != nil {
			return e
		}
		if e := iFds.fdAttr.SetFdNonCloseOnProcessExit(fd); e != nil {
			return e
		}
		iFds.tcpLock.Lock()
		iFds.tcpListeners = append(iFds.tcpListeners, fd)
		iFds.tcpLock.Unlock()
	case *net.UDPConn:
		fd, e := getFd(v)
		if e != nil {
			return e
		}
		if e := iFds.fdAttr.SetFdNonCloseOnProcessExit(fd); e != nil {
			return e
		}
		iFds.udpLock.Lock()
		iFds.udpConns = append(iFds.udpConns, fd)
		iFds.udpLock.Unlock()
	case *net.UnixListener:
		fd, e := getFd(v)
		if e != nil {
			return e
		}
		if e := iFds.fdAttr.SetFdNonCloseOnProcessExit(fd); e != nil {
			return e
		}
		iFds.unixLock.Lock()
		iFds.unixListeners = append(iFds.unixListeners, fd)
		iFds.unixLock.Unlock()
	case *os.File:
		fd := int(v.Fd())
		if e := iFds.fdAttr.SetFdNonCloseOnProcessExit(fd); e != nil {
			return e
		}
		iFds.fileLock.Lock()
		iFds.fileFds = append(iFds.fileFds, fd)
		iFds.fileLock.Unlock()
	default:
		return errors.New("Non-inheritable File Type")
	}
	return nil
}

func (iFds *inheritFds)GetInheritFds() map[string][]int {
	return nil
}

type iFile interface {
	File() (f *os.File, err error)
}

func getFd(v iFile) (int, error) {
	f, e := v.File()
	if e != nil {
		return 0, e
	}
	return int(f.Fd()), nil
}
