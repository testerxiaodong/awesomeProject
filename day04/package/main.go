package main

import (
	// 导入别名
	importPackage "awesomeProject/day04/import_package"
	"fmt"
)

// init函数没有参数也没有返回值
func init() {
	fmt.Println("我是package包里的init函数，我被导入时自动执行")
}

func main() {
	// init函数执行时机：全局声明 => init函数 => main函数
	// init函数执行顺序：先执行被导入包的init函数
	result := importPackage.Add(1, 2)
	fmt.Println(result)
}
