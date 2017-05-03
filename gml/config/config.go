package config

import (
    "os"
    "runtime"
    "path"
    "github.com/jinzhu/configor"
)

type configuration struct{
    ServiceUrl string `default:"localhost:7007"`
    ServiceHost string `default:"localhost"`
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