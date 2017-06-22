package controller

import(
    "fmt"
    "net/http"
    // "encoding/json"
    // "bytes"
    "os"
    "io"
    glog "github.com/golang/glog"
)

var Path string

func AddImage(w http.ResponseWriter, r *http.Request) {
    //32 << 20 == 32 * 2^20 == 32M?
    r.ParseMultipartForm(32 << 20)

    servicePath := Path + "/" + r.Form.Get("service") + "/"

    file, handler, err := r.FormFile("file")
    if err != nil {
        glog.Fatalf("FormFile: %s", err)
        return
    }
    defer file.Close()

    // fmt.Fprintf(w, "%v", handler.Header)

    if _, err := os.Stat(servicePath); os.IsNotExist(err) {
       os.MkdirAll(servicePath, 0777)
    }

    f, err := os.OpenFile(servicePath + handler.Filename, os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
        glog.Fatalf("OpenFile: %s", err)
        return
    }
    defer f.Close()

    _, err = io.Copy(f, file)

    if err != nil {
        glog.Fatalf("Copy: %s", err)
        fmt.Fprintf(w, "%s\n", "Fail")
    }else {
        glog.Info("Sucess " + servicePath + handler.Filename)
        fmt.Fprintf(w, "%s\n", "Success")
    }

    return
}