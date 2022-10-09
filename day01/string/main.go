package main

import (
	"fmt"
	"strings"
)

func main() {
	// 转义符：\
	path := "/Users/cengdong"
	fmt.Println(path)
	// 多行字符串
	s1 := `床前明月光，疑似地上霜。
举头望明月，低头思故乡。`
	fmt.Println(s1)
	// 字符串拼接：加号或者Sprintf
	s1 += path
	fmt.Println(s1)
	// Sprintf
	path = fmt.Sprintf("%s_%s", path, path)
	fmt.Println(path)
	// 字符串分隔
	ret := strings.Split(path, "_")
	fmt.Println(ret)
	// 切片拼接字符串
	joinRet := strings.Join(ret, " ")
	fmt.Println(joinRet)
	// 包含
	containFlag := strings.Contains(path, "ceng")
	fmt.Println(containFlag)
	// 前缀/后缀
	prefixFlag := strings.HasPrefix(path, "abc")
	suffixFlag := strings.HasSuffix(path, "abc")
	fmt.Println(prefixFlag)
	fmt.Println(suffixFlag)
	// 字符串遍历：range遍历
	s2 := "hello沙河"
	for _, c := range s2 {
		fmt.Printf("%T: %c\n", c, c)
	}
	// len函数返回的是字符串字节数
	fmt.Println(len(s2))
	// 字符串修改：不能直接修改，先转换为rune切片，对切片进行修改，再转成字符串
	s3 := "白萝卜"
	s4 := []rune(s3)
	s4[0] = '红'
	fmt.Println(string(s4))
	// 单个字符默认是int32类型
	// byte是uint8
}
