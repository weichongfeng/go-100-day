package main

/**
1.下面代码输出什么？
func main() {
    str := "hello"
    str[0] = 'x'
    fmt.Println(str)
}
A. hello

B. xello

C. compilation error

参考代码及解析：C。知识点：常量，Go 语言中的字符串是只读的。
 */

/**
2.下面代码输出什么？
func incr(p *int) int {
    *p++
    return *p
}

func main() {
    p :=1
    incr(&p)
    fmt.Println(p)
}
A. 1

B. 2

C. 3

参考答案及解析：B。知识点：指针，incr() 函数里的 p 是 *int 类型的指针，指向的是 main() 函数的变量 p 的地址。
第 2 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果。
 */

/**
3.对 add() 函数调用正确的是（）
func add(args ...int) int {

    sum := 0
    for _, arg := range args {
        sum += arg
    }
    return sum
}
A. add(1, 2)

B. add(1, 3, 7)

C. add([]int{1, 2})

D. add([]int{1, 3, 7}…)

参考答案及解析：ABD。知识点：可变函数。
 */
