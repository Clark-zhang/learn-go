package main

import(
    "net/http"
    "github.com/justinas/alice"
)

func New(){
    login := alice.New(c.Login)

    http.HandleFunc("/bookId", c.GetBookId)
    http.Handle("/getBook", login.ThenFunc(c.GetBook))

    http.Handle("/needLogin", login.ThenFunc(c.GetBook))
    http.Handle("/redirect", login.ThenFunc(c.GetBook))
}