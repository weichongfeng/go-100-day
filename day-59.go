package main
/**
1.下面的代码输出什么？
type N int

func (n *N) test(){
    fmt.Println(*n)
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
参考答案及解析：13 13 13。知识点：方法值。当目标方法的接收者是指针类型时，那么被复制的就是指针。

引自：《Go语言学习笔记》· 方法
 */

/**
2.下面哪一行代码会 panic，请说明原因？
package main

func main() {
  var m map[int]bool // nil
  _ = m[123]
  var p *[5]string // nil
  for range p {
    _ = len(p)
  }
  var s []int // nil
  _ = s[:]
  s, s[0] = []int{1, 2}, 9
}

参考答案及解析：第 12 行。因为左侧的 s[0] 中的 s 为 nil。

引自：《Go语言101》
 */
