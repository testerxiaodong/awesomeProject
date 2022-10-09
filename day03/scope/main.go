package main

import "fmt"

var (
	globalVar string
)

func f1() {
	localVar := "我是局部作用域"
	fmt.Println(localVar)
	fmt.Println(globalVar)
}

func main() {
	// go语言有三种作用域
	// 全局作用域
	// 函数局部作用域
	// 语句块作用域
	// 变量查找：当前作用域找，找不到就向外部作用域找，直至全局作用域，找到为止
	fmt.Println(globalVar)
	f1()
}
