package main

import (
	"fmt"
	"time"
)

func main(){
	requests := make(chan int, 5)
	for i := 1;i<=5;i++{
		requests<- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond*200)

	for req := range requests{
		<- limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLimiter := make(chan time.Time,3)

	for i := 0; i<3;i++{
		burstyLimiter<-time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200){
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i<=5; i++{
		burstyRequests<-i
	}
	close(burstyRequests)
	for req := range burstyRequests{
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
// request 1 2019-11-26 10:34:26.4004105 +0800 CST m=+0.205552701
//request 2 2019-11-26 10:34:26.6000255 +0800 CST m=+0.405167701
//request 3 2019-11-26 10:34:26.8005154 +0800 CST m=+0.605657601
//request 4 2019-11-26 10:34:27.0001528 +0800 CST m=+0.805295001
//request 5 2019-11-26 10:34:27.2017837 +0800 CST m=+1.006925901

//request 1 2019-11-26 10:34:27.2017837 +0800 CST m=+1.006925901
//request 2 2019-11-26 10:34:27.2017837 +0800 CST m=+1.006925901
//request 3 2019-11-26 10:34:27.2017837 +0800 CST m=+1.006925901
//request 4 2019-11-26 10:34:27.4034427 +0800 CST m=+1.208584901
//request 5 2019-11-26 10:34:27.6031026 +0800 CST m=+1.408244801