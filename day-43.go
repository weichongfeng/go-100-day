package main

/**
import (
    "fmt"
    "log"
    "time"
)
func main() {
}
参考答案及解析：导入的包没有被使用。如果引入一个包，但是未使用其中如何函数、接口、结构体或变量的话，代码将编译失败。

如果你真的需要引入包，可以使用下划线操作符，_，来作为这个包的名字，从而避免失败。下划线操作符用于引入，但不使用。

我们还可以注释或者移除未使用的包。

修复代码：

import (
    _ "fmt"
    "log"
    "time"
)
var _ = log.Println
func main() {
    _ = time.Now
}
引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
 */

/**
2.下面代码输出什么？
func main() {
    x := interface{}(nil)
    y := (*int)(nil)
    a := y == x false
    b := y == nil true
    _, c := x.(interface{}) true
    println(a, b, c)
}
A. true true false

B. false true true

C. true true true

D. false true false

参考答案及解析：D。知识点：类型断言。类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。
编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败
 */

/**
3.下面代码有几处错误的地方？请说明原因。
func main() {

    var s []int
    s = append(s,1)

    var m map[string]int
    m["one"] = 1
}
参考答案及解析：有 1 出错误，不能对 nil 的 map 直接赋值，需要使用 make() 初始化。但可以使用 append() 函数对为 nil 的 slice 增加元素。

修复代码：

func main() {
    var m map[string]int
    m = make(map[string]int)
    m["one"] = 1
}
 */
