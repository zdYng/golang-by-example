package main

import "fmt"

func main(){

	s := make([] string, 3)
	fmt.Println("emp: ", s)
	// emp:  [  ]
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	// set:  [a b c]
 	// get:  c

	fmt.Println("len: ", len(s))
	// len:  3
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("app: ", s)
	// app:  [a b c d e f]
	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("copy: ",c)
	// copy:  [a b c d e f]
	lis2 := s[2:5]
	fmt.Println("sli1: ", lis2)
	// sli1:  [c d e]
	lis2 = s[:5]
	fmt.Println("sl2: ", lis2)
	// sl2:  [a b c d e]
	lis2 = s[2:]
	fmt.Println("sl3: ", lis2)
	// sl3:  [c d e f]
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: " , t)
	// dcl:  [g h i]
	twoD := make([][]int, 3)
	for i:=0;i<3;i++{
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j :=0;j<innerLen;j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ",twoD)
	// 2d:  [[0] [1 2] [2 3 4]]

}
