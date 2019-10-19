package main

import (
	"fmt"
	"time"
)

// 创建通道
var ch = make(chan int)
var ch1 = make(chan string)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
	ch <- 666 // 往通道写入数据

}

func person2() {
	<-ch // 如果通道没有数据则会阻塞
	Printer("world")
	ch1 <- "子协程完毕"

}

func main() {

	defer fmt.Println("主协程结束")

	go person1()
	go person2()

	// 主协程结束，子协程跟着结束
	str := <-ch1
	fmt.Println(str)
}
