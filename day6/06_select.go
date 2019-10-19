package main

import (
	"fmt"
)

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x: // 写入数据到通道
			x, y = y, x+y
		case flag := <-quit: // 读取通道中的数据
			fmt.Printf("\n退出 %t\n", flag)
			return
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	// 子协程
	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch // 接收通道中的数据
			fmt.Printf("%d ", num)
		}
		quit <- true // 发送数据到通道
	}()

	// 主协程
	fibonacci(ch, quit)
}
