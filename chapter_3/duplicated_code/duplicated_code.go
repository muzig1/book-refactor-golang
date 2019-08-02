package main

/*
问题：Duplicated Code - 重复代码

解决：
1. Extract Method - 提炼函数
2. Pull Up Method - 提升函数权限
3. Extract Class  - 提炼类，将提炼的函数放在最合适的位置，其他需要使用的就用此调用
*/

func main() {
}

// ------ Not Good Code >>>>>>
type fruitSet []*fruit

type fruit struct {
	SN   int32
	Name string
}

// GetFruit 根据水果名获取水果
func (set fruitSet) GetFruit(name string) *fruit {
	for _, f := range set {
		if f.Name == name {
			return f
		}
	}
	return nil
}

// GetFruit2 根据水果sn 码获取水果
func (set fruitSet) GetFruit2(sn int32) *fruit {
	for _, f := range set {
		if f.SN == sn {
			return f
		}
	}
	return nil
}

// ------ Not Good Code <<<<<<

// ------ Extract Method >>>>>>
// 优化目标：提炼函数，提炼for 循环

// getFruit 将函数穿入，类似于函数式编程，抽出共同的for 循环
func (set fruitSet) getFruit(selector func(f *fruit) (match bool)) *fruit {
	for _, f := range set {
		if selector(f) {
			return f
		}
	}
	return nil
}

// GetFruitByName 可以动态定制函数传入， 不用再编写for 循环
func (set fruitSet) GetFruitByName(name string) *fruit {
	return set.getFruit(func(f *fruit) (match bool) {
		if f.Name == name {
			match = true
		}
		return
	})
}

// ------ Extract Method <<<<<<

// ------------------------------------------------------------------------

// ------ Not Good Code >>>>>>
// Module --->>>
// medal 奖章
type medal struct {
	CfgID int32
	Name  string
	Level int32
}

func (m *medal) LevelUp() {
	// 获取配置表信息
	cfg := GetMedalCfg(m.CfgID)
	if cfg == nil {
		return
	}
	// 约束检查
	if m.Level < cfg.MaxLvl {
		m.Level++
	}
}

func (m *medal) LevelDown() {
	// 获取配置表信息
	cfg := GetMedalCfg(m.CfgID)
	if cfg == nil {
		return
	}
	// 约束检查
	if m.Level != 0 {
		m.Level--
	}
}

var MedalCfgs = []*MedalCfg{
	{
		CfgID:  1,
		MaxLvl: 10,
	},
}

type MedalCfg struct {
	CfgID  int32 `bson:"cfg_id"`  // 配置ID
	MaxLvl int32 `bson:"max_lvl"` // 最大等级 - 用于检查
}

func GetMedalCfg(cfgID int32) *MedalCfg {
	for _, cfg := range MedalCfgs {
		if cfg.CfgID == cfgID {
			return cfg
		}
	}
	return nil
}

// Controller --->>>
func DoMedalLevelUp() {
	m := &medal{
		CfgID: 1,
		Name:  "medal",
		Level: 5,
	}
	// 此处一下调用，其实会调用同样的代码，也就是**配置获取** - 我们需要优化一下
	m.LevelUp()
	m.LevelDown()
}

// ------ Not Good Code <<<<<<

// ------ Pull Up Method >>>>>>
func (m *medal) newLevelUp(cfg *MedalCfg) {
	// 约束检查
	if m.Level < cfg.MaxLvl {
		m.Level++
	}
}

func (m *medal) newLevelDown(cfg *MedalCfg) {
	// 约束检查
	if m.Level != 0 {
		m.Level--
	}
}

// Controller --->>>
func newDoMedalLevelUp() {
	m := &medal{
		CfgID: 1,
		Name:  "medal",
		Level: 5,
	}

	// 优化后，采用传参的方式， 提升了函数的权限
	cfg := GetMedalCfg(m.CfgID)
	if cfg == nil {
		return
	}
	m.newLevelUp(cfg)
	m.newLevelDown(cfg)
}

// ------ Pull Up Method <<<<<<

// ------------------------------------------------------------------------

// ------ Not Good Code >>>>>>

// ------ Not Good Code <<<<<<

// ------ Extract Class >>>>>>

// ------ Extract Class <<<<<<
