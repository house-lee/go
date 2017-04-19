package sys

type inheritFds struct {
	listeners []int
	files     []int
}

const (
    ListenerFd = iota
    FileFd
)
func RegisterInheritFd(fdType, fd int) {
    //TODO
    switch fdType {
    case ListenerFd:
    case FileFd:
    default:

    }
}

func GetInheritFds() {

}

func SetFdClosedOnProcessExit() {

}

func SetFdNonClosedOnProcessExit() {

}
