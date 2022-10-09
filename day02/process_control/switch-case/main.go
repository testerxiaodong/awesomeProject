package main

import "fmt"

func main() {
	n := 6
	// fallthrough可以穿过一层
	switch n {
	case 1:
		fmt.Println("第一条分支")
		fallthrough
	case 2:
		fmt.Println("第二条分支")
	case 3:
		fmt.Println("第三条分支")
	case 4:
		fmt.Println("第四条分支")
	case 5:
		fmt.Println("第五条分支")
	default:
		fmt.Println("默认分支")
	}
	// 初始化可以写在switch里
	// case里可以写多个匹配项，逗号分隔
	// case里可以写逻辑语句
	switch m := 1; m {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}
}
