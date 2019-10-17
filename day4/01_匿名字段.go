package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 匿名字段，继承Person的成员
	id     int
	addr   string
}

func main() {
	// 顺序初始化
	var p1 Student = Student{Person{"Harry", 'm', 23}, 1, "sz"}
	fmt.Println(p1)
	// 自动推导
	p2 := Student{Person{"Mike", 'm', 22}, 2, "bj"}
	// %+v 详细显示
	fmt.Printf("%+v\n", p2)
	// 没有初始化的会自动赋值为0
}
