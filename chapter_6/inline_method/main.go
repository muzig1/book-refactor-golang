package main

import "fmt"

/*
手法：Inline Method - 内联函数

目的：用最简单的例子描述重构的动机，但实际情况并非如此，此重构手法需谨慎使用，具体查看注意事项

注意：
1. 检查函数不具备多态性
	具有多态性的函数，如果内联函数的话，就不具有重写的功能了；
	在Go语言里面，指代继承接口的类型，其函数不能使用此重构手法
	现在基本使用IDE的情况下，一眼就能看出是否多态或继承接口

步骤：
1. 检查函数的所有引用点
2. 可以使用全局匹配函数，并进行替换操作
3. 编译，测试
4. 删除函数定义

*/

func main() {
	// 优化前：
	c := &Clothes{Size: 179}
	fmt.Println(c.GetMeasure())

	// 优化后：
	cc := &NewClothes{Size: 181}
	fmt.Println(cc.GetMeasure())
}

type Clothes struct {
	Size int32
}

//GetMeasure 获取衣服尺寸
func (c *Clothes) GetMeasure() int32 {
	if c.moreThan180() {
		return 1
	}
	return 0
}

func (c *Clothes) moreThan180() bool {
	return c.Size > 180
}

type NewClothes struct {
	Size int32
}

//GetMeasure 获取衣服尺寸
func (c *NewClothes) GetMeasure() int32 {
	if c.Size > 180 {
		return 1
	}
	return 0
}
