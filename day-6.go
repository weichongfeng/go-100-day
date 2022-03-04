package main

import "fmt"

type MyInt1 int  //定义一个新类型 MyInt1
type MyInt2 = int //定义一个int类型的别名

func main() {
	var i int =0
	var i1 MyInt1 = i  // int类型不能直接赋值个MyInt1类型 可以这样定义 var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i  //MyInt2 本质上还是 int  MyInt2只是int别名
	fmt.Println(i1,i2)
}
