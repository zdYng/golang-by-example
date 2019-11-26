package main

import "fmt"

func ping(pings chan<- string, msg string){
	pings <- msg
}
// The ping function defines a channel that is only allowed to send data.
// Attempts to receive data using this channel will result in a compile-time error.

func pong(pings <- chan string, pongs chan<- string){
	msg := <- pings // send
	pongs <- msg  // receive
}
// The pong function allows channels (pings) to receive data and
// another channel (pongs) to send data.


func main(){
	pings := make(chan string , 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<- pongs)

}
// passed message