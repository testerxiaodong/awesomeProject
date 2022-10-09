package main

import "fmt"

func main() {
	score := 99
	if score > 90 {
		fmt.Println("优秀的学霸")
	} else if score >= 80 {
		fmt.Println("好孩子")
	} else if score >= 70 {
		fmt.Println("加把油")
	} else if score >= 60 {
		fmt.Println("老师尽力了")
	} else {
		fmt.Println("重修大礼包")
	}
}
