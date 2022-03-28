package main

/**
1.下面的代码有什么问题？
type X struct {}
func (x *X) test()  {
   println(x)
}
func main() {
    var a *X
    a.test()
    X{}.test()
}
参考答案及解析：X{} 是不可寻址的，不能直接调用方法。知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。

修复代码：

func main() {
    var a *X
    a.test()    // 相当于 test(nil)
    var x = X{}
    x.test()
}
 */

/**
2.下面代码有什么不规范的地方吗？
func main() {
    x := map[string]string{"one":"a","two":"","three":"c"}
    if v := x["two"]; v == "" {
        fmt.Println("no entry")
    }
}
参考答案及解析：检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。最可靠的操作是使用访问 map 时返回的第二个值。

修复代码如下：

func main() {
    x := map[string]string{"one":"a","two":"","three":"c"}
    if _,ok := x["two"]; !ok {
        fmt.Println("no entry")
    }
}
 */
