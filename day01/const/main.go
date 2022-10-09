package main

import "fmt"

// 常量声明使用关键字const
const (
	pi  = 3.1415926
	num = 114514
)

// 批量声明常量，某一行声明后没有赋值，默认和上一行值相等
const (
	n1 = 100
	n2
	n3
)

// iota：iota在关键字const出现的时候被重置为0
// const中每新增一行常量声明将使iota的值+1
// 常使用_跳过一行声明
const (
	a0 = iota
	a1
	a2
	_
	a4
)

func main() {
	// 定义了常量之后不能修改
	fmt.Println(pi)
	fmt.Println(num)
	fmt.Println(n1, n2, n3)
	fmt.Println(a0, a1, a2, a4)
}
