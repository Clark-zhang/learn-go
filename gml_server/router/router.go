package router

import(
    "net/http"
    "github.com/justinas/alice"
    c "github.com/Clark-zhang/learn-go/gml_server/controller"
)

func New(){
    login := alice.New(c.Login)

    http.HandleFunc("/bookId", c.GetBookId)
    http.Handle("/getBook", login.ThenFunc(c.GetBook))

    http.Handle("/needLogin", login.ThenFunc(c.GetBook))
    http.Handle("/redirect", login.ThenFunc(c.GetBook))
}