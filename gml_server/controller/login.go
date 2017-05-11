package controller

import(
    "net/http"
    "fmt"
)

func Login(next http.Handler) http.Handler{

    fn := func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/needLogin" {
             LoginHandler(w, r)
        }else if r.URL.Path == "/redirect" {
             http.Redirect(w, r, "/needLogin", 301)
        }else{
             next.ServeHTTP(w, r)
        }
    }

    return http.HandlerFunc(fn)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Please login to access")
}