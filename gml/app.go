package main

import (
    "fmt"
    "github.com/jinzhu/configor"
)

var Config = struct {
    APPName string `default:"app name"`

    DB struct {
        IP  string
        Port     uint   `default:"3306"`
        Name     string
        User     string `default:"root"`
        Password string `required:"true" env:"DBPassword"`
    }
}{}

func main() {
    configor.Load(&Config, "config.yml")
    fmt.Println(Config)
}