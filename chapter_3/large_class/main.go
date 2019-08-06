package main

/*
问题：Large Class - 过大的类

解决：
	1. Extract Method
	2. Extract Class | Extract Subclass

tips:
	1. 尽量抽象相似属性，并使用组合的思维
	2. 避免一个结构内，拥有很多字段，这样会让该结构的成员函数变向增多
*/

func main() {

}

// 举栗子：
// 		针对于游戏中士兵，分为步兵(foot soldier)，骑兵(ride soldier)，弓兵(bow soldier)

// ------ Not Good Code --->>>

type (
	//FootSoldier 步兵
	FootSoldier struct {
		ID     int64  // 唯一ID
		CfgID  int32  // 配置ID
		Level  int32  // 等级
		Name   string // 姓名
		Atk    int32  // 攻击力
		Def    int32  // 防御力
		Action int32  // 行动力
		// ...此处简略，暂时列举一部分，实际需求比这个复杂的多
	}

	//RideSoldier 骑兵
	RideSoldier struct {
		ID     int64
		CfgID  int32
		Level  int32
		Name   string
		Atk    int32
		Def    int32
		Action int32
	}

	//BowSoldier 弓兵
	BowSoldier struct {
		ID     int64
		CfgID  int32
		Level  int32
		Name   string
		Atk    int32
		Def    int32
		Action int32
	}
)

// ------ Not Good Code ---<<<

// ------ Extract Class | Extract Subclass --->>>

// Soldier 士兵基类 - 此处嵌套结构体，一般较少使用，常分离设计，若只会使用一次，则建议内嵌定义
type Soldier struct {
	ID    int64
	CfgID int32

	Basic struct {
		Name  string
		Level int32
	}

	Prop struct {
		Atk, Def, Action int32
	}
}

// 此处使用的结构体，非指针对象，这个后续再做细节描述 - 什么时候用指针对象? 什么使用结构对象?
type (
	//FootSoldierDef 步兵
	FootSoldierDef struct {
		Soldier
	}

	//RideSoldier 骑兵
	RideSoldierDef struct {
		Soldier
	}

	//BowSoldier 弓兵
	BowSoldierDef struct {
		Soldier
	}
)

// ------ Extract Class | Extract Subclass ---<<<
