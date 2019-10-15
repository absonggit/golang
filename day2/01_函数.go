package main

import (
	"fmt"
	"os"
)

func MyFunc1() { // 函数名首字母大写为public，小写为private
	fmt.Println("无参无返回值函数")
}

func MyFunc2(x, y int) { //x, y是形参，为了接收实参传递的参数
	fmt.Printf("有参无返回值函数 ")
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

func MyFunc3(args ...int) { //...int 不定参数类型，传入的参数可以是0个或多个,只能放在形参的最后面
	fmt.Printf("传入了%d个参数: ", len(args))
	for i := range args {
		fmt.Printf("%d ", args[i])
	}
	MyFunc4(args...)     // 不定参数传递
	MyFunc5(args[:2]...) // 字符串切片来传递指定个数的参数传递
}

func MyFunc4(tmp ...int) {
	fmt.Printf("\n不定参数传递：\n")
	for i, data := range tmp {
		fmt.Printf("\t%d,%d\n", i, data)
	}
}

func MyFunc5(tmp ...int) {

	fmt.Println("传递了指定个数的参数", tmp)
}

func MyFunc6() (result string) {
	result = "无参有一个返回值"
	return // 有返回值，需要return 关键字
}
func MyFunc7() (id int, result string) {
	id = 1
	result = "无参有多个返回值"
	return // 有返回值，需要return 关键字
}

func MyFunc8(a int, b string) (id int, result string) {
	id, result = a, b
	return // 有返回值，需要return 关键字
}

func test1() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func test2(i int) int {
	if i == 1 {
		return 1
	}
	return i + test2(i-1)
}

func Add(a, b int) int {
	return a + b // 加法
}

func Minus(a, b int) int {
	return a - b // 减法
}

func Mul(a, b int) int {
	return a * b // 乘法
}

func Div(a, b int) int {
	return a / b // 求商
}

// type 定义函数类型,FuncType 就是函数类型
type FuncType func(int, int) int

// 回调函数：函数有一个参数是函数类型
// 多态，调用一个接口，实现不同表现
func Calc(a, b int, Ftest FuncType) (result int) {
	result = Ftest(a, b)
	return
}

// 闭包是由函数和与其相关的引用环境组合而成的实体
func test3() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	a, b := 1, 2

	MyFunc1() // 调用函数

	MyFunc2(a, b) // a ,b 是实参，传递参数给函数形参

	MyFunc3(5, 3, 7, 9) // 可以传入任意个数的参数

	str1 := MyFunc6() // 函数返回值赋值
	fmt.Println(str1) // 打印变量

	r1, r2 := MyFunc7() // 函数多个返回值赋值
	fmt.Println(r1, r2) // 打印变量

	a1, a2 := MyFunc8(b, "有参有返回值")
	fmt.Println(a1, a2)

	sum1 := test1()
	fmt.Println("for循环计算1-100的累加和：", sum1)

	sum2 := test2(100)
	fmt.Println("函数递归计算1-100的累加和：", sum2)

	Ftest := Add // Ftest为一个函数类型的变量
	result := Ftest(10, 20)
	fmt.Println("定义函数类型，计算两个数相加和：", result)

	Ftest = Minus
	result = Ftest(10, 20)
	fmt.Println("定义函数类型计，算两个数相减差：", result)

	result = Calc(10, 20, Add)
	fmt.Println("使用函数回调，计算两个数相加和：", result)

	result = Calc(10, 20, Minus)
	fmt.Println("使用函数回调，算两个数相减差：", result)

	result = Calc(10, 20, Mul)
	fmt.Println("使用函数回调，算两个数相乘积：", result)

	result = Calc(10, 20, Mul)
	fmt.Println("使用函数回调，算两个数相乘积：", result)

	result = Calc(20, 10, Div)
	fmt.Println("使用函数回调，算两个数相除商：", result)

	// 匿名函数
	f1 := func() {
		fmt.Println("定义匿名函数，再调用", a)
	}
	// 匿名函数取别名
	type AnonType func()
	f2 := f1
	f2()

	func() {
		fmt.Println("同时定义、调用匿名函数", b)
	}()
	x, y := 10, 20
	func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(x, y)
	fmt.Println("有参有返回值的匿名函数：", x, y)

	defer fmt.Println("defer 延迟执行，函数结束前执行")
	defer fmt.Println("defer 1") // 多个defer，保持先入后出的原则来执行，即使函数出错，也会执行defer 后面的代码
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	f3 := test3() // f3 来调用闭包函数，不关心变量或常量是否已经超过作用域，所以只要闭包还在使用它，这些变量就还存在
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())

	num1 := 10
	num2 := 20
	defer func(a, b int) {
		fmt.Println("defer与匿名函数结合使用 1 ", a, b)
	}(num1, num2) // 这里的变量传递是已经传入匿名函数，只是没有调用
	num1 = 100
	num2 = 200
	fmt.Println("defer与匿名函数结合使用 2 ", a, b)

	// 接收到的命令行参数，都是以字符串传递的
	list := os.Args
	fmt.Println(list)
	fmt.Println(len(list))
}
