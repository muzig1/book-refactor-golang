# Self Encapsulate Field - 自我封装

在自我封装这个概念下, 有两种不同的流派, 一种是直接访问字段; 一种是间接访问字段. 两种方式各有好处.

**间接访问字段:**

1. 复写 - 子类能够覆盖父类的方法(golang也通过间接访问方式,达到复写的效果)
2. 灵活 - 对数据管理方式更友好,弹性,可自定义; 如延迟初始化之类的

**直接访问字段:**

1. 简单 - 代码容易阅读
2. 轻量 - 取值函数代码相对少一些

书中给的建议是: 配合团队协作为主, 先采用简单的直接访问方式, 当遇到了相对复杂的需求的时候, 再进行重构修改.

## Demo

> [直接访问方式](demo/../straight/main.go)

```go
// IntRange .
type IntRange struct {
	_low, _high int
}

func (ir IntRange) includes(arg int) bool {
	return arg >= ir._low && arg <= ir._high
}

func (ir *IntRange) grow(factor int) {
	ir._high *= factor
}
```

> [间接访问方式](demo/../indirect/main.go)

```go
type (
	// IntRange 父类
	IntRange struct {
		_low, _high int
	}

	// CappedRange 子类
	CappedRange struct {
		*IntRange // 组合
		_cap      int
	}
)

// >>> 父类

func (ir IntRange) getLow() int {
	return ir._low
}

func (ir IntRange) getHigh() int {
	return ir._high
}

func (ir *IntRange) setLow(arg int) {
	if ir.getHigh() < arg {
		return
	}
	ir._low = arg
}

func (ir *IntRange) setHigh(arg int) {
	if ir.getLow() > arg {
		return
	}
	ir._high = arg
}

func (ir IntRange) includes(arg int) bool {
	return arg >= ir._low && arg <= ir._high
}

func (ir *IntRange) grow(factor int) {
	ir._high *= factor
}

// <<<

// >>> 子类

func (cr CappedRange) getCap() int {
	return cr._cap
}

func (cr *CappedRange) setCap(arg int) {
	cr._cap = arg
}

func (cr CappedRange) getHigh() int {
    // 此处进行复写父类的方法, 达到约束的效果
    return int(math.Min(float64(cr.IntRange.getHigh()), float64(cr.getCap())))
}

// <<<

```