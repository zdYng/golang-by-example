package main

import "fmt"

// 如果你有一个含有多个值的 slice，
// 想把它们作为参数 使用，
// 你要这样调用 func(slice...)。
func sum(nums ...int){
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main(){
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1,2,3,4}
	sum(nums...)
}
