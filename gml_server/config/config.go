package config

import (
    "os"
    "runtime"
    "path"
    "github.com/jinzhu/configor"
)

type configuration struct{
    APPName string `default:"app name"`

    DB struct {
        IP  string
        Port     uint   `default:"3306"`
        Name     string
        User     string `default:"root"`
        Password string `required:"true" env:"DBPassword"`
    }
}

var Config = configuration{}

func GetConfig() configuration{
    //just for test
    os.Setenv("CONFIGOR_ENV", "local")
    defer os.Setenv("CONFIGOR_ENV", "")

    //the path of this file
    _, currentFilePath, _, _ := runtime.Caller(0)
    dirpath := path.Dir(currentFilePath)

    env := os.Getenv("CONFIGOR_ENV")
    configor.Load(&Config, dirpath + "/config." + env + ".yml")

    return Config
}