//
//
//

package main

import (
    "fmt"
    "time"
)


func dtmNowFmt() string   {
   
    return time.Now().Format(time.RFC3339Nano)
}


func main() {

    //  loop, time acquisition, and printf duration research

    for j := 0; j <= 10; j++ {
            fmt.Printf("loop     : exec       : %s\n",       dtmNowFmt() )
    }




    //  these three "channels" provide most of the logic within this program

                //  this channel receives content and quits after "duration"
    chTimer     :=  time.NewTimer(3*time.Second).C 
    
                //  this channel receives content after every "duration"
    chTicker    :=  time.NewTicker(time.Millisecond * 700).C
 
                //  this channel lives (keeps us live), until we choose to stop
    chDone      :=  make(chan bool)


    go func() { //  this goroutine sleeps, waking to end the Done chan 
        time.Sleep(time.Second * 9)
        chDone <- true
    }   ()
    

    for {

        select {

        case <- chTimer:
            fmt.Printf("chTimer  : expires    : %s\n",       dtmNowFmt() )

        case <- chTicker:
            fmt.Printf("chTicker : fires      : %s\n",       dtmNowFmt() )

        case <- chDone:
            fmt.Printf("chDone   : terminates : %s\n",       dtmNowFmt() )
            return
        }
    }

}
