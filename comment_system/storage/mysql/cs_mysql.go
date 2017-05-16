package cs_mysql

import (
    "database/sql"
     _ "github.com/go-sql-driver/mysql"
     "log"
     )

type Relation struct {
    Cid uint64
    Pid uint64
    Fpid uint64
    Eid uint64
}

type Comment struct {
    Cid uint64 `json:"cid"`
    Pid uint64 `json:"pid"`
    Fpid uint64 `json:"fpid"`
    Eid uint64 `json:"-"`
    Uid uint64 `json:"uid"`
    Content string `json:"comment"`
    Ip string `json:"-"`
    CreatedAt string `json:"created"`
}

type Comments []Comment


func (c Comment) Add(db *sql.DB) (err error){
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }

    stmt, _ := db.Prepare("INSERT INTO comments(e_id, user_id, content, ip, created_at) VALUES(?, ?, ?, ?, NOW())")
    res, err := stmt.Exec(c.Eid, c.Uid, c.Content, c.Ip)
    if err != nil {
        log.Fatal(err)
        tx.Rollback()
    }
    lastId, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
        tx.Rollback()
    }

    stmt, _ = db.Prepare("INSERT INTO comments_relationship(c_id, p_id, fp_id, e_id) VALUES(?, ?, ?, ?)")
    res, err = stmt.Exec(lastId, c.Pid, c.Fpid, c.Eid)
    if err != nil {
        log.Fatal(err)
        tx.Rollback()
    }

    tx.Commit()

    return
}

func (c Comment) Delete(db *sql.DB) (err error){
    return
}

func (c Comment) Update(db *sql.DB) (comment Comment, err error){
    return
}

func Comments_Select(db *sql.DB, Eid int, pageIndex int, pageSize int) (comments []Comment, err error){
    from := (pageIndex - 1) * pageSize

    rows, err := db.Query("select id, user_id, content, created_at from comments where e_id = ? limit ?, ?", Eid, from, pageSize)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // comments = []Comment{}
    c := Comment{}
    for rows.Next() {
        err := rows.Scan(&c.Cid, &c.Uid, &c.Content, &c.CreatedAt)
        if err != nil {
            log.Fatal(err)
        }
        comments = append(comments, c)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    return
}