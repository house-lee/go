package main

import (
    "fmt"
    "github.com/house-lee/SoarGO/config"
    "github.com/house-lee/SoarGO/config/example/conf"
    "os"
)

func main()  {
    if err := config.LoadConf(&conf.ServerFromFile, "./runtime.conf"); err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%+v\n", conf.ServerFromFile)

    prepareEnv()
    if err := config.LoadEnv(&conf.ServerFromEnv); err != nil {
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
