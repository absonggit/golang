package main

import (
	"fmt"
)

func main() {
	//算数
	a, b := 40, 30
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d ÷ %d = %d\n", a, b, a%b)
	a++
	b--
	fmt.Printf("a++ = %d\n", a)
	fmt.Printf("b++ = %d\n", b)

	// 关系比较
	fmt.Printf("%d > %d is %t\n", a, b, a > b)
	fmt.Printf("%d < %d is %t\n", a, b, a < b)
	fmt.Printf("%d >= %d is %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d is %t\n", a, b, a <= b)
	fmt.Printf("%d != %d is %t\n", a, b, a != b)

	// 逻辑运算
	fmt.Printf("!(%d != %d) is %t\n", a, b, !(a != b))
	fmt.Println("true && false is", true && false)
	fmt.Println("true || false is", true || false)

	// 位运算

	fmt.Println(b << 2)
	fmt.Println(b >> 2)
	fmt.Println(b & a)
	fmt.Println(b | a)
	fmt.Println(b ^ a)

	// 赋值运算
	a += 2
	fmt.Println(a)
	a -= 2
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	a %= 2
	fmt.Println(a)
	a <<= 2
	fmt.Println(a)
	a >>= 2
	fmt.Println(a)
	a &= 2
	fmt.Println(a)
	a ^= 2
	fmt.Println(a)
	a |= 2
	fmt.Println(a)
	// 其他运算符
	// &a  取a的内存地址
	// *a  指针变量a所指向内存的值
	// 使用()括起来的运算会先执行
}
