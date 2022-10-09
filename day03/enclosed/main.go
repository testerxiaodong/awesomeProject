package main

import (
	"fmt"
	"time"
)

func f1(f func()) {
	fmt.Println("我是函数f1")
	f()
}

func f2(x, y int) {
	fmt.Println("我是函数f2")
	fmt.Println(x + y)
}

func f3(f func(x, y int), x, y int) func() {
	f(x, y)
	return func() {
		fmt.Println("我是闭包里的函数")
	}
}

// 生成斐波那契数列
func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = f1+f2, f2
		return f1
	}
}

// Counter 函数运行次数计数
func Counter(f func()) func() int {
	count := 0
	return func() int {
		f()
		count++
		return count
	}
}

// Timer 函数运行时间计算
func Timer(f func()) func() {
	return func() {
		time1 := time.Now().UnixMilli()
		f()
		time2 := time.Now().UnixMilli()
		//fmt.Println(time2 - time1)
		fmt.Printf("函数运行开始时间为%v\n", time1)
		fmt.Printf("函数运行结束时间为%v\n", time2)
		fmt.Printf("函数的运行时长为%v\n", time2-time1)
	}
}

func foo() {
	fmt.Println("foo run")
}

func foo1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	// 闭包
	f := f3(f2, 1, 2)
	f1(f)
	// 隔离数据
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
	cnt := Counter(foo)
	cnt()
	cnt()
	cnt()
	fmt.Println(cnt())
	timer := Timer(foo1)
	timer()
}
