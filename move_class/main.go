package main

/*
手法：Move Class - 移动类

目的：
1. 细化职责

tips:
1. 拆分类更多，支持并发锁的对象越多 - (暂时不知道好处与坏处)
*/

func main() {

}

type Person struct {
	Name string
	PersonOffice
}

type PersonOffice struct {
	OfficeAreaCode, OfficeNumber string
}
