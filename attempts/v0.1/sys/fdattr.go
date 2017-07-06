package sys

import (
    "syscall"
)

type fdAttr struct {}

func (fda *fdAttr) SetFdCloseOnProcessExit(fd int) error {
    _,err := fcntl(fd, syscall.F_SETFD, syscall.FD_CLOEXEC)
    return err
}

func (fda *fdAttr) SetFdNonCloseOnProcessExit(fd int) error  {
    _,err := fcntl(fd, syscall.F_SETFD, ^syscall.FD_CLOEXEC)
    return err
}

func fcntl(fd int, cmd int, arg int) (int, error) {
    val,_,err := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
    if err != 0 {
        return 0,err
    }
    return int(val),nil
}