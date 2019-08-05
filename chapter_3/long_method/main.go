package main

import "fmt"

/*
问题: Long Method - 超长函数

解决：
	1. Extract Method - 提炼函数
	2. Replace Temp with Query
	3. Introduce Parameter Object | Preserve Whole Object
	4. Replace Method with Method Object
	5. Decompose Conditional

tips:
	1. "间接层的好处"
		解释能力、共享能力、选择能力
	2. 提炼信号
		a. 条件表达式以及循环一般是提炼信号
		b. 需要注释的地方，一般需要提炼函数
*/

func main() {
	var (
		p = &Player{
			ID: 960907,
		}
		p2 = &Player{
			ID: 960506,
		}
		item = &UseItemReq{
			item: Tuple{
				Typ: ICAddExp,
				ID:  0,
				Val: 100,
			},
		}
	)

	// 使用道具请求发送
	// 优化前：
	_ = OnUseItemReq(p, item)
	fmt.Println(p.ToString())
}

// ItemCategory 道具类型
type ItemCategory string

const (
	_        ItemCategory = ""
	ICAddExp              = "add_exp" // 添加经验
	ICAddRss              = "add_rss" // 添加资源
)

// UseItemReq 使用道具请求
type UseItemReq struct {
	item Tuple
}

// Tuple 三元组
type Tuple struct {
	Typ string
	ID  int32
	Val int32
}

// Player 玩家
type Player struct {
	ID  int64
	Exp int32 // 玩家经验
	Rss int32 // 玩家资源
}

// ToString 玩家信息
func (p *Player) ToString() string {
	return fmt.Sprintf("player ID:%d Exp:%d, Rss:%d", p.ID, p.Exp, p.Rss)
}

// 举🌰(栗子)：游戏逻辑中，常出现的使用道具的效果实现
// ------ Not Good Code --->>>

// UseItem 使用道具
func OnUseItemReq(p *Player, req *UseItemReq) (err error) {
	// 此处模拟的是最简单的情况，搭配配置获取等，会让其检查变得臃肿
	switch req.item.Typ {
	case ICAddExp:
		p.Exp += req.item.Val
	case ICAddRss:
		p.Rss += req.item.Val
	default:
		err = fmt.Errorf("error: OnUseItemReq item typ is not match, playerID:%d itemTyp:%d", p.ID, req.item.Typ)
	}
	return
}

// ------ Not Good Code ---<<<
