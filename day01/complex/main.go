package main

import "fmt"

var (
	c1 complex64
	c2 complex128
)

func main() {
	// 复数
	c1 = 1 + 2i
	c2 = 1 + 3i
	// complex64
	fmt.Printf("%T: %v\n", c1, c1)
	fmt.Printf("%T: %v\n", c2, c2)
	// complex128
}
