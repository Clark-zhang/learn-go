package main

import (
    "fmt"
    "github.com/jinzhu/configor"
    "github.com/Clark-zhang/learn-go/gml/request"
)

func main() {

    req := request.Request{
        Url: "/bookId?bookId=123"
        Method: "get"
        R: make(chan string)
    }
    go req.MakeRequest();

    fmt.Println(req.R)
}