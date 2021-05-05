package main

import "fmt"

func main()  {
	const foo = `
改名不仅仅是修改名字而已。
如果你想不出一个好名字，说明背后很可能潜 藏着更深的设计问题。
为一个恼人的名字所付出的纠结，常常能推动我们对代码 进行精简。
`
	fmt.Print(foo)
}
