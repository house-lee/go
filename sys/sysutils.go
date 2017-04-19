package sys

type inheritFds struct {
	listeners []int
	files     []int
}

const (
    TCPListenerFd = iota
    UDPListenerFd
    UnixListenerFd
    FileFd
)
func RegisterInheritFd(fdType, fd int) {
    //TODO
    switch fdType {
    case TCPListenerFd:
    case UDPListenerFd:
    case UnixListenerFd:
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
