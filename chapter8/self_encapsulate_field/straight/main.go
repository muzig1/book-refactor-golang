package main

import "fmt"

func main() {
	r := &IntRange{_low: 10, _high: 20}
	fmt.Println(r.includes(15))

	r.grow(2)
	fmt.Println(r.includes(15))
}

// IntRange .
type IntRange struct {
	_low, _high int
}

func (ir IntRange) includes(arg int) bool {
	return arg >= ir._low && arg <= ir._high
}

func (ir *IntRange) grow(factor int) {
	ir._high *= factor
}
