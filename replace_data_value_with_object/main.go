package main

import "fmt"

/*
手法: Replace Data Value With Object - 使用对象替换值

目的: 基于后续可扩展的原则

细节: 在Go里面其实个人觉得有两种设计思路(有不同的应用场景)
1. 使用的新的struct
2. 别名类型
*/

func main()  {
	o := new(Order)
	o.name = "xiaoming"

	fmt.Println(o.getName())
}

type Order struct {
	name string
}

func (o *Order) getName() string {
	return o.name
}
