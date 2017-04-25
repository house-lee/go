package sys

type iInherit interface {
	RegisterInheritFd(v interface{}) error
	GetInheritFds() map[string][]int
}

type iFdAttr interface {
    SetFdCloseOnProcessExit(fd int) error
    SetFdNonCloseOnProcessExit(fd int) error
}
