package main

/**
1.下面代码有什么问题？
func main() {
    m := make(map[string]int,2)
    cap(m)
}
参考答案及解析：问题：使用 cap() 获取 map 的容量。1.使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。
2.cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数。

引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
 */
/**
2.下面的代码有什么问题？
func main() {
    var x = nil
    _ = x
}
参考答案及解析：nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。如果不指定变量的类型，
编译器猜不出变量的具体类型，导致编译错误。

修复代码：

func main() {
    var x interface{} = nil
    _ = x
}
 */

/**
3.下面代码能编译通过吗？
type info struct {
    result int
}
func work() (int,error) {
    return 13,nil
}
func main() {
    var data info
    data.result, err := work()
    fmt.Printf("info: %+v\n",data)
}
参考答案及解析：编译失败。

    non-name data.result on left side of :=
不能使用短变量声明设置结构体字段值，修复代码：

func main() {
    var data info
    var err error
    data.result, err = work() //ok
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(data)
}
引自：http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html
 */
