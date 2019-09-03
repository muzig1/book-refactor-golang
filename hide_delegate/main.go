package main

import "fmt"

/*
手法：Hide Delegate - 隐藏代理

目的：
1. 代理发生变化的时候，不会影响客户
*/

func main() {
	p := NewPerson("xiaoming", "xiaohua", "golang")
	managerName := p.GetManagerDelegate()
	fmt.Println("manager name:", managerName)
}

func NewPerson(name, mName, departName string) *Person {
	depart := Department{
		Name:    departName,
		Manager: mName,
	}
	return &Person{
		Name:               name,
		Depart:             depart,
		GetManagerDelegate: depart.GetManager,
	}
}

type Person struct {
	Name               string
	Depart             Department
	GetManagerDelegate func() string
}

func (p *Person) GetManager() string {
	return p.GetManagerDelegate()
}

type Department struct {
	Name    string
	Manager string
}

func (d *Department) GetManager() string {
	return d.Manager
}
