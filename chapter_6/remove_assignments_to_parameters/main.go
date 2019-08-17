package main

/*
手法：Remove Assignments To Parameters
*/

func main() {
	// 优化前:
	func(val, rate int32) int32 {
		if val > 10 {
			// 此处对参数进行二次使用 - 拒绝此用法
			val = 1
		}
		return val
	}(10, 20)

	// 优化后
	func(val, rate int32) int32 {
		result := val
		if val > 10 {
			result = 1
		}
		return result
	}(10, 20)
}
