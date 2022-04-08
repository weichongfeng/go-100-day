package main

/**
1.下面的代码有什么问题？
package main

import "fmt"

func main() {
     s := make([]int, 3, 9)
     fmt.Println(len(s))
     s2 := s[4:8]
     fmt.Println(len(s2))
}
参考答案及解析：代码没问题，输出 3 4。从一个基础切片派生出的子切片的长度可能大于基础切片的长度。
假设基础切片是 baseSlice，使用操作符 [low,high]，有如下规则：0 <= low <= high <= cap(baseSlice)，只要上述满足这个关系，
下标 low 和 high 都可以大于 len(baseSlice)。

引自：《Go语言101》
 */

/**
2.下面代码输出什么？
type N int

func (n N) test(){
     fmt.Println(n)
 }

 func main()  {
     var n N = 10
     p := &n

    n++
    f1 := n.test

    n++
    f2 := p.test

    n++
    fmt.Println(n)

    f1()
    f2()
}
 */
