package request

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "sync"
    "github.com/Clark-zhang/learn-go/gml/config"
)

//http://blog.narenarya.in/concurrent-http-in-go.html

type Request struct{
    Url string
    Method string
    Data map[string]string
    R chan string
}

func (r Request) MakeRequest(wg *sync.WaitGroup){
    wg.Add(1)
    defer wg.Done()

    Config := config.GetConfig()

    reqUrl := Config.ServiceUrl + r.Url

    var req *http.Request

    switch r.Method {
    case "get":
        req, _ = http.NewRequest("GET", reqUrl, nil)

    case "post":
        form := url.Values{}
        for k,v := range r.Data{
            form.Add(k, v)
        }

        req, _ = http.NewRequest("POST", reqUrl, strings.NewReader(form.Encode()))
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    }

    req.Host = Config.ServiceHost

    hc := http.Client{}
    resp, err := hc.Do(req)

    if err != nil {
        r.R <- "false"
        return
    }else{
        defer resp.Body.Close()
    }

    body, _ := ioutil.ReadAll(resp.Body)

    r.R <- string(body)
}

