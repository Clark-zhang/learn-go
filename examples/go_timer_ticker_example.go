package main

import (
    "time"
    "fmt"
    "sync"
    "github.com/go-redis/redis"
)

func refreshTime(c *redis.Client) {
    ticker := time.NewTicker(time.Second * 5)
    for t := range ticker.C{
        c.Set("time", t.String(), 0).Err()
    }
}

func refreshPrefixTime(c *redis.Client) {
    ticker := time.NewTicker(time.Second * 5)
    for _ = range ticker.C{
        val, err := c.Get("time").Result()
        if err != nil {
            panic(err)
        }

        c.Set("prefix-time", "prefix " + string(val), 0).Err()
    }
}

func refreshCache() {
    client := redis.NewClient(&redis.Options{
            Addr:     "localhost:6379",
            Password: "", // no password set
            DB:       0,  // use default DB
    })
    var wg sync.WaitGroup
    //add one channel to wait, but it will never receive the channel
    //so it will go into infinite loop
    //to make sure refresh go routines will not exit
    wg.Add(1)

    go refreshTime(client)
    go refreshPrefixTime(client)

    wg.Wait()
}

func main(){
    refreshCache()
}

func simpleExample() {
  ticker := time.NewTicker(time.Millisecond * 500)
  go func() {
  for t := range ticker.C {
  fmt.Println("Tick at", t)
  }
  }()
  time.Sleep(time.Millisecond * 1500)
  ticker.Stop()
  fmt.Println("Ticker stopped")
}

