package main

import (
	"fmt"
)

func main() {

	// if 语句
	var status string
	if age := 20; age > 0 && age < 18 {
		status = "boy"
	} else if age > 18 && age < 50 {
		status = "man"
	} else {
		status = "old man"
	}

	if name := "Harry"; name == "Harry" {
		fmt.Printf("你好, %s %s\n", name, status)
	} else {
		fmt.Printf("%s,你不是 Harry%s,Fuck off!\n", name, status)
	}

	// switch 语句
	var num int
	fmt.Printf("输入要进入的楼层编号0-9：")
	fmt.Scan(&num)
	switch num {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 0:
		fmt.Printf("\n按下%d楼\n", num)
	default:
		fmt.Println("非法输入")
	}
}
