package sys

import "sync"

var inheritObj IInherit
var initInheritFd sync.Once

func InheritFd() IInherit {
    initInheritFd.Do(func() {
        inheritObj = &inheritFds{
            fdAttr:        FdAttr(),
            tcpListeners:  make([]int, 0, 10),
            udpConns:      make([]int, 0, 10),
            unixListeners: make([]int, 0, 10),
            fileFds:       make([]int, 0, 10),
        }
    })
    return inheritObj
}

var initFdAttr sync.Once
var fdAttrObj IFdAttr

func FdAttr() IFdAttr {
    initFdAttr.Do(func() {
        fdAttrObj = &fdAttr{}
    })
    return fdAttrObj
}