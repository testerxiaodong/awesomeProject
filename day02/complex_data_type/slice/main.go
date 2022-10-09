package main

import "fmt"

func main() {
	// 切片的定义
	var s1 []int
	var s2 []string
	fmt.Println(s1)
	fmt.Println(s2)
	// 数组初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"cengdong", "xiaoming"}
	fmt.Println(s1)
	fmt.Println(s2)
	// 切片的长度和容量
	fmt.Printf("len: %v: cap: %v\n", len(s1), cap(s1))
	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4}
	s3 := a1[1:]
	fmt.Printf("%T: %v\n", s3, s3)
	// make函数创造切片
	s4 := make([]int, 5, 10)
	fmt.Println(s4)
	// 切片是引用类型，不可以和其他切片比较，可以和nil比较
	// 切片定义后没有初始化，值是nil
	// 改变底层数组，切片的值也随之改变
	fmt.Println("---------")
	// 切片追加元素：append
	s3 = append(s3, s4...)
	fmt.Println(s3)
	// 切片删除元素：append
	s3 = append(s3[:3], s3[4:]...)
	fmt.Println(s3)
	// 切片拷贝：copy(dst, src)
	x1 := []int{1, 3, 5}
	x2 := x1
	var x3 = make([]int, 3)
	copy(x3, x1)
	fmt.Println(x1, x2, x3)
}
