package main

/**
1.关于字符串连接，下面语法正确的是？
A. str := ‘abc’ + ‘123’
B. str := “abc” + “123”
str := ‘123’ + “abc”
D. fmt.Sprintf(“abc%d”, 123)
参考答案及解析：BD。知识点：字符串连接。除了以上两种连接方式，还有 strings.Join()、buffer.WriteString()等。
 */

/**
const (
     x = iota
     _
     y
     z = "zz"
     k
     p = iota
 )

func main()  {
    fmt.Println(x,y,z,k,p)
}
参考答案及解析：编译通过，输出：0 2 zz zz 5。知识点：iota 的使用。给大家贴篇文章，讲的很详细
https://www.cnblogs.com/zsy/p/5370052.html
 */

/**
3.下面赋值正确的是()
A. var x = nil
B. var x interface{} = nil
C. var x string = nil
D. var x error = nil
参考答案及解析：BD。知识点：nil 值。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 D 选项的 error 类型，它是一种内置接口类型，看下方贴出的源码就知道，所以 D 是对的。

type error interface {
    Error() string
}
 */
