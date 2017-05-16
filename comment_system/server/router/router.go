package router

import(
    "net/http"
    // "github.com/justinas/alice"
    c "github.com/Clark-zhang/learn-go/comment_system/server/controller"
)

func New(){
    http.HandleFunc("/add", c.AddComment)
    http.HandleFunc("/get", c.GetComments)
}