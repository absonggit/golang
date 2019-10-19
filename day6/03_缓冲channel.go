package main

import (
	"fmt"
)

func main() {
	// 创建一个无缓冲的channel
	ch := make(chan int, 3)
	fmt.Printf("长度：%d\t容量：%d\n", len(ch), cap(ch))
	// 创建协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 写数据
			fmt.Println("写入数据：", i)
		}
		close(ch) // 循环执行完毕，关闭通道
	}()

	// for {
	// 	// 如果ok为true，说明通道没有关闭
	// 	if num, ok := <-ch; ok == true {
	// 		fmt.Println("读取数据：", num)
	// 	} else {
	// 		break
	// 	}
	// }

	// 使用range迭代接收channel,通道关闭自动结束
	for num := range ch {
		fmt.Println("读取数据：", num)
	}
}
