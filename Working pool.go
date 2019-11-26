package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <- chan int, results chan <- int){
	for j := range jobs{
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main(){
	jobs := make(chan int , 100)
	results := make(chan int, 100)

	for w := 1; w<= 3; w++{
		go workers(w, jobs, results)
	}
	for j := 1; j<= 9; j ++{
		jobs<- j
	}
	close(jobs)
	for a := 1; a <= 9; a++{
		<- results
	}
}
// worker 1 processing job 2
// worker 2 processing job 3
// worker 3 processing job 1
// worker 3 processing job 4
// worker 1 processing job 5
// worker 2 processing job 6
// worker 2 processing job 7
// worker 1 processing job 8
// worker 3 processing job 9

