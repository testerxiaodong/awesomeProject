package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// 格式化
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 只格式化时分秒
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	nowTimeStamp := now.Unix()
	// 时间戳转时间
	// time.Unix()
	transTime := time.Unix(nowTimeStamp, 0)
	fmt.Println(transTime)
	// 时间间隔
	// time.Duration
	nextTime := now.Add(time.Hour * 2)
	fmt.Println(nextTime)
	// 定时器：本质是一个通道
	//time.Tick(time.Second)
}
