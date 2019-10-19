package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer 创建延时定时器
	timer := time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("子协程执行")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("时间到")

	<-time.After(2 * time.Second)
	fmt.Println("时间到")

	timer.Reset(4 * time.Second) // 重新设置
	<-timer.C
	fmt.Println("时间到")

	// Ticker 创建周期定时器
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println(i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}

}
