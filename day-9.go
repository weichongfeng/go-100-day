package main

/**
1.关于channel，下面语法正确的是()
A. var ch chan int

B. ch := make(chan int)

C. <- ch

D. ch <-

参考答案及解析：ABC。A、B都是声明 channel；C 读取 channel；写 channel 是必须带上值，所以 D 错误。
 */

/**
2.下面这段代码输出什么？
type person struct {
    name string
}

func main() {
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
A.0

B.1

C.Compilation error

参考答案及解析：A。打印一个 map 中不存在的值时，返回元素类型的零值。这个例子中，m 的类型是 map[person]int，
因为 m 中不存在 p，所以打印 int 类型的零值，即 0。

 */

/**
3.下面这段代码输出什么？
func hello(num ...int) {
    num[0] = 18
}

func main() {
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
A.18

B.5

C.Compilation error

参考答案及解析：18。知识点：可变函数。
 */
