package main

/**
1.下面的代码有几处语法问题，各是什么？
package main
import (
    "fmt"
)
func main() {
    var x string = nil
    if x == nil {
        x = "default"
    }
    fmt.Println(x)
}
参考答案及解析：两个地方有语法问题。golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。
 */
/**
2.return 之后的 defer 语句会执行吗，下面这段代码输出什么？
var a bool = true
func main() {
    defer func(){
        fmt.Println("1")
    }()
    if a == true {
        fmt.Println("2")
        return
    }
    defer func(){
        fmt.Println("3")
    }()
}
参考答案及解析：2 1。defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法
 */
