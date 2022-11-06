package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 32
	// 数字转字符串
	strI := strconv.Itoa(i)
	fmt.Printf("%T\n", strI)
	// 字符串转数字
	intI, err := strconv.Atoi(strI)
	if err != nil {
		fmt.Println("转化时出错")
		return
	}
	fmt.Printf("%T", intI)
	// Parse系列
	// func ParseBool(str string) (value bool, err error)
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	// func ParseFloat(s string, bitSize int) (f float64, err error)
	// Format系列
	// func FormatBool(b bool) string
	// func FormatInt(i int64, base int) string
	// func FormatUint(i uint64, base int) string
	// func FormatFloat(f float64, fmt byte, prec int, bitSize int) string
}
