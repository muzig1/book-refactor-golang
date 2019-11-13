package main

import "fmt"

/*
手法: self encapsulate field - 封装字段

目的: 分为两种流派

1. 直接访问字段, 不需要封装
2. 间接访问字段, 需要进行封装

优点:
	前者更方便使用,而且减少代码量
	后者更够更加灵活的控制字段的数据, 父类还可以覆盖子类实现的获取函数

缺点:
	两者的优点也是相互的缺点

细节:
*/

// 栗子说明: 间接获取,封装的好处
// 具体栗子: 获取猪肉的价格

func main()  {
	s := new(Shop)
	s.price = 1000

	// 从商店获取猪肉价格
	fmt.Println(s.getPrice())
	// 获取市场内部价
	fmt.Println(s.Pork.getPrice())
}

type (
	Shop struct {
		Pork
	}

	Pork struct {
		price int
	}
)

func (s *Shop) getPrice() int {
	// 由于店铺猪肉涨价了, 所以可以直接在Shop的方法里面去控制价格
	return s.Pork.getPrice()*2
}

func (p *Pork) getPrice() int {
	return p.price
}

