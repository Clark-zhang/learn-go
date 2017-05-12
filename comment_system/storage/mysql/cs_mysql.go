/*
todo:
    mysql connection pool
    atomic,mutex?

*/
package cs_mysql

import (
    "database/sql"
     _ "github.com/go-sql-driver/mysql"
     "fmt"
     )

type Relation struct {
    cId uint64
    pId uint64
    fpId uint64
    eId uint64
}

type Comment struct {
    cId uint64
    pId uint64
    fpId uint64
    eId uint64
}

type Comments struct{
    comments *[]Comment
}

func (c Comment) Create() (Err){

}

func (c Comment) Update() (Err){

}

func test(){
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:7008)/cs?charset=utf8")
    if err != nil{
        panic(err.Error())
    }
}