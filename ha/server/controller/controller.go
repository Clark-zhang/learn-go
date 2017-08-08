package controller

import(
    "fmt"
    "net/http"
    "github.com/go-redis/redis"
    syncByHttp "github.com/Clark-zhang/learn-go/ha/sync/http"
    // "encoding/json"
    // "bytes"
)

func SetRedisData(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    key := r.Form.Get("key")
    value := r.Form.Get("value")

    client := redis.NewClient(&redis.Options{
        Addr:     Conf["storage"],
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := client.Ping().Result()
    // fmt.Println(pong, err)
    // Output: PONG <nil>

    err = client.Set(key, value, 0).Err()
    if err != nil {
        panic(err)
    }else{
        fmt.Fprintf(w, "receive key:%s value:%s", key, value)
    }

    //@todo decorator pattern?
    if(r.Header.Get("ha-sync") != "true"){
        syncByHttp.SyncByHttp(r)
    }
}