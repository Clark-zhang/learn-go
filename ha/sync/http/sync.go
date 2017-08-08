package syncByHttp

import(
    "fmt"
    "net/http"
    netUrl "net/url"
    "strings"
)

func SyncByHttp(req *http.Request) {
    url := req.URL
    url.Host = Conf["syncAddr"]

    req.ParseForm()

    //@todo check if it's support file upload
    form := netUrl.Values{}
    for key, vals := range req.Form {
        for _, val := range vals {
            form.Add(key, val)
        }
    }

    proxyReq, err := http.NewRequest(req.Method, "http:" + url.String(), strings.NewReader(form.Encode()))
    if err != nil {
        // handle error
    }

    //@todo support multi instances
    proxyReq.Header.Set("ha-sync", "true")
    proxyReq.Header.Set("Content-Type", "multipart/form-data")

    for header, values := range req.Header {
        for _, value := range values {
            proxyReq.Header.Set(header, value)
        }
    }

    client := &http.Client{}
    proxyRes, err := client.Do(proxyReq)

    if err != nil {
       fmt.Println(err)
    }else if proxyRes.StatusCode != 200{
        //log
        fmt.Println("sync response not 200")
        fmt.Println(err)
    }
}