package main

import "fmt"

func add(x, y int) int {
	sum := x + y
	return sum
}

func f1(f func(x, y int) int, x, y int) int {
	res := f(x, y)
	return res
}

func main() {
	// 函数是一种类型
	fmt.Printf("%T\n", add)
	// 函数可以作为参数，也可以作为返回值
	res := f1(add, 1, 2)
	fmt.Println(res)
}
