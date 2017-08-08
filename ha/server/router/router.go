package router

import(
    "net/http"
    c "github.com/Clark-zhang/learn-go/ha/server/controller"
)

func New(){
    http.HandleFunc("/setData", c.SetRedisData)
}