package main

import (
	"fmt"
)

func swap(i, j *int) {
	*i, *j = *j, *i
}
func main() {
	var (
		a int  // 初始化变量a
		b *int // 初始化指针变量b
	)
	a = 10 // 变量a赋值10
	b = &a // b指针指向a的内存地址
	fmt.Println("取变量a的内存地址", &a)
	fmt.Println("对b的指针取值", *b)

	*b = 20 // 赋值b指针所指向的值为20
	fmt.Println("a取值", a)
	fmt.Println("对b的指针取值", *b)
	// 注意b指针必须指向某个内存地址才能赋值，可以new()函数创建空指针

	c := new(int) //new函数动态申请内存空间，使用完毕后自回收
	*c = 666
	fmt.Println("c指针变量的值为", *c)

	//指针做函数参数
	d, e := 10, 20
	swap(&d, &e) //地址传递
	fmt.Println(d, e)
}
