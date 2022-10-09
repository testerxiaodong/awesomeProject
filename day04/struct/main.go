package main

import "fmt"

// myInt 自定义类型
type myInt int

// 类型别名
// type rune = int32

// Address 结构体嵌套
type Address struct {
	province string
	city     string
}

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
	*Address
}

type Company struct {
	name string
	*Address
}

// Person构造函数
func newPerson(name string, age int, gender string, hobby []string, address *Address) *Person {
	return &Person{
		name,
		age,
		gender,
		hobby,
		address,
	}
}

// Person类型work方法
func (p *Person) work() {
	fmt.Printf("%v is working\n", p.name)
}

// Address类型show方法，看看继承
func (ad *Address) show() {
	fmt.Printf("我是一个地址，我在%v\n", ad.province)
}

func main() {
	var mint myInt
	fmt.Printf("%T\n", mint)
	// 定义结构体
	person := newPerson("cengdong", 18, "男", []string{"吃饭", "睡觉", "打豆豆"},
		&Address{"湖南", "娄底"})
	fmt.Println(person)
	// 调用方法
	person.work()
	// 调用继承方法
	person.show()
	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "嘿嘿嘿"
	s.age = 100
	fmt.Printf("type: %T, value: %v\n", s, s)
	// 结构体是值类型
	// 结构体内部内存是连续的
}
