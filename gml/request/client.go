package request

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "github.com/Clark-zhang/learn-go/gml/config"
)


type Request struct{
    Url string
    Method string
    Data map[string]string
    R chan<-string
}

func (r Request) MakeRequest(){
    Config := config.GetConfig()

    reqUrl := Config.ServiceUrl + r.Url

    switch r.Method {
    case "get":
        req, err := http.NewRequest("GET", reqUrl, nil)

        if err != nil {
            panic(err)
        }

    case "post":
        form := url.Values{}
        for k,v := range r.Data{
            form.Add(k, v)
        }

        req, _ := http.NewRequest("POST", reqUrl, strings.NewReader(form.Encode()))
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    }

    req.Host = Config.ServiceHost

    hc := http.Client{}
    resp, _ := hc.Do(req)
    defer resp.Body.Close()

    r.R <- ioutil.ReadAll(resp.Body)
}

