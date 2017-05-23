package router

import(
    "net/http"
    // "github.com/justinas/alice"
    c "github.com/Clark-zhang/learn-go/image_service/controller"
)

func New(){
    // login := alice.New(c.Login)

    http.HandleFunc("/imageservice/add", c.AddImage)
    // http.HandleFunc("/imageservice/newId", c.NewId)
    // http.HandleFunc("/imageservice/tree", c.Tree)
    // http.Handle("/getBook", login.ThenFunc(c.GetBook))
}