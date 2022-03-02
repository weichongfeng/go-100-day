package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "--->", *v)
	}
}
/**输出
0 ---> 3
1 ---> 3
2 ---> 3
3 ---> 3
 */
