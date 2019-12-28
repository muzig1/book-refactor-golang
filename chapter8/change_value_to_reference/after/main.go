package main

import "fmt"

func main() {
	xiaoming := &Customer{name: "xiaoming"}
	o1 := Order{cus: xiaoming} // 订单一
	o2 := Order{cus: xiaoming} // 订单二

	// 引用同一个对象,其中一个修改,其他也会修改
	o1.cus.name = "xiaohua"
	fmt.Println(o1.getName())
	fmt.Println(o2.getName())
}

type (
	// Order .
	Order struct {
		cus *Customer // 此处的Customer为引用类型
	}

	// Customer .
	Customer struct {
		name string
	}
)

func (o Order) getName() string {
	return o.cus.name
}
