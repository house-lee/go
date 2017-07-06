package sys

type IInherit interface {
	RegisterInheritFd(v interface{}) error
	GetInheritFds() map[string][]int
}

type IFdAttr interface {
    SetFdCloseOnProcessExit(fd int) error
    SetFdNonCloseOnProcessExit(fd int) error
}
