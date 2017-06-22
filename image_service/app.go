//Usage:
// mkdir -p /tmp/clark-log
// go run app.go -path=/tmp/clark-service -port=7007 -log_dir=/tmp/clark-log &



package main

import (
    "net/http"
    "log"
    "flag"
    "github.com/Clark-zhang/learn-go/image_service/router"
    "github.com/golang/glog"
    controller "github.com/Clark-zhang/learn-go/image_service/controller"
)

func main() {
    path := flag.String("path", "/tmp/learn-go/file_service/", "path to store files")
    port := flag.String("port", "7007", "port the service will run")
    flag.Parse()

    controller.Path = *path

    router.New()

    glog.Info("Server start with path" + *path + " on port" + *port)

    log.Fatal(http.ListenAndServe(":" + *port, nil))
}