package main

import (
    "github.com/house-lee/SoarGO/v1/sys"
    "net"
    "fmt"
    "os"
)

func main()  {
    inheritObj := sys.InheritFd()
    var tcplistener_1, tcplistener_2 net.Listener
    var err error
    tcplistener_1, err = net.Listen("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println(err.Error())
    }
    tcplistener_2, err = net.Listen("tcp4", ":9999")
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(tcplistener_1)
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(tcplistener_2)
    if err != nil {
        fmt.Println(err.Error())
    }

    addr1, _ := net.ResolveUDPAddr("udp", ":19999")
    addr2, _ := net.ResolveUDPAddr("udp", ":29999")
    var udpConn_1, udpConn_2 *net.UDPConn
    udpConn_1, err = net.ListenUDP("udp", addr1)
    if err != nil {
        fmt.Println(err.Error())
    }
    udpConn_2, err = net.ListenUDP("udp", addr2)
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(udpConn_1)
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(udpConn_2)
    if err != nil {
        fmt.Println(err.Error())
    }

    var unixListener_1, unixListener_2 net.Listener
    unixListener_1, err = net.Listen("unix", "127.0.0.1:33333")
    if err != nil {
        fmt.Println(err.Error())
    }
    unixListener_2, err = net.Listen("unix", ":22222")
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(unixListener_1)
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(unixListener_2)
    if err != nil {
        fmt.Println(err.Error())
    }

    var file1,file2 *os.File
    file1, err = os.Open("a.txt")
    if err != nil {
        fmt.Println(err.Error())
    }
    file2, err = os.Open("b.txt")
    if err != nil {
        fmt.Println(err.Error())
    }

    err = inheritObj.RegisterInheritFd(file1)
    if err != nil {
        fmt.Println(err.Error())
    }
    err = inheritObj.RegisterInheritFd(file2)
    if err != nil {
        fmt.Println(err.Error())
    }

    m := inheritObj.GetInheritFds()
    fmt.Printf("%+v\n", m)
}
