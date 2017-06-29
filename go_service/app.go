package main

import (
    "net/http"
    "log"
    "flag"
    "github.com/golang/glog"
    "github.com/Clark-zhang/learn-go/go_service/router"
    "github.com/Clark-zhang/learn-go/go_service/config"
)

func main() {
    conf := config.GetOverseaConfig()

    flag.Parse()
    flag.Lookup("log_dir").Value.Set(conf["logFolder"])

    router.New()

    glog.Info("Server start at " + conf["port"])

    log.Fatal(http.ListenAndServe(":" + conf["port"], nil))
}