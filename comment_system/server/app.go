package cs_server

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/Clark-zhang/learn-go/comment_system/server/router"
    "github.com/Clark-zhang/learn-go/comment_system/server/controller"
    "log"
    "net/http"
    "fmt"
)

var DB *sql.DB

func Run(){
    DB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:7008)/cs?charset=utf8")
    if err != nil {
        log.Fatalf("Error on initializing database connection: %s", err.Error())
    }

    DB.SetMaxIdleConns(100)

    if err != nil {
        log.Fatalf("Error on opening database connection: %s", err.Error())
    }

    controller.DB = DB

    router.New()
    fmt.Println("Sever start on 8007")
    log.Fatal(http.ListenAndServe(":8007", nil))
}