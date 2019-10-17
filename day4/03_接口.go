package main

import (
	"fmt"
)

// 定义接口类型
type Humaner interface {
	// 方法，只有申明，没有实现,由自定义类型实现
	sayhi()
}

type Personer interface {
	Humaner // 继承sayhi()
	Sing(song string)
}

// 自定义类型
type Student struct {
	name string
	id   int
}
type Techer struct {
	addr  string
	group string
}
type MyStr string

// 方法
func (tmp *Student) sayhi() {
	fmt.Println("你好", tmp.name)
}

func (tmp *Techer) sayhi() {
	fmt.Println(tmp.addr, tmp.group)
}

func (tmp *MyStr) sayhi() {
	fmt.Println(*tmp)
}

func (tmp *Student) Sing(song string) {
	fmt.Println("唱着：", song)
}

// 定义一个函数，参数为接口类型

func WhoSayHi(i Humaner) {
	i.sayhi()
}

// 空接口万能类型，保存任意类型的值
func NullInterface(args ...interface{}) {
	for _, data := range args {
		switch value := data.(type) { // 类型断言
		case int:
			fmt.Printf("内容是%v 类型为int\n", value)
		case string:
			fmt.Printf("内容是%v 类型为string\n", value)
		case bool:
			fmt.Printf("内容是%v 类型为bool\n", value)
		default:
			fmt.Println("other", value)
		}
	}
}
func main() {
	s := &Student{"Harry", 1}
	t := &Techer{"北京", "go"}
	var str MyStr = "开心"
	// 多态 调用同一函数，实现不同表现
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, i := range x {
		WhoSayHi(i)
	}
	var i1 Personer
	s1 := &Student{"Tom", 2}
	i1 = s1
	i1.sayhi()
	i1.Sing("心太软")

	//任意类型都可以传参
	NullInterface("asadsa")
	NullInterface(1)
	NullInterface(false)
}
