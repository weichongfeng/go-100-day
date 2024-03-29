package main

/**
func main() {
    var s1 []int //nil切片
    var s2 = []int{} //空切片
    if __ == nil {
        fmt.Println("yes nil")
    }else{
        fmt.Println("no nil")
    }
}
A. s1

B. s2

C. s1、s2 都可以
 */

/**
2.下面这段代码输出什么？
func main() {
    i := 65
    fmt.Println(string(i))
}
A. A

B. 65

C. compilation error

参考答案及解析：A。UTF-8 编码中，十进制数字 65 对应的符号是 A。
 */

/**
3.下面这段代码输出什么？
type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowA())
    fmt.Println(b.ShowB())
}

参考答案及解析：13 23。知识点：接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，
所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。
 */
