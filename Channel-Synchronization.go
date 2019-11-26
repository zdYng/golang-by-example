package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Print("working ...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	done <- true
}

func main(){
	done := make(chan bool , 1)

	go worker(done)

	<- done
	fmt.Println("no <- done : ....")
	done2 := make(chan bool , 1)

	go worker(done2)
	// If you remove the line <-done from the program,
	// the program will end even before the worker has started

}
// working ...done
// no <- done : ....