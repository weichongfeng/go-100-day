package main

/**
1.下面代码有什么问题吗？
func main()  {

    for i:=0;i<10 ;i++  {
    loop:
        println(i)
    }
    goto loop
}
参考答案及解析：goto 不能跳转到其他函数或者内层代码。编译报错：

    goto loop jumps into block starting at
 */

/**
2.下面代码输出什么，请说明。
func main() {
    x := []int{0, 1, 2}
    y := [3]*int{}
    for i, v := range x {
        defer func() {
            print(v) 2 2 2
        }()
        y[i] = &v 2 2 2
    }
    print(*y[0], *y[1], *y[2])
}
 */

func main() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2])
}