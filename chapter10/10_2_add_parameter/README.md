# Add Parameter - 添加参数

添加参数是一种常用的重构手法. 书中描述的比较细节, 此处按照个人最近重构的理解, 做一个重构指导.

## 步骤

1. 整理函数内, 变量的作用域
    - 函数内部变量
    - 传入参数(外部变量)
    - 对象参数(即类实现的方法)
2. 分析这个函数的复用性, 是否可以单独实现, 不作为对象方法.
    - 体现在, 尽量可以把对象的数据, 按照参数的方式传入
3. 实际操作细节
    - 作用域小 - 尽量降低函数的作用域, 能访问的变量尽可能的少, 这样简单, 可读性高
    - 复用性强 - 体现在不需要隶属于函数对象, 可以理解为全局函数(但类函数和这个还是要区分开)
    
## 栗子(🌰)

> [重构前](before/main.go)

```go
// RefreshAllComp 刷新所有电脑价格
func (cc *CompCity) RefreshAllComp() {
	// 此处将所有逻辑都放在最顶层CompCity对象下, 无法复用对单个Comp进行刷新操作
	for _, comp := range cc.Comps {
		switch comp.Typ {
		case Simple:
			comp.Price *= cc.Volatility * 2
		case HighEnd:
			comp.Price *= cc.Volatility * 3
		}
	}
}
```

> [重构后](after/main.go)

```go
// RefreshAllComp 刷新所有电脑价格
func (cc *CompCity) RefreshAllComp() {
	// 此处将函数作用域分散, 层层调用, 外部就可以获取单个对象, 进行刷新
	cc.Comps.refresh(cc.Volatility)
}
func (set CompSet) refresh(volatility int32) {
	for _, comp := range set {
		comp.refresh(volatility)
	}
}

func (c *Comp) refresh(volatility int32) {
	switch c.Typ {
	case Simple:
		c.Price *= volatility * 2
	case HighEnd:
		c.Price *= volatility * 3
	}
}
```