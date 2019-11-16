package main

import "fmt"

type data struct {
	i int
	s string
}

func main() {
	/*data2 := [3]string{"aa","bb","cc"}
	for i := range data2{
		fmt.Println(i)
	}
	var data2Copy [3]*string
	for i,s := range data2{
		fmt.Println(i, s)
		data2Copy[i] = &s
	}

	for i := range data2{
		fmt.Println(i)
	}*/
	b := sayHello()
	fmt.Println(b)
}

func sayHello() *int {
	a := 100
	fmt.Println(a, &a)
	return &a
}
