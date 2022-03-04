package main

import "fmt"

func main() {
	s := make([]int, 5) //[0,0,0,0,0]
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	//[0,0,0,0,0,1,2,3]

	v := make([]int,0) //[]
	v = append(v,1,2,3,4)
	fmt.Println(v)
	//[1,2,3,4]
}