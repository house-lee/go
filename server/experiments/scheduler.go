package main

import (
    "fmt"
    "runtime"
    "time"
)

func cpuIntensive(p *int64) {
    for i := int64(1); i <= 10000000; i++ {
        *p = *p + 1
    }
    //fmt.Println("Done intensive thing")
}

func printVar(p *int64) {
    fmt.Printf("print x = %d.\n", *p)
}

func main() {
    runtime.GOMAXPROCS(1)

    x := int64(0)
    for i := 0; i != 10; i++ {
        go cpuIntensive(&x) // This should go into background
        go printVar(&x)
    }

    // This won't get scheduled until everything has finished.
    time.Sleep(1 * time.Second) // Wait for goroutines to finish
}