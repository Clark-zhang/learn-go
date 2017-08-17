package main

import (
    "flag"
    "github.com/Clark-zhang/learn-go/ha/server"
)

// go run app.go --storage=localhost:30002 --port=localhost:30001 --syncAddr=localhost:30003
// go run app.go --storage=localhost:30004 --port=localhost:30003 --syncAddr=localhost:30001
//curl localhost:30001/setData -d "key=alicekey1&value=alicevalue1"
//curl localhost:30003/setData -d "key=bobkey1&value=bobvalue1"

//@todo add reverse proxy for fetching data?
//@todo add register center to register/heartbeat new instance?

func main() {
    storage := flag.String("storage", "localhost:30002", "redis addr")
    port := flag.String("port", "localhost:30001", "the port that Alice will run")

    syncAddr := flag.String("syncAddr", "localhost:30003", "Ha friend - Bob")

    flag.Parse()


    c := map[string]string{
        "storage": *storage,
        "port": *port,
        "syncAddr": *syncAddr,
    }

    go server.Run(c)
    select {}
}