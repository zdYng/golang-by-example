package main

import (
	"fmt"
	"math"
)

const a string = "constant"

func main(){
	fmt.Println(a)

	const n = 50000000

	const d = 3e11/n

	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}

// constant
// 6000
// 6000
// 0.8256467432733234
