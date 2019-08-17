package main

import (
	"math/rand"
	"strings"
)

/*
手法：Introduce Explain Variable - 引入解释性变量

使用场景：
	1. Extract Method需要大量时间的情况下，只想简化代码理解复杂度的时候
*/

type String string

func (s *String) ToUpper() String {
	ret := strings.ToUpper(s.ToString())
	return String(ret)
}

func (s String) IndexOf(aim String) int {
	return strings.Index(s.ToString(), aim.ToString())
}

func (s String) ToString() string {
	return string(s)
}

func main() {
	var (
		platform String = "Mac os 1.0"
		browser  String = "IE browser 1.0"
		resize          = rand.Int31n(10)
	)

	// 优化前：
	if platform.ToUpper().IndexOf("MAC") != -1 &&
		browser.ToUpper().IndexOf("IE") != -1 &&
		resize > 0 {
		// do something
	}

	// 优化后
	var (
		isMacOS     = platform.ToUpper().IndexOf("MAC") != -1
		isIEBrowser = browser.ToUpper().IndexOf("IE") != -1
		isResized   = resize > 0
	)

	if isMacOS && isIEBrowser && isResized {
		// do something
	}
}
