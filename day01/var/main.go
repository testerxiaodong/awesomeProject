package main

import "fmt"

func main() {
	// 声明变量
	var (
		name string
		age  int
		isOk bool
	)
	// 变量必须声明后才能使用
	// 变量声明后没有初始化，变量的值是对应类型的零值
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)
	// 局部变量变量声明后必须使用，否则不能通过编译
	// 声明变量同时赋值，类型推导
	var s1 = "haha"
	fmt.Println(s1)
	// 短变量声明
	s2 := "沙河小王子"
	fmt.Println(s2)
	// 匿名变量：不占用命名空间，不会分配内存
	// 同一个作用域中，同一个变量不能重复声明
}
