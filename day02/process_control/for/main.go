package main

import "fmt"

// GenerateMulti 生成九九乘法表
func GenerateMulti() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	// for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种：初始化语句写在语句块外面、自增写在语句块里面、无限循环
	// for range循环
	s := "hello沙河"
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
	GenerateMulti()
	// 跳出多层for循环
	// break只能跳出一层for循环
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j == 5 {
				break
			}
			fmt.Println(i * j)
		}
	}
	fmt.Println("-------------------")
	// continue也只能跳过一层循环
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j == 5 {
				continue
			}
			fmt.Println(i * j)
		}
	}
	// 跳出多层循环break + 标签
}
