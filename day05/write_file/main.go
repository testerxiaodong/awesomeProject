package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteFileByWrite() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer fileObj.Close()
	str := "hello沙河"
	n, err := fileObj.Write([]byte(str))
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
	fmt.Printf("write %v 字节\n", n)
	n, err = fileObj.WriteString(str)
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
	fmt.Printf("write %v 字节\n", n)
}

func WriteFileByBufferIo() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer fileObj.Close()
	str := "hello沙河"
	writer := bufio.NewWriter(fileObj)
	n, err := writer.Write([]byte(str))
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
	fmt.Printf("write %v 字节\n", n)
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
	fmt.Printf("write %v 字节\n", n)
	writer.Flush()
}

func WriteFileByIoUtil() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./write.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file by ioUtil, err: ", err)
		return
	}
}

func main() {
	//WriteFileByWrite()
	WriteFileByBufferIo()
	WriteFileByIoUtil()
}
