package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 把go语言中结构体转换为json
	// 把json格式字符串转换为go语言结构体
	p1 := &Person{
		"cengdong",
		18,
	}
	pb, err := json.Marshal(*p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(pb))
	str := `{"name":"cengdong","age":18}`
	p2 := new(Person)
	err = json.Unmarshal([]byte(str), p2)
	if err != nil {
		fmt.Printf("unmarshal failed, err: %v\n", err)
	}
	fmt.Println(p2)
}
