package main

import (
    "crypto/tls"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main()  {
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        fmt.Println(err)
        return
    }
    config := &tls.Config{Certificates:[]tls.Certificate{cert}}
    listener, err := tls.Listen("tcp", ":1443", config)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()

    server := &http.Server{
        Addr: ":1443",
        Handler:buildRouter(),
    }
    err = server.Serve(listener)
    if err != nil {
        fmt.Println(err)
    }

}

func buildRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/hello/", Hello)
    return router
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(
        w,
        "Index: Hello World!\n Serving by %d",
        os.Getpid(),
    )
}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(
        w,
        "Hello: Hello World!\n Serving by %d",
        os.Getpid(),
    )
}