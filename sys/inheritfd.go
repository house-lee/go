package sys

import "sync"

type inheritFds struct {
	fdAttr        iFdAttr
	tcpListeners  []int
	udpConns      []int
	unixListeners []int
	fileFds       []int
}

var fds *inheritFds
var initInheritFd sync.Once

func InheritFd() *inheritFds {
	initInheritFd.Do(func() {
		fds = &inheritFds{
            fdAttr:       FdAttr(),
			tcpListeners:  make([]int, 0, 10),
			udpConns:      make([]int, 0, 10),
			unixListeners: make([]int, 0, 10),
			fileFds:       make([]int, 0, 10),
		}
	})
	return fds
}

func RegisterInheritFd(v interface{}) error {
    return nil
}