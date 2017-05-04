package main

import (
    "fmt"
    "sync"
    "github.com/Clark-zhang/learn-go/gml/request"
)


func sendReq(){
    var wg sync.WaitGroup

    getReq := request.Request{
        Url: "/bookId?bookId=123",
        Method: "get",
        R: make(chan string),
    }
    go getReq.MakeRequest(&wg);

    postReq := request.Request{
        Url: "/getBook",
        Method: "post",
        Data: map[string]string{"bookId": "1"},
        R: make(chan string),
    }
    go postReq.MakeRequest(&wg);

    wg.Wait()

    fmt.Println(<-getReq.R)
    fmt.Println(<-postReq.R)
}

func main(){
    sendReq()
}