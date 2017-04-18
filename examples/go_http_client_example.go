package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func post(){
    urlStr := "http://127.0.0.1:7007/getBook"
    apiHost := "localhost"

    fmt.Println("URL:>", urlStr)

    form := url.Values{}
    form.Add("bookId", "321")

    hc := http.Client{}
    req, _ := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Host = apiHost

    resp, _ := hc.Do(req)
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

func get() {
	apiUrl := "http://127.0.0.1:7007/bookId?bookId=123"
	fmt.Println("URL:>", apiUrl)

	req, err := http.NewRequest("GET", apiUrl, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	get()
	post()
}
