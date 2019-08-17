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
}

type Item struct {
	Quality int32
	Price   int32
}

func (i *Item) getPrice() int32 {
	return i.Quality * i.Price
}
