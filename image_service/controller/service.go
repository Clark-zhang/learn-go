package controller

import(
    "fmt"
    "net/http"
    // "encoding/json"
    "github.com/Clark-zhang/learn-go/image_service/config"
    // "bytes"
    "os"
    "io"
)

func AddImage(w http.ResponseWriter, r *http.Request) {
    //32 << 20 == 32 * 2^20 == 32M?
    r.ParseMultipartForm(32 << 20)

    conf := config.GetConfig()

    servicePath := conf.Folder + r.Form.Get("service") + "/"

    file, handler, err := r.FormFile("file")
    if err != nil {
       fmt.Println(err)
       return
    }
    defer file.Close()

    // fmt.Fprintf(w, "%v", handler.Header)

    f, err := os.OpenFile(servicePath + handler.Filename, os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
       fmt.Println(err)
       return
    }
    defer f.Close()

    _, err = io.Copy(f, file)

    if err != nil {
        fmt.Fprintf(w, "%s\n", "Fail")
    }else {
        fmt.Fprintf(w, "%s\n", "Success")
    }

    return
}