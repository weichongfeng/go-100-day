package main

/**
1.关于 cap() 函数的适用类型，下面说法正确的是()
A. array

B. slice

C. map

D. channel

参考答案及解析：ABD。知识点：cap()，cap() 函数不适用 map。
 */

/**
2.下面这段代码输出什么？
func main() {
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
A. nil

B. not nil

C. compilation error

参考答案及解析：A。当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。
 */

/**
3.下面这段代码输出什么？
func main() {
    s := make(map[string]int)
    delete(s, "h")
    fmt.Println(s["h"])
}
A. runtime panic

B. 0

C. compilation error

参考答案及解析：B。删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。
 */
