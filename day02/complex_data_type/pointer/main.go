package main

import "fmt"

func main() {
	// &: 取地址
	n := 18
	fmt.Println(&n)
	// 根据地址取值
	p := &n
	m := *p
	fmt.Println(m)
	// make函数与new区别
	// make函数一般为map、chan、slice申请内存空间，返回值就是对应类型
	// new函数返回对应类型的指针，一般用于基本数据类型
}
