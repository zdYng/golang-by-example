package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Alice", 30})
	fmt.Println(person{"Fred", 22})
	fmt.Println(person{"Ann", 40})
	s := person{"ann", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
// {Bob 20}
// {Alice 30}
// {Fred 22}
// {Ann 40}
// ann
// 50
// 51