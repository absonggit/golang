package main

import (
	"fmt"
)

func producer(in chan<- int) {
	for i := 0; i < 5; i++ {
		in <- i
		fmt.Println("写入：", i)
	}
	close(in)
}

func customer(out <-chan int) {
	for data := range out {
		fmt.Println("读取：", data)
	}
}
func main() {
	ch := make(chan int)

	// 生产者，写入channel
	go producer(ch)

	// 消费者，读取channel
	customer(ch)

}
