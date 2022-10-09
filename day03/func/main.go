package main

import "fmt"

func add(x, y int) (sum int) {
	sum = x + y
	return sum
}

// Min 可变长参数的函数
func Min(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	min := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

// SumSub 多个返回值函数
func SumSub(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// DeferDemo defer语句延迟调用，多个defer语句压栈后进先出
func DeferDemo() {
	fmt.Println("函数开始执行")
	defer fmt.Println("我是defer1")
	defer fmt.Println("我是defer2")
	defer fmt.Println("我是defer3")
	fmt.Println("函数结束执行")
}

func main() {
	// 函数的定义：关键字func
	// 函数的参数、函数的返回值
	result := add(1, 2)
	fmt.Println(result)
	// 匿名函数
	var f1 = func(x, y int) int {
		sum := x + y
		return sum
	}
	fmt.Println(f1(1, 2))
	// 可变长参数
	mini := Min(9, 2, 4, 3, 5, 7, 1, 1, 8)
	fmt.Println(mini)
	// 多个返回值
	sum, sub := SumSub(2, 1)
	fmt.Println(sum, sub)
	// defer语句：延迟调用，多个defer语句压栈后进先出
	// defer语句多用于释放资源
	DeferDemo()
}
