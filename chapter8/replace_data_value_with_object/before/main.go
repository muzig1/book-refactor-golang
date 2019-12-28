package main

import "fmt"

func main() {
	orders := []*Order{
		NewOrder("xiaoming"),
		NewOrder("xiaohua"),
		NewOrder("xiaoxue"),
		NewOrder("xiaoming"),
	}

	cnt := NumberofOrders(orders, "xiaoming")
	fmt.Println("xiaoming order number is ", cnt)
}

// NewOrder 构造函数
func NewOrder(name string) *Order {
	return &Order{
		customer: name,
	}
}

// Order 订单
type Order struct {
	customer string
}

// NumberofOrders 订单中指定内容的数量
func NumberofOrders(orders []*Order, name string) int {
	var sum int
	for _, order := range orders {
		if order.customer == name {
			sum++
		}
	}
	return sum
}
