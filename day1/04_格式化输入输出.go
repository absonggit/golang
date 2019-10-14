package main

import (
	"fmt"
)

func main() {
	/* fmt.Printf   格式化输出,
	%T 变量所属类型
	%d 整数类型
	%f 浮点型
	%s 字符串
	%c 字符型
	%v 自动匹配格式
	%t 布尔
	*/
	var a int
	fmt.Printf("输入整数赋值给变量a：")
	fmt.Scan(&a)
	fmt.Printf("\n变量a = %d\n", a)

}
