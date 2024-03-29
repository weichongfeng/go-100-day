package main
/**
1.下面选项正确的是？
A. 类型可以声明的函数体内；
B. Go 语言支持 ++i 或者 –i 操作；
C. nil 是关键字；
D. 匿名函数可以直接赋值给一个变量或者直接执行；
参考答案及解析：AD。
 */

/**
2.下面的代码输出什么？
func F(n int) func() int {
    return func() int {
        n++
        return n
    }
}

func main() {
    f := F(5)
    defer func() {
        fmt.Println(f())
    }()
    defer fmt.Println(f())
    i := f()
    fmt.Println(i)
}
参考答案及解析：768。知识点：匿名函数、defer()。defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。

引自《Go语言101》
 */
