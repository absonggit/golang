package main

import (
	"fmt"
)

// 定义结构体
type Person struct {
	name string
	age  int
	sex  byte
}

// 继承Person字段，成员和方法都继承
type Student struct {
	Person
	id   int
	addr string
}

// 带有接收者的函数叫方法 只要接收者的类型不一样，方法就算同名，也不会出现重复定义错误
// 创建方法，打印信息
func (tmp Person) PrintInfo() {
	fmt.Printf("name=%s,sex=%c,age=%d\n", tmp.name, tmp.sex, tmp.age)
}

// 创建方法，修改信息
func (p *Person) SetInfo(n string, a int, s byte) {
	p.name = n
	p.age = a
	p.sex = s
}

func (p *Student) SetInfo(n string, a int, s byte, i int, addr string) {
	p.name = n
	p.age = a
	p.sex = s
	p.id = i
	p.addr = addr
}

func main() {
	// 定义同时初始化
	p1 := Person{"Harry", 23, 'm'}

	// 调用方法
	p1.PrintInfo()
	p1.SetInfo("哈利", 23, 'm')
	p1.PrintInfo()

	// p2 继承了方法，如果有相同的方法，先找本作用域的方法，找不到再用继承的方法
	p2 := Student{Person{"Tom", 24, 'm'}, 1, "sz"}
	p2.PrintInfo()
	p2.SetInfo("汤姆", 24, 'm', 1, "sz")
	p2.PrintInfo()
}
