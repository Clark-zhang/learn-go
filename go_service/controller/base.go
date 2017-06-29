package controller

import(
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    glog "github.com/golang/glog"
)

func CheckMultipartFormMiddleware(next http.Handler) http.Handler{

    fn := func(w http.ResponseWriter, r *http.Request) {
        if ct := r.Header.Get("Content-Type"); strings.Contains(ct, "multipart/form-data") {
            next.ServeHTTP(w, r)
        }else{
            glog.Info("checkMultipartFormg: " , r.Header)
            glog.Info("123" , r.Header.Get("Content-Type"), "321")

            res := make(map[string]string)

            res["code"] = "1"
            res["message"] = "header 'Content-Type' must be 'multipart/form-data'"
            response, _ := json.Marshal(res)
            fmt.Fprintf(w, "%s\n", response)
        }
    }

    return http.HandlerFunc(fn)
}