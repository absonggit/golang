package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(slice []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		slice[i] = rand.Intn(100)
	}

}

func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
func main() {
	// 数组切片(动态数组)是通过内部指针和相关属性引用数组片段,指针指向底层数组
	var s1 []int
	s1 = []int{1, 2, 3, 4, 5}
	s2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(s1, s2)
	array := [5]int{1, 2, 3, 4, 5}                                                                        // 数组
	slice := []int{1, 2, 3, 4, 5}                                                                         // 切片
	s3 := array[1:4:5]                                                                                    // [1:4:5] [low:high:max] low = 下标；len = high - low ；cap = max - low
	fmt.Printf("初始化切片:%v, 切片长度:%v,切片容量:%v,切片截取[2:4:5]为%v\n", slice, len(slice), cap(slice), slice[2:4:5]) // 切片截取
	fmt.Printf("数组:%v, 切片:%v, 切片长度:%v,切片容量:%v\n", array, s3, len(s3), cap(s3))
	slice = append(slice, 6) // 在切片尾部增加元素，如果超过容量，通常以两倍自动扩容
	fmt.Println("切片增加元素", slice, len(slice), cap(slice))
	s4 := make([]int, 5, 10) // []int 切片类型 5 切片长度 10 切片容量(如果没有申明，则和长度一样)
	fmt.Printf("make()创建切片:%v, 切片长度:%v,切片容量:%v\n", s4, len(s4), cap(s4))

	// copy()
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s5 := data[8:]
	s6 := data[:5]
	fmt.Println("s5切片", s5)
	fmt.Println("s6切片", s6)
	copy(s6, s5) //用于将源的数据（第二个参数），复制到目标（第一个参数）。返回值为拷贝了的数据个数，是len(dst)和len(src)中的最小值。
	fmt.Println("copy后的s5", s5)
	fmt.Println("copy后的s6", s6)
	fmt.Println("data切片", data)

	s7 := []int{1, 2}
	s8 := []int{6, 7, 8, 9}
	copy(s8, s7)
	fmt.Println(s8)

	// 切片做函数参数
	slice1 := make([]int, 10)
	InitData(slice1)
	fmt.Println("排序前: ", slice1)
	BubbleSort(slice1)
	fmt.Println("排序后: ", slice1)
}
