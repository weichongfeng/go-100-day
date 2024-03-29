package main
/**
1.下面这段代码输出什么？
func main() {
    var k = 1
    var s = []int{1, 2}
    k, s[k] = 0, 3
    fmt.Println(s[0] + s[1]) 4
}
参考答案及解析：4。知识点：多重赋值。

多重赋值分为两个步骤，有先后顺序：

计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；

赋值；

所以本例，会先计算 s[k]，等号右边是两个表达式是常量，所以赋值运算等同于 k, s[1] = 0, 3。
 */

/**
2.下面代码输出什么？
func main() {
    var k = 9
    for k = range []int{} {}
    fmt.Println(k)

    for k = 0; k < 3; k++ {
    }
    fmt.Println(k)


    for k = range (*[3]int)(nil) {
    }
    fmt.Println(k)
}
参考答案及解析：932。
 */
