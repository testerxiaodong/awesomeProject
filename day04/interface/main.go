package main

import "fmt"

type eater interface {
	eat()
}

type car interface {
	run()
}

type Person struct {
	name string
	age  int
}

func (p *Person) eat() {
	fmt.Printf("%v is eatting\n", p.name)
}

type Animal struct {
	name string
	kind string
}

func (a *Animal) eat() {
	fmt.Printf("%v is eatting\n", a.name)
}

func doSomething(e eater) {
	e.eat()
}

type Benz struct {
}

func (bz Benz) run() {
	fmt.Println("奔驰在开...")
}

type Bmw struct {
}

func (bm Bmw) run() {
	fmt.Println("宝马在开...")
}

// 空接口：所有类型都实现了空接口
func typeAssert(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串：", t)
	case int:
		fmt.Println("是一个整型：", t)
	case bool:
		fmt.Println("是一个bool值：", t)
	}
}

func main() {
	// 接口是一种类型
	var e eater
	e = &Person{
		"cengdong",
		18,
	}
	// 接口定义了方法的集合，无论什么变量，只要实现了这些方法，就实现了这个接口，可以成为这个接口类型的变量
	doSomething(e)
	// 接口底层存储：动态类型 + 动态值
	fmt.Printf("%T: %v\n", e, e)
	// 使用值接收者实现接口所有方法时，可以传值也可以传指针
	var c car
	bz := Benz{}
	c = bz
	fmt.Println(c)
	bm := &Bmw{}
	c = bm
	fmt.Println(c)
	// 使用指针接收者实现接口所有方法时，只能传指针
	// 接口和类型是多对多的关系
	// 接口嵌套
	// 类型断言
	a := 10
	typeAssert(a)
}
