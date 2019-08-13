package chapter_6

import "fmt"

/*
问题：Extract Method - 提炼函数

解决：
1.
*/

func main() {

}

// 情景一：无变量的情况 --->>>
func printOwing() {
	var (
		name        string
		outstanding int32
	)

	// print banner
	fmt.Println("************")
	fmt.Println("***** Customer Owes *****")
	fmt.Println("************")

	// calculate outstanding
	// 略...

	// print details
	fmt.Println("name: ", name)
	fmt.Println("outstanding: ", outstanding)
}

// 优化后：
func _printOwing() {
	var (
		name        string
		outstanding int32
	)

	// print banner
	printBanner()

	// calculate outstanding
	// 略...

	// print details
	fmt.Println("name: ", name)
	fmt.Println("outstanding: ", outstanding)
}

func printBanner() {
	fmt.Println("************")
	fmt.Println("***** Customer Owes *****")
	fmt.Println("************")
}

// 情景一：无变量的情况 ---<<<

// 情景二：有变量的情况 --->>>
func printDemo() {
	var sum int

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("start")
	fmt.Println("sum: ", sum)
	fmt.Println("end")
}

// 优化后：
func _printDemo() {
	var sum int

	for i := 0; i < 10; i++ {
		sum += i
	}

	printDetail(sum)
}

func printDetail(sum int) {
	fmt.Println("start")
	fmt.Println("sum: ", sum)
	fmt.Println("end")
}

// 情景二：有变量的情况 ---<<<

// 情景三：对局部变量再赋值的情况 --->>>

// 优化后：
// 情景三：对局部变量再赋值的情况 ---<<<
