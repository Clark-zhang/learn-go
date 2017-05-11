//docker run --name cs -e MYSQL_ROOT_PASSWORD=root -d -p 3306:7008 mysql

/*
CREATE TABLE IF NOT EXISTS `comments` (
    `id` bigint(30) NOT NULL auto_increment,
    `e_id` bigint(30),
    `p_id` bigint(30),
    `fp_id` bigint(30),
    `user_id` int(2),
    `content` varchar(1024),
    `ip` varchar(25),
    `created_at` datetime,
    PRIMARY KEY  (`id`)
);
*/

package main

import (
    "database/sql"
     _ "github.com/go-sql-driver/mysql"
     "fmt"
     )


func main(){
    db, err := sql.Open("mysql", "root:root@tcp(localhost:7008)/cs?charset=utf8")
    if err != nil{
        panic(fmt.Println("connect error"))
    }

    stmt, _ := db.Prepare("INSERT qd_en_review_report SET user_id=?,content=?")

    res, _ := stmt.Exec("1", "comment content")

    id, _ := res.LastInsertId()

    fmt.Println(id)
}