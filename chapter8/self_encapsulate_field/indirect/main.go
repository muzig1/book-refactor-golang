package main

import (
	"fmt"
	"math"
)

func main() {
	r := &IntRange{}
	r.setLow(10)
	r.setHigh(20)

	subR := &CappedRange{}
	subR.IntRange = r
	subR.setCap(15)

	fmt.Println(subR.getHigh())
}

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
	return int(math.Min(float64(cr.IntRange.getHigkh()), float64(cr.getCap())))
}

// <<<
