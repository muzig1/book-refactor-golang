package main

/*
手法：Replace Method With Method Object - 以函数对象替代函数
*/

func main() {
	delta := func() int {
		return 10
	}

	// 大概理解文中想表明的意思，但是感觉这个例子不是特别容易理解
	// 后续先保留文中的例子，然后后面是自己对此手法的理解以及例子

	// 书中例子
	// 优化前：
	_ = func(inputVal, quantity, yearToData int) int {
		importantValue1 := (inputVal * quantity) + delta()
		importantValue2 := (inputVal * yearToData) + 100
		if (yearToData - importantValue1) > 100 {
			importantValue2 -= 20
		}
		importantValue3 := importantValue2 * 7
		return importantValue3 - 2*importantValue1
	}

	// 优化后：
	_ = gamma(delta)
}

func gamma(delta func() int) int {
	return NewGammaDef(delta, 10, 20, 30).Compute()
}

func NewGammaDef(src func() int, inputVal, quantity, yearToData int) *GammaDef {
	return &GammaDef{
		delta:      src,
		inputVal:   inputVal,
		quantity:   quantity,
		yearToDate: yearToData,
	}
}

type GammaDef struct {
	delta                          func() int
	inputVal, quantity, yearToDate int
	importantValue1                int
	importantValue2                int
	importantValue3                int
}

func (g *GammaDef) Compute() int {
	g.importantValue1 = (g.inputVal * g.quantity) + g.delta()
	g.importantValue2 = (g.inputVal * g.yearToDate) + 100
	g.importantThing()
	g.importantValue3 = g.importantValue2 * 7
	return g.importantValue3 - 2*g.importantValue1
}

func (g *GammaDef) importantThing() {
	if (g.yearToDate - g.importantValue1) > 100 {
		g.importantValue2 -= 20
	}
}
