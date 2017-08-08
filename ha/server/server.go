package server

import (
    "fmt"
    "net/http"
    "log"
    "github.com/Clark-zhang/learn-go/ha/server/router"
    "github.com/Clark-zhang/learn-go/ha/server/controller"
    syncByHttp "github.com/Clark-zhang/learn-go/ha/sync/http"
)

func Run(conf map[string]string) {

    fmt.Println("Server start with config: ")
    fmt.Printf("%v", conf)

    controller.Conf = conf
    syncByHttp.Conf = conf

    router.New()

    log.Fatal(http.ListenAndServe(conf["port"], nil))
}