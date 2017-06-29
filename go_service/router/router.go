package router

import(
    "net/http"
    "github.com/justinas/alice"
    c "github.com/Clark-zhang/learn-go/go_service/controller"
)

func New(){
    // login := alice.New(c.Login)
    // http.HandleFunc("/imageservice/newId", c.NewId)
    // http.HandleFunc("/imageservice/tree", c.Tree)
    // http.Handle("/getBook", login.ThenFunc(c.GetBook))

    checkHeader := alice.New(c.CheckMultipartFormMiddleware)

    http.Handle("/imageservice/add", checkHeader.ThenFunc(c.ImageServiceAdd))
    http.HandleFunc("/actsservice/add", c.ActsServiceAdd)
}