package main

import (
	"fmt"
)

// 定义结构体

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test01(s Student) {
	s.age = 33
	fmt.Println("s5值传递:", s)
}

func test02(s *Student) {
	s.age = 33
	fmt.Println("s5指针传递:", *s)
}
func main() {
	// 顺序初始化，每一个结构体成员都要初始化
	var s1 Student = Student{1, "张三", 'm', 23, "中国"}
	fmt.Println(s1)

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{id: 2, name: "李四"}
	fmt.Println(s2)

	// 指针初始化
	s3 := &Student{id: 3, name: "王五"}
	fmt.Println(*s3)

	// 修改成员的值
	s3.addr = "japan"
	fmt.Println(*s3)

	// new()申明结构体
	s4 := new(Student)
	fmt.Println(*s4)

	// 结构体比较
	s5 := Student{1, "张三", 'm', 23, "中国"}
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s5)

	// 相同结构体变量可以相互赋值

	// 结构体做函数参数值传递,形参无法修改实参
	test01(s5)
	fmt.Println("s5源值:", s5)

	// 结构体做函数参数指针传递,形参实参联动变化
	test02(&s5)
	fmt.Println("s5源值:", s5)
}
