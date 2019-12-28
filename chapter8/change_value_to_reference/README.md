# Change Value to Reference - 值类型替换为引用类型

## 动机

在许多系统中, 按照数据的存储位置(堆栈)分类: 值类型&引用类型; 有时候希望多个地方所引用的是同一个对象, 即修改了一处, 其他地方也跟着修改了

## demo

> [重构前](before/main.go)

```go
type (
	// Order .
	Order struct {
		cus Customer // 此处的Customer为值类型
	}

	// Customer .
	Customer struct {
		name string
	}
)

func (o Order) getName() string {
	return o.cus.name
}
```

> [重构后](after/main.go)

```go
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
```