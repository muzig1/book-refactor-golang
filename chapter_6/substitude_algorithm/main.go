package main

import (
	"strings"
)

/*
手法：Substitude Algorithm - 替换算法

tips:
	1. 此手法从写法上面去优化算法，并不是从计算本身去优化算法，这点区别需要注意
*/

func main() {
	// 优化前：
	_ = func(people []string) string {
		for _, p := range people {
			if strings.EqualFold(p, "Don") {
				return "Don"
			}
			if strings.EqualFold(p, "John") {
				return "John"
			}
			if strings.EqualFold(p, "Kent") {
				return "Kent"
			}
		}
		return ""
	}

	// 优化后：
	_ = func(people []string) string {
		candidates := Candidate{"Don", "John", "Kent"}
		for _, p := range people {
			if candidates.Containers(p) {
				return p
			}
		}
		return ""
	}
}

type Candidate []string

func (c Candidate) Containers(name string) bool {
	for _, n := range c {
		if n == name {
			return true
		}
	}
	return false
}
