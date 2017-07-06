package main

import (
    "fmt"
    "github.com/house-lee/SoarGO/cfg"
    "github.com/house-lee/SoarGO/cfg/example/conf"
    "os"
)

func main()  {
    if err := cfg.LoadConf(&conf.ServerFromFile, "./runtime.conf"); err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%+v\n", conf.ServerFromFile)

    prepareEnv()
    if err := cfg.LoadEnv(&conf.ServerFromEnv); err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%+v\n", conf.ServerFromEnv)
}

func prepareEnv() {
    env := map[string] string {
        "PORT" : "6379",
        "MYSQL_HOST" : "localhost",
        "MYSQL_PORT" : "3306",
        "RUN_IN_BACKGROUND": "yes",
        "PI" : "3.1415926",
    }
    for k,v := range env {
        os.Setenv(k,v)
    }
}
