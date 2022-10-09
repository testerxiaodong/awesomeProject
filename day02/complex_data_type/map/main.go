package main

import (
	"fmt"
	"strings"
)

func countWord(s string) (hashMap map[string]int) {
	hashMap = make(map[string]int)
	stringSlice := strings.Split(s, " ")
	for _, v := range stringSlice {
		if _, ok := hashMap[v]; ok {
			hashMap[v] += 1
		} else {
			hashMap[v] = 1
		}
	}
	return hashMap
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	// map定义
	var m1 map[string]int
	fmt.Println(m1)
	// 必须初始化
	m1 = make(map[string]int, 8)
	m1["理想"] = 18
	m1["cengdong"] = 23
	fmt.Println(m1)
	// 判断key是否在map中
	if p, ok := m1["理想"]; !ok {
		fmt.Println("不在map中")
	} else {
		fmt.Println(p)
	}
	// map遍历：只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// map遍历：只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除某个键
	delete(m1, "cengdong")
	fmt.Println(m1)
	// 删除不存在的键，不会报错
	result := countWord("how do you do")
	fmt.Println(result)
	palindromeFlag := isPalindrome("abcba")
	fmt.Println(palindromeFlag)
}
