package main

import "fmt"

/*
手法： Split Temporary Variable - 拆分临时变量

tips：
	1. 此处文中提到使用final 关键字，go中没有此相似的关键字，但可以注意一下
	2. rrr
*/

func main() {
	var (
		height = 10
		width  = 5
	)

	// 优化前：
	tmp := height * width
	fmt.Println(tmp)
	tmp = 2 * (height + width)
	fmt.Println(tmp)

	// 优化后:
	area := height * width
	fmt.Println(area)
	perimeter := 2 * (height + width)
	fmt.Println(perimeter)
}
