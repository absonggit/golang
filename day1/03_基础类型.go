package main

import (
	"fmt"
)

func main() {
	var (
		num   int        = 10    //整型
		f1    float64    = 3.14  //浮点型
		b1    byte       = 'a'   //字符类型
		str1  string     = "abc" //字符串
		bool1 bool       = false //布尔
		c1    complex128 = 2.1 + 3.14i
	)
	fmt.Printf("num type is %T\nf1 type is %T\nb1 type is %T\nstr1 type is %T\nbool1 type is %T \nc1 type is %T \n", num, f1, b1, str1, bool1, c1)
	fmt.Printf("%c,%d\n", b1, b1)
	fmt.Println(real(c1), imag(c1)) //取复数的实部和虚部
	// 类型转换，兼容的类型才可以转换
	// 类型别名
	type bigint int64
	type (
		myint int
		mystr string
	)
}
