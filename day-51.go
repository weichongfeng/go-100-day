package main

/**
1.下面的代码能否正确输出？
func main() {
    var fn1 = func() {}
    var fn2 = func() {}
    if fn1 != fn2 {
       println("fn1 not equal fn2")
    }
}
参考答案及解析：编译错误

    invalid operation: fn1 != fn2 (func can only be compared to nil)
函数只能与 nil 比较。
 */

/**
2.下面代码输出什么？
type T struct {
    n int
}
func main() {
    m := make(map[int]T)
    m[0].n = 1
    fmt.Println(m[0].n)
}
A. 1

B. compilation error

参考答案及解析：B。编译错误：

    cannot assign to struct field m[0].n in map
map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。

修复代码：

type T struct {
    n int
}
func main() {
    m := make(map[int]T)
    t := T{1}
    m[0] = t
    fmt.Println(m[0].n)
}

 */
