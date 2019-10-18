package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// strings.Contains 查询字符串中是否包含go，包含返回true反之为false
	fmt.Println(strings.Contains("hello go", "go"))
	fmt.Println(strings.Contains("hello go", "abc"))

	// strings.Join 组合字符串
	s := []string{"Hello", "Maria", "I", "Love", "You"}
	j := strings.Join(s, "*")
	fmt.Println(j)

	// strings.Index 查找字符串索引下标,不存在返回-1
	fmt.Println(strings.Index(j, "Love"))
	fmt.Println(strings.Index(j, "Bella"))

	// strings.Repeat 重复打印字符串
	fmt.Println(strings.Repeat(j, 3))

	// strings.Split 指定分隔符分离字符串，返回一个切片
	fmt.Println(strings.Split(j, "*"))

	// strings.Trim 去掉两头的指定字符
	fmt.Println(strings.Trim(">>>help me<<<", ">|<"))

	// strings.Fields 以空格做分割将字符串转为切片
	fmt.Println(strings.Fields("Tian wang covers ground tiger"))

	// 转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10) // 10 指以10进制方式追加
	slice = strconv.AppendQuote(slice, "Icoming")
	fmt.Println(string(slice))

	// 其他类型转换为字符串
	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strconv.FormatFloat(3.14, 'f', -1, 64)) // f 以小数方式打印， -1 指小数点位数(-1 紧缩模式)，64 以float64处理
	fmt.Println(strconv.Itoa(123456))

	// 字符串转其他类型
	fmt.Println(strconv.ParseBool("false")) // 这里返回两个值，第二个是 error
	fmt.Println(strconv.Atoi("123"))        // 这里返回两个值，第二个是 error

	// 正则表达式
	c := "Tian wang covers ground tiger"
	// 1 解释规则
	reg1 := regexp.MustCompile(`t.+r`) // `` 引起来的内容 代表原生字符串
	// 2 根据规则匹配字符
	result := reg1.FindAllStringSubmatch(c, -1) // -1 匹配所有
	fmt.Println(result)
}
