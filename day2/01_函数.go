package main

import (
	"fmt"
)

func MyFunc1() {
	fmt.Println("我是无参无返回值函数")
}

func MyFunc2(x, y int) { //x, y是形参，为了接收实参传递的参数,只能放在形参的最后面
	fmt.Printf("我是有参无返回值函数 ")
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

func MyFunc3(args ...int) { //...int 不定参数类型
	fmt.Printf("传入了%d个参数: ", len(args))
	for i := range args {
		fmt.Printf("%d ", args[i])
	}
}

func main() {
	a, b := 1, 2
	MyFunc1()           // 调用函数
	MyFunc2(a, b)       // a ,b 是实参，传递参数给函数形参
	MyFunc3(5, 3, 7, 9) // 可以传入任意个数的参数
}
