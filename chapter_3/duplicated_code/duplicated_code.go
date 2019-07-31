package main

/*
问题：Duplicated Code - 重复代码

解决：
1. Extract Method - 提炼函数
2. Pull Up Method - 提升函数权限
3. Extract Class  - 提炼类，将提炼的函数放在最合适的位置，其他需要使用的就用此调用
*/

func main() {
	// 购买香烟
	OnShoppingReq(&Adult{
		Name:   "xiaoming",
		Age:    18,
		Cash:   100,
		Pocket: make([]*Cigarette, 0),
	}, Panda)
}

// ------ Bad Code >>>>>>
// Module:
// 概述：一个小卖部，销售香烟，只能销售给18岁以上的成年人
// 对象：小卖部(Shop), 香烟(Cigarette), 成年人(Adult)
type CigaretteType uint8 // 香烟类型

const (
	Creature CigaretteType = iota // 娇子
	Panda                         // 熊猫
)

type (
	// Shop 小卖部
	Shop struct {
		CigSet []*Cigarette // 香烟集
	}

	// Cigarette 香烟
	Cigarette struct {
		Type  CigaretteType
		Price int // 价格
	}

	// Adult 成年人
	Adult struct {
		Name string // 姓名
		Age  uint8  // 年龄
		Cash int    // 现金

		Pocket []*Cigarette // 装香烟的口袋
	}
)

// Shopping 购买香烟
func (s *Shop) Shopping(adult *Adult, cigaretteType CigaretteType) (sell *Cigarette) {
	if !(adult.Age >= 18) {
		return
	}
	for _, cig := range s.CigSet {
		if cig.Type == cigaretteType {
			sell = cig
		}
	}
	if sell == nil {
		return
	}
	if adult.Cash < sell.Price {
		return nil
	}
	return
}

// Controller:
var ShopInstance = &Shop{[]*Cigarette{
	{
		Type:  Panda,
		Price: 10,
	},
	{
		Type:  Creature,
		Price: 20,
	},
}}

// OnShoppingReq 购物请求
func OnShoppingReq(adult *Adult, cigaretteType CigaretteType) {
	// 检查
	if !(adult.Age >= 18) {
		return
	}
	var cigarette *Cigarette
	for _, cig := range ShopInstance.CigSet {
		if cig.Type == cigaretteType {
			cigarette = cig
		}
	}
	if cigarette == nil {
		return
	}
	if adult.Cash < cigarette.Price {
		return
	}

	// 执行
	adult.Pocket = append(adult.Pocket, ShopInstance.Shopping(adult, cigaretteType))
}

// ------ Bad Code >>>>>>

// ------ Extract Method >>>>>>

// ------ Extract Method <<<<<<

// ------ Pull Up Method >>>>>>

// ------ Pull Up Method <<<<<<

// ------ Extract Class >>>>>>

// ------ Extract Class <<<<<<
