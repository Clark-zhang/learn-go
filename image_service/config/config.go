package config

import (
    "os"
)

type configuration struct{
    Folder string `default:"/tmp/clark_image_service"`
}

var Config = configuration{}

func GetConfig() configuration{
    //just for test
    os.Setenv("CONFIGOR_ENV", "local")
    defer os.Setenv("CONFIGOR_ENV", "")

    env := os.Getenv("CONFIGOR_ENV")

    switch env {
    case "local":
        Config = configuration{
            Folder: "/tmp/clark_image_service/",
        }
    }

    return Config
}