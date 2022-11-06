package main

import (
	"fmt"
	"reflect"
)

type B struct {
}

type Enum int

const (
	Zero Enum = 0
)

type Cat struct {
}

type StructField struct {
	Name      string            `json:"name"` // 字段名
	PkgPath   string            // 字段路径
	Type      reflect.Type      // 字段反射类型对象
	Tag       reflect.StructTag // 字段的结构体标签
	Offset    uintptr           // 字段在结构体中的相对偏移
	Index     []int             // Type.FieldByIndex中的返回的索引值
	Anonymous bool              // 是否为匿名字段
}

func main() {
	// 反射
	var a int
	var b B
	typeOfB := reflect.TypeOf(b)
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	// Name方法返回（内置或者自定义）名称，kind方法返回大种类名称
	fmt.Println(typeOfB.Name(), typeOfB.Kind())
	typeOfZ := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfZ.Name(), typeOfZ.Kind())
	// 指针与指向指针的元素
	c := new(Cat)
	typeOfC := reflect.TypeOf(c)
	// 指针类型的Name为空字符串
	fmt.Println(typeOfC.Name(), typeOfC.Kind())
	// 通过Elem()方法获得指针指向的元素
	typeOfC = typeOfC.Elem()
	fmt.Println(typeOfC.Name(), typeOfC.Kind())
	// 通过反射获取结构体的成员类型
	var tf StructField
	typeOfTf := reflect.TypeOf(tf)
	// 获取结构体的成员数量
	fmt.Println(typeOfTf.NumField())
	for i := 0; i < typeOfTf.NumField(); i++ {
		// 通过索引获取对应成员的信息
		fmt.Println(typeOfTf.Field(i))
		fmt.Println(typeOfTf.Field(i).Name, typeOfTf.Field(i).Tag)
	}
	valueOfA := reflect.ValueOf(a)
	valueOfA.Interface()
	valueOfA.Int()
	valueOfA.NumField()
	valueOfA.Field(0)
}
