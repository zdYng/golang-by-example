package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	fmt.Println(t)
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i<2;i++{
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}
	t2 := time.Now()
	fmt.Println(t2)

	// Note that concurrent executions from the first and
	// second Sleeps only took about two seconds to run
}

// 2019-11-26 09:37:13.9882916 +0800 CST m=+0.004985901
// received one
// received two
// 2019-11-26 09:37:16.0242861 +0800 CST m=+2.040980401