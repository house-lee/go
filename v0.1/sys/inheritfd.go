package sys

import (
	"errors"
	"net"
	"os"
	"sync"
)

type inheritFds struct {
	fdAttr        IFdAttr
	tcpListeners  []int
	tcpLock       sync.RWMutex
	udpConns      []int
	udpLock       sync.RWMutex
	unixListeners []int
	unixLock      sync.RWMutex
	fileFds       []int
	fileLock      sync.RWMutex
}


func (iFds *inheritFds) RegisterInheritFd(v interface{}) error {
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

func (iFds *inheritFds) GetInheritFds() map[string][]int {
	storedFds := map[string][]int{}
	defer func() {
		iFds.tcpLock.RUnlock()
		iFds.udpLock.RUnlock()
		iFds.unixLock.RUnlock()
		iFds.fileLock.RUnlock()
	}()

	iFds.tcpLock.RLock()
	l := len(iFds.tcpListeners)
	if l != 0 {
		storedFds["tcp"] = make([]int, l)
		for i := 0; i != l; i++ {
			storedFds["tcp"][i] = iFds.tcpListeners[i]
		}
	}

	iFds.udpLock.RLock()
	l = len(iFds.udpConns)
	if l != 0 {
		storedFds["udp"] = make([]int, l)
		for i := 0; i != l; i++ {
			storedFds["udp"][i] = iFds.udpConns[i]
		}
	}

	iFds.unixLock.RLock()
	l = len(iFds.unixListeners)
	if l != 0 {
		storedFds["unix"] = make([]int, l)
		for i := 0; i != l; i++ {
			storedFds["unix"][i] = iFds.unixListeners[i]
		}
	}

	iFds.fileLock.RLock()
	l = len(iFds.fileFds)
	if l != 0 {
		storedFds["file"] = make([]int, l)
		for i := 0; i != l; i++ {
			storedFds["file"][i] = iFds.fileFds[i]
		}
	}

	return storedFds
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
