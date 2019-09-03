package main

/*
手法：Move Class - 移动类

目的：
1. 细化职责
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
