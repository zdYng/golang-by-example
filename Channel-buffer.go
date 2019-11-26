package main

import "fmt"

func main(){
	messages := make(chan string , 3) // 最多3个缓存
	messages<- "buffered"
	messages<- "channel"
	messages<- "lastt"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<- messages)

	fmt.Println("######################")

	messages_2 := make(chan string , 2) // 最多2个缓存
	messages_2<- "one"
	messages_2<- "two"
	fmt.Println(<-messages_2)
	fmt.Println(<-messages_2)
}

// buffered
// channel
// lastt
// ######################
// one
// two