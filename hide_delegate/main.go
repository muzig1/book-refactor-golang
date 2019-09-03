package main

/*
手法：Hide Delegate - 隐藏代理

目的：
1. 代理发生变化的时候，不会影响客户
*/

func main() {

}

type Person struct {
	Name   string
	Depart Department
}

func (p *Person) GetManager() string {
	return p.Depart.Manager
}

type Department struct {
	Name    string
	Manager string
}
