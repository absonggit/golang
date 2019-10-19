package main

import (
	"fmt"
	// "runtime"
	"time"
)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")

}

func person2() {
	Printer("world")
}

func main() {
	//runtime.GOMAXPROCS(1) // 指定CPU核数运行程序
	//runtime.Gosched()     // 让出时间片，先让别的协程执行完，再回来执行此协程
	//runtime.Goexit()      // 结束当前协程
	go person1()
	go person2()
	// 主协程结束，子协程跟着结束
	for {
	}
}
