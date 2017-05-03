package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/julienschmidt/httprouter"
    "github.com/Clark-zhang/learn-go/gml_server/router"
    "github.com/Clark-zhang/learn-go/gml_server/config"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func main() {
    Config := config.GetConfig()

    fmt.Println("Server start with config: ")
    fmt.Printf("%v", Config)

    router := router.New()
    log.Fatal(http.ListenAndServe(":7007", router))
}