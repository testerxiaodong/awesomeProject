package main

import "fmt"

// SumArray 求整型切片的和
func SumArray(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	// 数组：存放相同数据类型的容器
	// 数组声明必须指定长度或通过元素个数推导
	// 数组的长度和容量是数组类型的一部分
	var a1 [3]bool
	// 长度自动推导
	var a2 = [...]int{1, 2, 3}
	fmt.Println(a1)
	fmt.Println(a2)
	// 按照索引部分初始化
	a3 := [10]int{1: 1, 2: 2}
	fmt.Println(a3)
	// 数组的遍历：for range
	for i, v := range a3 {
		fmt.Println(i, v)
	}
	// 数组遍历：for循环
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	// 多维数组
	a4 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a4)
	// 多维数组遍历
	for _, v := range a4 {
		for _, value := range v {
			fmt.Println(value)
		}
	}
	// 数组是值类型
	b1 := [...]int{1, 2, 3}
	b2 := b1
	// 要么传递数组指针
	b2[0] = 100
	fmt.Println(b1)
	fmt.Println(b2)
	result := SumArray(a2[:])
	fmt.Println(result)
}
