package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}
func main() {
	// error  异常处理
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("ERRORS: ", err)
	} else {
		fmt.Println("reslut = ", result)
	}
	// panic 异常处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic 测试")
}
