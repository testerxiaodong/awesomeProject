package import_package

import "fmt"

func init() {
	fmt.Println("我是import_package包里的init函数，我被导入时自动执行")
}

func Add(x, y int) int {
	return x + y
}
