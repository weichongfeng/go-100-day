package main

/**
1.下面的代码有什么问题？
func main() {
    const x = 123
    const y = 1.23
    fmt.Println(x)
}
参考答案及解析：编译可以通过。知识点：常量。常量是一个简单值的标识符，在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。
 */

/**
2.下面代码输出什么？
const (
     x uint16 = 120
     y
     s = "abc"
     z
 )
 func main() {
    fmt.Printf("%T %v\n", y, y) uint16 120
    fmt.Printf("%T %v\n", z, z) string abc"
}
参考答案及解析：知识点：常量。

输出：

    uint16 120
    string abc
常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
 */

/**
3.下面代码有什么问题？
func main() {
    var x string = nil
    if x == nil {
        x = "default"
    }
}
参考答案及解析：将 nil 分配给 string 类型的变量。这是个大多数新手会犯的错误。修复代码：

func main() {
    var x string //defaults to "" (zero value)

    if x == "" {
        x = "default"
    }
}
引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
 */
