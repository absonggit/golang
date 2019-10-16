package main

import (
	"fmt"
)

func Delkv(m map[int]string) {
	delete(m, 1) // 删除键值
}
func main() {
	// 定义一个变量，类型为map[int]string。键是唯一的，map有len没有cap.
	// 给map 赋值时如果指定的键不存在，则增加这对键值

	m1 := make(map[int]string)
	fmt.Println(m1)
	m1[1] = "harry"
	m1[2] = "wayne"
	fmt.Println(m1)
	m1[1] = "carry"
	fmt.Println(m1[1])

	// 遍历map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 判断一个键值是否存在
	Delkv(m1)
	v, ok := m1[1]
	if ok == true {
		fmt.Println(v)
	} else {
		fmt.Println("键值不存在")
	}
}
