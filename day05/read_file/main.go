package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFileByRead() {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {

		}
	}(fileObj)
	var tmp = make([]byte, 128)
	var content []byte
	for {
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err: %v\n", err)
			return
		}
		content = append(content, tmp[:n]...)
		fmt.Printf("读取了%v字节数据\n", n)
		fmt.Println("读取的数据为：", tmp[:n])
	}
	fmt.Println(content)
}

func ReadFileByBufferIo() {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
		}
	}(fileObj)
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		fmt.Println(line)
	}
}

func ReadFileByIoUtil() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("open file by ioUtil failed, err: ", err)
		return
	}
	fmt.Println(content)
}

func main() {
	// 读取文件的三种方式
	// fileObj.read([]byte, length)
	ReadFileByRead()
	ReadFileByBufferIo()
	ReadFileByIoUtil()
}
