package controller

import(
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
    cs_mysql "github.com/Clark-zhang/learn-go/comment_system/storage/mysql"
    "github.com/tomasen/realip"
)

/*
@todo
1. data should be a map
   use <T> (interface) to make common response type
2. reduce duplicate of code
*/

type GetCommentsR struct{
    Code int `json:"code"`
    Data []cs_mysql.Comment `json:"data"`
}

func GetComments(w http.ResponseWriter, r *http.Request) {
    qv := r.URL.Query()
    index, _:= strconv.Atoi(qv.Get("pageIndex"))
    size, _ := strconv.Atoi(qv.Get("pageSize"))
    event, _ := strconv.Atoi(qv.Get("eventId"))

    comments, err := cs_mysql.Comments_Select(DB, int(event), int(index), int(size))

    if err != nil{
        ret, _ := json.Marshal(GetCommentsR{Code: 1})
        fmt.Fprintf(w, "%s", string(ret))
    }else{
        ret, _ := json.Marshal(GetCommentsR{
                Code: 0,
                Data: comments,
            })
        fmt.Fprintf(w, "%s", string(ret))
    }
}

func AddComment(w http.ResponseWriter, r *http.Request){
    r.ParseForm()

    pid, _ := strconv.ParseUint(r.Form.Get("pid"), 10, 64)
    fpid, _ := strconv.ParseUint(r.Form.Get("fpid"), 10, 64)
    eid, _ := strconv.ParseUint(r.Form.Get("eid"), 10, 64)
    uid, _ := strconv.ParseUint(r.Form.Get("uid"), 10, 64)
    c := r.Form.Get("comment")

    comment := &cs_mysql.Comment{
        Pid: pid,
        Fpid: fpid,
        Eid: eid,
        Uid: uid,
        Content: c,
        Ip: realip.RealIP(r),
    }

    err := comment.Add(DB)

    if err != nil{
        ret, _ := json.Marshal(GetCommentsR{Code: 1})
        fmt.Fprintf(w, "%s", string(ret))
    }else{
        ret, _ := json.Marshal(GetCommentsR{Code: 0})
        fmt.Fprintf(w, "%s", string(ret))
    }
}