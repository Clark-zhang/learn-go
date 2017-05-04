package timingWheel

import(
    "time"
)

//todo DEBUG
//<-ch will get variable from channel and put it to non-value
//is there a way to share variables between routines and not reset it?

func TimingWheel(wCh chan [][]string, cSlotCh chan int){
    ticker := time.NewTicker(time.Second * 5)
    func() {
        for _ = range ticker.C {
            //todo notify the connection

            wheel := <-wCh
            CSlot := <-cSlotCh

            wheel[CSlot] = wheel[CSlot][:0]
            wCh <- wheel


            if CSlot == len(wheel){
                CSlot = 0
            }else{
                CSlot++
            }
            cSlotCh <- CSlot
        }
    }()
}

func MockConnections(wCh chan [][]string, cSlotCh chan int){
    ticker := time.NewTicker(time.Second * 3)
    func() {
        for t := range ticker.C {
            wheel := <- wCh
            cSlot := <- cSlotCh
            wheel[cSlot] = append(wheel[cSlot], t.String() + "1")
            wheel[cSlot] = append(wheel[cSlot], t.String() + "2")
            wCh <- wheel
            cSlotCh <- cSlot
        }
    }()
}