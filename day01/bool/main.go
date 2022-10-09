package main

import "fmt"

func main() {
	// 布尔类型
	// 不允许整型转换为布尔型
	// 布尔型不能参与运算
	b1 := false
	fmt.Printf("%T: %v", b1, b1)
}
