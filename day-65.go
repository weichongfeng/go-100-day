package main

import "fmt"

/**
1.flag 是 bool 型变量，下面 if 表达式符合编码规范的是？
A. if flag == 1
B. if flag
C. if flag == false
D. if !flag
参考答案及解析：BCD。
 */

/**
2.下面的代码输出什么，请说明？
func main() {
    defer func() {
        fmt.Print(recover())
    }()
    defer func() {
        defer func() {
            fmt.Print(recover())
        }()
        panic(1)
    }()
    defer recover()
    panic(2)
}
参考答案及解析：12。相关知识点请看 第64天 题目解析。
 */

func main() {
	defer func() {
		fmt.Println(1111)
		fmt.Print(recover())
	}()
	defer func() {
		//defer func() {
		//	fmt.Print(recover())
		//}()
		defer fmt.Println(recover(),222)
		panic(1)
	}()
	defer recover()
	panic(2)
}
