package controller

import(
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "bytes"
)

type book struct {
    Id string `json:"bookId"`
    Name string `json:"bookName"`
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