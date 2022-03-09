package main

/**
下面代码中，x 已声明，y 没有声明，判断每条语句的对错。
    1. x, _ := f() false
    2. x, _ = f() true
    3. x, y := f() true
    4. x, y = f() false

参考答案及解析：错、对、对、错。知识点：变量的声明。1.错，x 已经声明，不能使用 :=；2.对；3.对，
当多值赋值时，:= 左边的变量无论声明与否都可以；4.错，y 没有声明。
 */

/**
2.下面代码输出什么？
func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}

func increaseB() (r int) {
    defer func() {
        r++
    }()
    return r
}

func main() {
    fmt.Println(increaseA()) 0
    fmt.Println(increaseB()) 1
}
A. 1 1

B. 0 1

C. 1 0

D. 0 0

参考答案及解析：B。知识点：defer、返回值。注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。
关于 defer 与返回值的知识点，后面我会写篇文章详细分析，到时候可以看下文章的讲解。
 */

/**
下面代码输出什么？
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
	var a A = Work{3}
    s := a.(Work)
    fmt.Println(s.ShowA()) 13
    fmt.Println(s.ShowB()) 33
}

A. 13 23

B. compilation error

参考答案及解析：A。知识点：类型断言。这道题可以和第 15 天的第三题 和第 16 天的第三题结合起来看

*/