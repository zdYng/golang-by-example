package main

import "fmt"

func intSeq() func() int{
	i := 0
	return func() int {
		i ++
		return i
	}
}

func main(){
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 1
	// 2
	// 3
	newInts := intSeq()
	fmt.Println(newInts())
	// 1
}
