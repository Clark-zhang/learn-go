package main

import(
    "time"
    "fmt"
    "github.com/Clark-zhang/learn-go/examples/timing_wheel"
)

//todo DEBUG

func main(){
    wheel := make([][]string, 8)
    wCh := make(chan [][]string)
    wCh <- wheel

    sSlotCh := make(chan int)
    sSlotCh <- 0

    go timingWheel.MockConnections(wCh, sSlotCh)
    go timingWheel.TimingWheel(wCh, sSlotCh)

    ticker := time.NewTicker(time.Second * 3)
    func() {
        for t := range ticker.C {
            target := <- wCh
            fmt.Println(target, " at ", t)
        }
    }()
}
