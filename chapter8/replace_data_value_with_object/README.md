
# Replace Data Value with Object - 以对象替代数据值

## 动机

开发初期可能已最简单的数据形式去开发,但是随着开发进行,需求越来越多,越来越复杂,可能原有数据结构不再适合了. 比如, 开始可能用一个字符串表示电话号码,后来增加了"格式化", "抽取区号"之类的特殊行为,一个简单的string类型就无法满足,就只有侵入上层的函数成员; Duplicate Code 以及Feature Envy 的坏坏味道就从代码中散发出来. 当出现这种情况的时候就应该对代码进行重构

## Demo

> 简单案例描述

```go
type Order struct {
    Customer string
}

--- 重构前后 对比线

type (
    Order {
        customer Customer
    }

    Customer {
        name string
    }
)
```

> [重构前](before/main.go)

```go
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
```

> [重构后](after/main.go)

```go
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
```