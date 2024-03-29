package main

/**
1.关于协程，下面说法正确是（）
A. 协程和线程都可以实现程序的并发执行；

B. 线程比协程更轻量级；

C. 协程不存在死锁问题；

D. 通过 channel 来进行协程间的通信；

参考答案及解析：AD。
 */

/**
2.关于循环语句，下面说法正确的有（）
A. 循环语句既支持 for 关键字，也支持 while 和 do-while；

B. 关键字 for 的基本使用方法与 C/C++ 中没有任何差异；

C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；

D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；
参考答案及解析：CD。
 */

/**
3.下面代码输出正确的是？
func main() {
    i := 1
    s := []string{"A", "B", "C"}
    i, s[i-1] = 2, "Z"
    fmt.Printf("s: %v \n", s)
}
A. s: [Z,B,C]

B. s: [A,Z,C]

3.下面代码输出正确的是？
func main() {
    i := 1
    s := []string{"A", "B", "C"}
    i, s[i-1] = 2, "Z"
    fmt.Printf("s: %v \n", s)
}
A. s: [Z,B,C]

B. s: [A,Z,C]
 */
