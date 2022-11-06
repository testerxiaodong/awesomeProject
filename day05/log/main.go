package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	fileObj, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file failed, err; ", err)
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {

		}
	}(fileObj)
	log.SetOutput(fileObj)
}

func main() {
	fileObj, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file failed, err; ", err)
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {

		}
	}(fileObj)
	// 设置日志flag
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	fmt.Println(log.Flags())
	log.SetOutput(fileObj)
	log.Println("这是一条测试日志")
	log.Printf("这是一条格式化的测试日志，时间：%v\n", time.Now())
	//log.Fatal("这是一条fatal级别日志")
	// 日志器构造函数
	// func New(out io.Writer, prefix string, flag int) *Logger
}
