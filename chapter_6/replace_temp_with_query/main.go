package main

import "fmt"

/*
手法：Replace Temp With Query - 封装临时变量

担心:
	1. 担心性能问题，因为会调用多次计算；重构暂时不考虑性能问题，后续会有更多时候来思考性能优化的问题
*/

func main() {
	item := &Item{
		Quality: 3,
		Price:   10,
	}
	// 优化前：
	price := func() int32 {
		basePrice := item.Quality * item.Price
		if basePrice > 15 {
			return basePrice * 2
		} else {
			return basePrice * 3
		}
	}()
	fmt.Println("price:", price)

	// 优化后：
	price = func() int32 {
		if item.getPrice() > 15 {
			return item.getPrice() * 2
		} else {
			return item.getPrice() * 3
		}
	}()
	fmt.Println("prince:", price)

	// ------案例二

	// 优化前：
	price = func() int32 {
		basePrice := item.Quality * item.Price
		var discountFactor int32
		if basePrice > 15 {
			discountFactor = 2
		} else {
			discountFactor = 3
		}
		return basePrice * discountFactor
	}()

	// 优化后：
	// 步骤一：
	price = func() int32 {
		var discountFactor int32
		if item.getPrice() > 15 {
			discountFactor = 2
		} else {
			discountFactor = 3
		}
		return item.getPrice() * discountFactor
	}()

	// 步骤二：
	price = func() int32 {
		return item.getPrice() * item.getDiscountFactor()
	}()

}

type Item struct {
	Quality int32
	Price   int32
}

func (i *Item) getPrice() int32 {
	return i.Quality * i.Price
}

func (i *Item) getDiscountFactor() int32 {
	if i.getPrice() > 15 {
		return 2
	} else {
		return 3
	}
}
