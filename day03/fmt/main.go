package main

import "fmt"

var (
	name string
	age  int
	char rune
)

func main() {
	// 格式化输出
	name, age, char = "cengdong", 18, 'c'
	fmt.Printf("%v\n", name)
	fmt.Printf("%T\n", char)
	fmt.Printf("%o\n", age)
	fmt.Printf("%x\n", age)
	fmt.Printf("%b\n", age)
	fmt.Printf("%c\n", char)
	// 换行输出
	fmt.Println("----------")
	// 字符串拼接：fmt.Sprintf
}
