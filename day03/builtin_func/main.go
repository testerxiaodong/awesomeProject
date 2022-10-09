package main

import "fmt"

func funcA() {
	fmt.Println("我是函数a")
}

func funcB() {
	//defer func() {
	//	fmt.Println("释放数据库链接")
	//}()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("recover in B")
		}
	}()
	panic("出错啦")
}

func funcC() {
	fmt.Println("这是函数c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
