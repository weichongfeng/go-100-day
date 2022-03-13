package main

/**
1.下面这段代码输出什么？
const (
    a = iota 1
    b = iota 2
)
const (
    name = "name"
    c    = iota 1
    d    = iota 2
)
func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
 */

/**
2.下面这段代码输出什么？为什么？
type People interface {
    Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func main() {

    var s *Student
    if s == nil {
        fmt.Println("s is nil")
    } else {
        fmt.Println("s is not nil")
    }
    var p People = s
    if p == nil {
        fmt.Println("p is nil")
    } else {
        fmt.Println("p is not nil")
    }
}

参考答案及解析：s is nil 和 p is not nil。这道题会不会有点诧异，我们分配给变量 p 的值明明是 nil，然而 p 却不是 nil。记住一点，
当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。上面的代码，给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，
是一个 nil 指针，所以相等条件不成立
 */

