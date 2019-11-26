package main

import (
	"fmt"
	"time"
)

func main(){
	// Here is a timeout operation using select.
	// res: = <-c1 waits for the result, <-Time.
	// After waits for the timeout value to be sent after 1 second.
	// Since select processes the first prepared receive operation by default,
	// if this operation exceeds the allowed 1 second,
	// a timeout case will be executed.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	/////////////////////////////////////
	// If I allow a longer timeout of 3 seconds,
	// I will successfully receive the value from c2 and print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <- c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
// timeout 1
// result 2