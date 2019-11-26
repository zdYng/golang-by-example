package main

import "fmt"

func main(){

	i := 1
	for i <=3{
		fmt.Println(i)
		i += 1
	}

	// 1
	//2
	//3


	for j := i ; j<=9 ; j++{
		fmt.Println(j)
	}

	// 4
	// 5
	// 6
	// 7
	// 8
	// 9

	c := 0
	for {
		c++
		fmt.Println("while")
		if c==4{
			break
		}
	}

	// while
	// while
	// while
	// while

	for n:= 0 ; n<14;n++{
		if n%2==0{
			continue
		}
		fmt.Println(n)
	}

	// 1
	// 3
	// 5
	// 7
	// 9
	// 11
	// 13

}
