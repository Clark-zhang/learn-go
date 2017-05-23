package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/Clark-zhang/learn-go/image_service/router"
    "github.com/Clark-zhang/learn-go/image_service/config"
)

func main() {
    Config := config.GetConfig()

    fmt.Println("Server start with config: ")
    fmt.Printf("%v", Config)

    router.New()
    log.Fatal(http.ListenAndServe(":7007", nil))
}