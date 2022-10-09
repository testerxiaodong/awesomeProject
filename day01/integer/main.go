package main

import "fmt"

var (
	uint8Num  uint8
	uint16Num uint16
	uint32Num uint32
	uint64Num uint64
)

var (
	int8Num  int8
	int16Num int16
	int32Num int32
	int64Num int64
)

var (
	uintNum uint
	intNum  int
	uintPtr uintptr
)

func main() {
	// 整型：默认是int型
	// 无符号整型 uint8、uint16、uint32、uint64
	fmt.Printf("%T: %v\n", uint8Num, uint8Num)
	fmt.Printf("%T: %v\n", uint16Num, uint16Num)
	fmt.Printf("%T: %v\n", uint32Num, uint32Num)
	fmt.Printf("%T: %v\n", uint64Num, uint64Num)
	// 有符号整型 int8、int16、int32、int64
	fmt.Printf("%T: %v\n", int8Num, int8Num)
	fmt.Printf("%T: %v\n", int16Num, int16Num)
	fmt.Printf("%T: %v\n", int32Num, int32Num)
	fmt.Printf("%T: %v\n", int64Num, int64Num)
	// 特殊整型 uint、int、uintptr
	fmt.Printf("%T: %v\n", uintNum, uintNum)
	fmt.Printf("%T: %v\n", intNum, intNum)
	fmt.Printf("%T: %v\n", uintPtr, uintPtr)
	// 进制转换：二进制、八进制、十进制、十六进制
}
