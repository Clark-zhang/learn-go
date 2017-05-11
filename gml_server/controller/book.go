package controller

import(
    "fmt"
    "net/http"
    "encoding/json"
    // "bytes"
)

type book struct {
    Id string `json:"bookId"`
    Name string `json:"bookName"`
}

func GetBookId(w http.ResponseWriter, r *http.Request) {
    queryValues := r.URL.Query()

    fmt.Fprintf(w, "bookId: %s \n", queryValues.Get("bookId"))
}

func GetBook(w http.ResponseWriter, r *http.Request){
    r.ParseForm()

    bookId := r.Form.Get("bookId")

    // var buffer bytes.Buffer
    // buffer.WriteString("The bookName of bookId ")
    // buffer.WriteString(bookId)
    // bookName := buffer.String()
    bookName := "The bookName of bookId " + bookId

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