package main

import "fmt"

func main() {
	cus := NewCustomer("xiaoming", "9527")
	o1 := NewOrder(cus)
	o2 := NewOrder(cus)
	o3 := NewOrder(NewCustomer("xiaohua", "1818"))
	orders := []*Order{o1, o2, o3}

	cnt := NumberofOrders(orders, "028-9527")
	fmt.Println("number is ", cnt)
}

// NewOrder 构造函数
func NewOrder(cus *Customer) *Order {
	return &Order{customer: cus}
}

// NewCustomer 构造函数
func NewCustomer(name, num string) *Customer {
	return &Customer{
		name:   name,
		number: num,
	}
}

type (
	// Order 订单
	Order struct {
		// 此处使用引用, 意味着同一用户可以再不同的订单结构下面
		// 之前为值类型, 就不能做到这一点
		customer *Customer
	}

	// Customer 用户
	Customer struct {
		name   string
		number string
	}
)

func (c Customer) getNumber() string {
	return "028-" + c.number
}

// NumberofOrders 订单中指定内容的数量
func NumberofOrders(orders []*Order, number string) int {
	var sum int
	for _, order := range orders {
		if order.customer.getNumber() == number {
			sum++
		}
	}
	return sum
}
