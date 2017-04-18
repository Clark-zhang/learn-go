package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "encoding/json"
    "bytes"
)


type book struct {
    Id string `json:"bookId"`
    Name string `json:"bookName"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func GetBookId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    queryValues := r.URL.Query()

    fmt.Fprintf(w, "bookId: %s \n", queryValues.Get("bookId"))
}

func GetBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
    r.ParseForm()

    bookId := r.Form.Get("bookId")

    var buffer bytes.Buffer
    buffer.WriteString("The bookName of bookId ")
    buffer.WriteString(bookId)
    bookName := buffer.String()

    book := book{Id: bookId, Name: bookName}

    //print struct
    // fmt.Printf("%+v\n", book)

    bookMarshaled, err := json.Marshal(book)
    if err != nil {
        fmt.Println(err)
        return
    }

    //print in server side
    // fmt.Printf("%s\n" , bookMarshaled)

    //response
    fmt.Fprintf(w, "%s\n", string(bookMarshaled))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
    router.GET("/bookId", GetBookId)
    router.POST("/getBook", GetBook)

    log.Fatal(http.ListenAndServe(":7007", router))
}