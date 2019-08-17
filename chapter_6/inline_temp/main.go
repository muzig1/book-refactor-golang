package main

import "fmt"

/*
手法：Inline Temp - 内联变量

步骤:
	1. 确认该函数内部，临时变量的使用地方只有一处
	2. 确认该函数的引用点，然后一一的替换为内联变量

注意：
	1. 若需要debug；内联变量无法查看具体的数据;此时可以选择使用临时变量
*/

func main() {
	// 优化前：
	base := getPctBase()
	result := 100 / base
	fmt.Println("result:%v", result)

	// 优化后：
	fmt.Println("result:%v", 100/getPctBase())
}

func getPctBase() int32 {
	return 10000
}
