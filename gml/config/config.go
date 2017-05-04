package config

import (
    "os"
)

type configuration struct{
    ServiceUrl string `default:"http://localhost:7007"`
    ServiceHost string `default:"localhost"`
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
            ServiceUrl: "http://localhost:7007",
            ServiceHost: "localhost",
        }
    }

    return Config
}