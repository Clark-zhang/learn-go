package controller

import(
    "fmt"
    "net/http"
    "encoding/json"
    // "bytes"
    "os"
    "io"
    glog "github.com/golang/glog"
    "github.com/Clark-zhang/learn-go/go_service/config"
    "github.com/Clark-zhang/learn-go/go_service/lib"
    "io/ioutil"
)

func ImageServiceAdd(w http.ResponseWriter, r *http.Request) {
    //32 << 20 == 32 * 2^20 == 32Mb
    r.ParseMultipartForm(32 << 20)
    res := make(map[string]string)

    conf := config.GetOverseaConfig()

    path := conf["imgFolder"]
    baseUrl := conf["imgBaseUrl"]

    file, handler, err := r.FormFile("file")
    if err != nil {
        glog.Error("FormFile: %s", err)
        res["code"] = "1"
        response, _ := json.Marshal(res)
        fmt.Fprintf(w, "%s\n", response)
        return
    }
    defer file.Close()

    if _, err := os.Stat(path); os.IsNotExist(err) {
       err := os.MkdirAll(path, 0777)
        if err != nil {
            glog.Error("MkdirAll: %s", err)
        }
    }

    fileName, _ := lib.RenameFileToUniqueId(handler.Filename)
    f, err := os.OpenFile(path + fileName, os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
        glog.Error("OpenFile: %s", err)
        return
    }
    defer f.Close()

    _, err = io.Copy(f, file)

    if err != nil {
        glog.Info("Copy: %s", err)
        res["code"] = "1"
    }else {
        glog.Info("Sucess " + path + fileName)
        res["code"] = "0"
        res["downloadUrl"] = baseUrl + fileName
    }

    response, err := json.Marshal(res)

    fmt.Fprintf(w, "%s\n", response)

    return
}

func ActsServiceAdd(w http.ResponseWriter, r *http.Request) {
    conf := config.GetOverseaConfig()
    basePath := conf["htmlFolder"]
    baseUrl := conf["htmlBaseUrl"]

    r.ParseMultipartForm(32 << 20)
    tapdId := r.Form.Get("tapdId")
    year := r.Form.Get("year")
    html := r.Form.Get("html")
    fileName := r.Form.Get("fileName")

    relativePath := year + "/" + tapdId + "/"
    path := basePath + relativePath

    if _, err := os.Stat(path); os.IsNotExist(err) {
       err := os.MkdirAll(path, 0777)
        if err != nil {
            glog.Error("MkdirAll: %s", err)
        }
    }


    err := ioutil.WriteFile(path + fileName, []byte(html), 0777)

    res := make(map[string]string)

    if err != nil{
        glog.Error("ioutil.WriteFile: %s", err)
        res["code"] = "1"
    }else {
        glog.Info("Sucess " + path + fileName)
        res["code"] = "0"
        res["downloadUrl"] = baseUrl + relativePath + fileName
    }

    response, err := json.Marshal(res)

    fmt.Fprintf(w, "%s\n", response)

    return
}