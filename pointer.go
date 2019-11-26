package main

import "fmt"

func zeroval(ival int) int{
	ival = 0
	return ival
}

func zeroptr(iptr *int) {
	*iptr = 3
}

func main(){
	i := 1
	fmt.Println("initial: ", i)

	result := zeroval(i)
	fmt.Println("zeroval: ", i)
	fmt.Println("result: ", result)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}
// initial:  1
// zeroval:  1
// result:  0
// zeroptr:  3
// pointer:  0xc00000a0a8
