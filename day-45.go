package main

/**
1.下面代码有什么错误？
func main() {
    one := 0
    one := 1
}
one = 1
参考答案及解析：变量重复声明。不能在单独的声明中重复声明一个变量，但在多变量声明的时候是可以的，但必须保证至少有一个变量是新声明的。

修复代码：

func main() {
    one := 0
    one, two := 1,2
    one,two = two,one
}
引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
 */

/**
2.下面代码有什么问题？
func main() {
    x := []int{
        1,
        2
    }
    _ = x
}
参考答案及解析：编译错误，第四行代码没有逗号。用字面量初始化数组、slice 和 map 时，最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错。

修复代码：

func main() {
    x := []int{    // 多行
        1,
        2,
    }
    x = x

    y := []int{3,4,} // 一行 no error
    y = y
}
引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
 */

/**
3.下面代码输出什么？
func test(x byte)  {
    fmt.Println(x)
}

func main() {
    var a byte = 0x11
    var b uint8 = a
    var c uint8 = a + b
    test(c)
}
参考答案及解析：34。与 rune 是 int32 的别名一样，byte 是 uint8 的别名，别名类型无序转换，可直接转换。
*/
