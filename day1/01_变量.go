package main

import (
	"fmt"
)

func main() {
	var a int  // 定义变量，格式：var 变量名 数据类型
	a = 30     // 已定义的变量赋值
	var b = 20 // 自动推导类型，先定义变量，再变量赋值
	c := 30    // 自动推导类型，先定义变量，再变量赋值

	fmt.Printf("a = %d ,b = %d , c = %d\n", a, b, c)
	fmt.Printf("a type %T ,b type %T , c type %T\n", a, b, c)

	// 变量多重赋值
	a, b = 10, 20
	fmt.Printf("a = %d,b = %d\n", a, b)

	// 交换变量的值
	b, a = a, b
	fmt.Printf("a = %d,b = %d\n", a, b)

	// _ 匿名变量，丢弃数据,匿名变量配合函数返回值使用才有优势
	x := 40
	y := 50
	a, _ = x, y
	fmt.Printf("a = %d\n", a)

	// 多变量定义，推导赋值
	var (
		e = 11
		f = 12
		g = 13
	)
	fmt.Println(e, f, g)
}
