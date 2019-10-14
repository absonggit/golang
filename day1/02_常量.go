package main

import (
	"fmt"
)

func main() {
	// 关键字const定义常量，常量不允许修改
	const a int = 10
	fmt.Println(a)
	const b = 20 //自动推导类型
	fmt.Println(b)
	fmt.Printf("b type is %T\n", b)

	// 多常量量定义，推导赋值
	const (
		aa = 10
		bb = 20
	)
	fmt.Println(aa, bb)

	// iota枚举,常量自动生成器，从0开始，每一行自动累加1，如果iota在同一行，值相同
	const (
		e = iota
		f
		g
		h, i, j = iota, iota, iota
		k       = iota
	)
	fmt.Println(e, f, g, h, i, j, k)

}
