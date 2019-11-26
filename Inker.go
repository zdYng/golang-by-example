package main

import (
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTimer(time.Millisecond * 500)
	go func() {
		for t := range ticker.C{
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// Tick at 2019-11-26 10:17:04.3834284 +0800 CST m=+0.505199101
// Ticker stopped