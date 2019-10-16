package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(a *[10]int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	var a [5]int       // 数组，创建同一类型的集合，[n]代表元素的个数，n必须是常量
	for i := range a { // 使用循环给a 数组赋值
		a[i] = i + 1
	}
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5} // 没有初始化的元素，自动赋值为0
	fmt.Println(b)

	c := [5]int{2: 10, 4: 20} // 使用下表赋值，下标:值
	fmt.Println(c)

	// 二维数组
	d := [3][4]int{
		{1, 3, 5, 7},
		{2, 4, 6, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(d)

	// 数组比较，只支持 == 或 !==，比较每个元素是否一样，两个数组类型必须一致
	fmt.Println(a == b)
	fmt.Println(b != c)

	// 随机数
	var e [10]int
	rand.Seed(time.Now().UnixNano()) // 使用当前时间数作为种子
	for i := 0; i < len(e); i++ {
		e[i] = rand.Intn(100) // 产生100以内的整数
	}
	fmt.Println("随机产生100以内的数组:", e)

	// 冒泡排序，数组函数传递
	BubbleSort(&e) // 数组做函数参数，是值传递，形参是实参的副本。所以这里使用数组指针传参给函数
	fmt.Println("使用指针的冒泡排序:", e)
}
